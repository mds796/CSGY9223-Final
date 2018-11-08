package auth

import (
	"net/http"
)

type RegisterAuthRequest struct {
	Username string
	Password string
}

type RegisterAuthResponse struct {
	Cookie http.Cookie
}

type LoginAuthRequest struct {
	Username string
	Password string
}

type LoginAuthResponse struct {
	Cookie http.Cookie
}

type VerifyAuthRequest struct {
	Cookie http.Cookie
}

type VerifyAuthResponse struct {
	Username string
}

type LogoutAuthRequest struct {
	Username string
}

type LogoutAuthResponse struct {
}

type Service interface {
	Register(request RegisterAuthRequest) (RegisterAuthResponse, error)
	Login(request LoginAuthRequest) (LoginAuthResponse, error)
	Verify(request VerifyAuthRequest) (VerifyAuthResponse, error)
	Logout(request LogoutAuthRequest) (LogoutAuthResponse, error)
}
