package auth

import (
	"crypto/sha256"
)

type StubService struct {
	UsersCache  map[string]string
	StatusCache map[string]string
}

func CreateStub() Service {
	stub := new(StubService)
	stub.UsersCache = make(map[string]string)
	stub.StatusCache = make(map[string]int)
	return stub
}

func (s *StubService) Register(request RegisterAuthRequest) (RegisterAuthResponse, error) {
}

func (s *StubService) Login(request LoginAuthRequest) (LoginAuthResponse, error) {
}

func (s *StubService) Verify(request VerifyAuthRequest) (VerifyAuthResponse, error) {
}

func (s *StubService) Logout(request LogoutAuthRequest) (LogoutAuthResponse, error) {
}
