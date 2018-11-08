package auth

import (
	"bytes"
	"crypto/sha256"
	"github.com/mds796/CSGY9223-Final/user"
)

const (
	LOGGED_OUT = iota // LOGGED_OUT == 0
	LOGGED_IN  = iota // LOGGED_IN == 1
)

type StubService struct {
	UserService   user.Service
	PasswordCache map[string][]byte
	StatusCache   map[string]int
}

func CreateStub(userService user.Service) Service {
	stub := new(StubService)
	stub.UserService = userService
	stub.PasswordCache = make(map[string][]byte)
	stub.StatusCache = make(map[string]int)
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
	return RegisterAuthResponse{}, nil
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
		return LoginAuthResponse{}, nil
	} else {
		return LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}
}

func (s *StubService) Verify(request VerifyAuthRequest) (VerifyAuthResponse, error) {
	return VerifyAuthResponse{}, nil
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
	return LogoutAuthResponse{}, nil
}
