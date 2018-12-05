package auth

import (
	"context"
	"fmt"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func createAuthService() *StubClient {
	userService := user.NewStubClient(user.CreateStub(storage.STUB))
	return &StubClient{service: CreateStub(storage.STUB, userService)}
}

func sendRegisterAuthRequest(client *StubClient, request *authpb.RegisterAuthRequest) (*authpb.RegisterAuthResponse, error) {
	return client.Register(context.Background(), request)
}

func sendLoginAuthRequest(client *StubClient, request *authpb.LoginAuthRequest) (*authpb.LoginAuthResponse, error) {
	return client.Login(context.Background(), request)
}

func sendVerifyAuthRequest(client *StubClient, request *authpb.VerifyAuthRequest) (*authpb.VerifyAuthResponse, error) {
	return client.Verify(context.Background(), request)
}

func sendLogoutAuthRequest(client *StubClient, request *authpb.LogoutAuthRequest) (*authpb.LogoutAuthResponse, error) {
	return client.Logout(context.Background(), request)
}

func TestAuthRegisterBasic(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendRegisterAuthRequest(authService, registerRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthRegisterExists(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	_, err := sendRegisterAuthRequest(authService, registerRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginBasic(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(authService, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(authService, loginRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginDoesNotExist(t *testing.T) {
	authService := createAuthService()

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(authService, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginPasswordIncorrect(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(authService, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := sendLoginAuthRequest(authService, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginAlreadyLoggedIn(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(authService, loginRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginAlreadyLoggedInIncorrectPassword(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := sendLoginAuthRequest(authService, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthVerifyBasic(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	registerResponse, _ := sendRegisterAuthRequest(authService, registerRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: registerResponse.Cookie}
	verifyResponse, err := sendVerifyAuthRequest(authService, verifyRequest)

	if err != nil {
		t.Fail()
	}

	if verifyResponse.Username != "mksavic" {
		t.Fail()
	}

	registerResponseCookie := DecodeCookie(registerResponse.Cookie)
	if verifyResponse.UID != registerResponseCookie.Value {
		t.Fail()
	}
}

func TestAuthVerifyLoggedOut(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	registerResponse, _ := sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(authService, logoutRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: registerResponse.Cookie}
	_, err := sendVerifyAuthRequest(authService, verifyRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthVerifyLogOutLogIn(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(authService, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	loginResponse, _ := sendLoginAuthRequest(authService, loginRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: loginResponse.Cookie}
	verifyResponse, err := sendVerifyAuthRequest(authService, verifyRequest)

	if err != nil {
		t.Fail()
	}

	if verifyResponse.Username != "mksavic" {
		t.Fail()
	}

	loginResponseCookie := DecodeCookie(loginResponse.Cookie)
	if verifyResponse.UID != loginResponseCookie.Value {
		fmt.Println(verifyResponse.UID)
		fmt.Println(loginResponseCookie.Value)
		t.Fail()
	}
}

func TestAuthLogoutBasic(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	_, err := sendLogoutAuthRequest(authService, logoutRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLogoutDoesNotExist(t *testing.T) {
	authService := createAuthService()

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	_, err := sendLogoutAuthRequest(authService, logoutRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLogoutAlreadyLoggedOut(t *testing.T) {
	authService := createAuthService()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(authService, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(authService, logoutRequest)

	_, err := sendLogoutAuthRequest(authService, logoutRequest)

	if err != nil {
		t.Fail()
	}
}
