package auth

import (
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func TestAuthRegisterStandard(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	request := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := authService.Register(request)

	if err != nil {
		t.Fail()
	}
}

func TestAuthRegisterExists(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	request := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(request)
	_, err := authService.Register(request)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginStandard(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	register_request := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(register_request)

	request := LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := authService.Login(request)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginDoesNotExist(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	request := LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := authService.Login(request)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginPasswordIncorrect(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	register_request := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(register_request)

	request := LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := authService.Login(request)

	if err == nil {
		t.Fail()
	}
}
