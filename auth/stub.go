package auth

import (
	"bytes"
	"crypto/sha256"
	"github.com/mds796/CSGY9223-Final/user"
	"net/http"
	"reflect"
	"time"
)

const (
	LOGGED_OUT = iota // LOGGED_OUT == 0
	LOGGED_IN  = iota // LOGGED_IN == 1
)

type StubService struct {
	UserService   user.Service
	PasswordCache map[string][]byte      // (UUID, sha256(password))
	StatusCache   map[string]int         // (UUID, status)
	CookieCache   map[string]http.Cookie // (username, cookie)
}

func CreateStub(userService user.Service) Service {
	stub := new(StubService)
	stub.UserService = userService
	stub.PasswordCache = make(map[string][]byte)
	stub.StatusCache = make(map[string]int)
	stub.CookieCache = make(map[string]http.Cookie)
	return stub
}

func (s *StubService) Register(request RegisterAuthRequest) (RegisterAuthResponse, error) {
	// create the user in user service
	createUserRequest := user.CreateUserRequest{Username: request.Username}
	createUserResponse, err := s.UserService.Create(createUserRequest)

	// something went wrong in user service
	if err != nil {
		return RegisterAuthResponse{}, &RegisterAuthError{request.Username}
	}

	// register the user
	h := sha256.New()
	h.Write([]byte(request.Password))
	s.PasswordCache[createUserResponse.Uuid] = h.Sum(nil)
	s.StatusCache[createUserResponse.Uuid] = LOGGED_IN
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: request.Username, Expires: expiration}
	s.CookieCache[request.Username] = cookie
	return RegisterAuthResponse{Cookie: cookie}, nil
}

func (s *StubService) Login(request LoginAuthRequest) (LoginAuthResponse, error) {
	// view the user in user service
	viewUserRequest := user.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}

	// check user password
	h := sha256.New()
	h.Write([]byte(request.Password))
	if bytes.Equal(s.PasswordCache[viewUserResponse.Uuid], h.Sum(nil)) {
		// login the user, their current status is irrelevant
		s.StatusCache[viewUserResponse.Uuid] = LOGGED_IN
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: request.Username, Expires: expiration}
		s.CookieCache[request.Username] = cookie
		return LoginAuthResponse{Cookie: cookie}, nil
	} else {
		return LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}
}

func (s *StubService) Verify(request VerifyAuthRequest) (VerifyAuthResponse, error) {
	// check if there is a username assigned to that cookie
	for username, cookie := range s.CookieCache {
		if reflect.DeepEqual(cookie, request.Cookie) {
			response := VerifyAuthResponse{Username: username}
			return response, nil
		}
	}

	// cookie is not known to us
	return VerifyAuthResponse{}, &VerifyAuthError{request.Cookie}
}

func (s *StubService) Logout(request LogoutAuthRequest) (LogoutAuthResponse, error) {
	// view the user in user service
	viewUserRequest := user.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return LogoutAuthResponse{}, &LogoutAuthError{request.Username}
	}

	// logout the user, their current status is irrelevant
	s.StatusCache[viewUserResponse.Uuid] = LOGGED_OUT
	delete(s.StatusCache, viewUserResponse.Uuid)
	return LogoutAuthResponse{}, nil
}
