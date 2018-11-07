package auth

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

type StubService struct {
	PasswordCache map[string][]byte
	StatusCache   map[string]int
}

func CreateStub() Service {
	stub := new(StubService)
	stub.PasswordCache = make(map[string][]byte)
	stub.StatusCache = make(map[string]int)
	return stub
}

func (s *StubService) Register(request RegisterAuthRequest) (RegisterAuthResponse, error) {
	if _, ok := s.PasswordCache[request.Uuid]; ok {
		// user is already registered
		err := errors.New("[AUTH]: Uuid already exists.")
		return RegisterAuthResponse{}, err
	} else {
		// hash the password and save it
		h := sha256.New()
		h.Write([]byte(request.Password))
		s.PasswordCache[request.Uuid] = h.Sum(nil)
		return RegisterAuthResponse{}, nil
	}
}

func (s *StubService) Login(request LoginAuthRequest) (LoginAuthResponse, error) {
	if _, ok := s.PasswordCache[request.Uuid]; !ok {
		// user is not registered
		err := errors.New("[AUTH]: Uuid does not exist.")
		return LoginAuthResponse{}, err
	}

	h := sha256.New()
	h.Write([]byte(request.Password))
	if bytes.Equal(s.PasswordCache[request.Uuid], h.Sum(nil)) {
		// login the user, their current status is irrelevant
		s.StatusCache[request.Uuid] = 1
		return LoginAuthResponse{}, nil
	} else {
		err := errors.New("[AUTH]: Password is incorrect.")
		return LoginAuthResponse{}, err
	}
}

func (s *StubService) Verify(request VerifyAuthRequest) (VerifyAuthResponse, error) {
	return VerifyAuthResponse{}, nil
}

func (s *StubService) Logout(request LogoutAuthRequest) (LogoutAuthResponse, error) {
	// logout the user, their current status is irrelevant
	s.StatusCache[request.Uuid] = 0
	return LogoutAuthResponse{}, nil
}
