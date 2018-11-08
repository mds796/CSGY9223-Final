package auth

type RegisterAuthRequest struct {
	Username string
	Password string
}

type RegisterAuthResponse struct {
}

type LoginAuthRequest struct {
	Username string
	Password string
}

type LoginAuthResponse struct {
}

type VerifyAuthRequest struct {
	Cookie string
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
