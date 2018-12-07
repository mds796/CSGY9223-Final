package auth

import (
	"context"
	"fmt"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func createClient() *StubClient {
	return NewStubClient(user.NewStubClient())
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
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendRegisterAuthRequest(client, registerRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthRegisterExists(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	_, err := sendRegisterAuthRequest(client, registerRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginBasic(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(client, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(client, loginRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginDoesNotExist(t *testing.T) {
	client := createClient()

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(client, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginPasswordIncorrect(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(client, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := sendLoginAuthRequest(client, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginAlreadyLoggedIn(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	_, err := sendLoginAuthRequest(client, loginRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginAlreadyLoggedInIncorrectPassword(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "123abc"}
	_, err := sendLoginAuthRequest(client, loginRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthVerifyBasic(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	registerResponse, _ := sendRegisterAuthRequest(client, registerRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: registerResponse.Cookie}
	verifyResponse, err := sendVerifyAuthRequest(client, verifyRequest)

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
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	registerResponse, _ := sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(client, logoutRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: registerResponse.Cookie}
	_, err := sendVerifyAuthRequest(client, verifyRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthVerifyLogOutLogIn(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(client, logoutRequest)

	loginRequest := &authpb.LoginAuthRequest{Username: "mksavic", Password: "abc123"}
	loginResponse, _ := sendLoginAuthRequest(client, loginRequest)

	verifyRequest := &authpb.VerifyAuthRequest{Cookie: loginResponse.Cookie}
	verifyResponse, err := sendVerifyAuthRequest(client, verifyRequest)

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
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	_, err := sendLogoutAuthRequest(client, logoutRequest)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLogoutDoesNotExist(t *testing.T) {
	client := createClient()

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	_, err := sendLogoutAuthRequest(client, logoutRequest)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLogoutAlreadyLoggedOut(t *testing.T) {
	client := createClient()

	registerRequest := &authpb.RegisterAuthRequest{Username: "mksavic", Password: "abc123"}
	sendRegisterAuthRequest(client, registerRequest)

	logoutRequest := &authpb.LogoutAuthRequest{Username: "mksavic"}
	sendLogoutAuthRequest(client, logoutRequest)

	_, err := sendLogoutAuthRequest(client, logoutRequest)

	if err != nil {
		t.Fail()
	}
}
