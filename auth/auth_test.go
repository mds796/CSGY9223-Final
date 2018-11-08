package auth

import (
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func TestAuthRegisterBasic(t *testing.T) {
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

func TestAuthLoginBasic(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	registerRequest := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(registerRequest)

	logoutRequest := LogoutAuthRequest{Username: "mksavic"}
	authService.Logout(logoutRequest)

	request := LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := authService.Login(request)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginAlreadyLoggedIn(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	registerRequest := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(registerRequest)

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

	registerRequest := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	authService.Register(registerRequest)

	request := LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := authService.Login(request)

	if err == nil {
		t.Fail()
	}
}

func TestAuthVerifyBasic(t *testing.T) {
	userService := user.CreateStub()
	authService := CreateStub(userService)

	registerRequest := RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	registerResponse, _ := authService.Register(registerRequest)

	verifyRequest := VerifyAuthRequest{Cookie: registerResponse.Cookie}
	verifyResponse, err := authService.Verify(verifyRequest)

	if err != nil {
		t.Fail()
	}

	if verifyResponse.Username != "mksavic" {
		t.Fail()
	}
}
