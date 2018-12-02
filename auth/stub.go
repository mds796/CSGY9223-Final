package auth

import (
	"bytes"
	"context"
	"crypto/sha256"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"log"
	"net/http"
	"time"
)

const (
	LOGGED_OUT = iota // LOGGED_OUT == 0
	LOGGED_IN  = iota // LOGGED_IN == 1
)

type StubService struct {
	UserService   userpb.UserClient
	PasswordCache map[string][]byte // (UUID, sha256(password))
	StatusCache   map[string]int    // (UUID, status)
	CookieCache   map[string]string // (username, cookie)
}

func DecodeCookie(cookie string) *http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}
	return request.Cookies()[0]
}

func CreateStub(userService userpb.UserClient) *StubService {
	stub := new(StubService)
	stub.UserService = userService
	stub.PasswordCache = make(map[string][]byte)
	stub.StatusCache = make(map[string]int)
	stub.CookieCache = make(map[string]string)
	return stub
}

func (s *StubService) Register(ctx context.Context, request *authpb.RegisterAuthRequest) (*authpb.RegisterAuthResponse, error) {
	// create the user in user service
	createUserRequest := &userpb.CreateUserRequest{Username: request.Username}
	createUserResponse, err := s.UserService.Create(ctx, createUserRequest)

	// something went wrong in user service
	if err != nil {
		log.Printf("[AUTH] %v", err)
		return &authpb.RegisterAuthResponse{}, &RegisterAuthError{request.Username}
	}

	// register the user
	h := sha256.New()
	h.Write([]byte(request.Password))
	s.PasswordCache[createUserResponse.UID] = h.Sum(nil)
	s.StatusCache[createUserResponse.UID] = LOGGED_IN
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := (&http.Cookie{Name: request.Username, Value: createUserResponse.UID, Expires: expiration}).String()
	s.CookieCache[request.Username] = cookie
	return &authpb.RegisterAuthResponse{Cookie: cookie}, nil
}

func (s *StubService) Login(ctx context.Context, request *authpb.LoginAuthRequest) (*authpb.LoginAuthResponse, error) {
	// view the user in user service
	viewUserRequest := &userpb.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(ctx, viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return &authpb.LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}

	// check user password
	h := sha256.New()
	h.Write([]byte(request.Password))
	if bytes.Equal(s.PasswordCache[viewUserResponse.UID], h.Sum(nil)) {
		// login the user, their current status is irrelevant
		s.StatusCache[viewUserResponse.UID] = LOGGED_IN
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := (&http.Cookie{Name: request.Username, Value: viewUserResponse.UID, Expires: expiration}).String()
		s.CookieCache[request.Username] = cookie
		return &authpb.LoginAuthResponse{Cookie: cookie}, nil
	} else {
		return &authpb.LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}
}

func (s *StubService) Verify(ctx context.Context, request *authpb.VerifyAuthRequest) (*authpb.VerifyAuthResponse, error) {
	// check if cookie is assigned to a username
	for username, cookie := range s.CookieCache {
		savedCookie := DecodeCookie(cookie)
		requestCookie := DecodeCookie(request.Cookie)
		if savedCookie.Name == requestCookie.Name && savedCookie.Path == requestCookie.Path && savedCookie.Value == requestCookie.Value {
			response := &authpb.VerifyAuthResponse{Username: username, UID: savedCookie.Value}
			return response, nil
		}
	}

	// cookie is not known to us
	return &authpb.VerifyAuthResponse{}, &VerifyAuthError{request.Cookie}
}

func (s *StubService) Logout(ctx context.Context, request *authpb.LogoutAuthRequest) (*authpb.LogoutAuthResponse, error) {
	// view the user in user service
	viewUserRequest := &userpb.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(ctx, viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return &authpb.LogoutAuthResponse{}, &LogoutAuthError{request.Username}
	}

	// logout the user, their current status is irrelevant
	s.StatusCache[viewUserResponse.UID] = LOGGED_OUT
	delete(s.CookieCache, request.Username)
	return &authpb.LogoutAuthResponse{}, nil
}
