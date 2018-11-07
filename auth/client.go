package auth

type RegisterAuthRequest struct {
	Username string
	Password string
}

type RegisterAuthResponse struct {
	Uuid string
}

type LoginAuthRequest struct {
	Username string
	Password string
}

type LoginAuthResponse struct {
	Uuid string
}

type VerifyAuthRequest struct {
	Cookie string
}

type VerifyAuthResponse struct {
	Uuid string
}

type LogoutAuthRequest struct {
	Uuid string
}

type LogoutAuthResponse struct {
}

type Service interface {
	Register(request RegisterAuthRequest) (RegisterAuthResponse, error)
	Login(request LoginAuthRequest) (LoginAuthResponse, error)
	Verify(request VerifyAuthRequest) (VerifyAuthResponse, error)
	Logout(request LogoutAuthRequest) (LogoutAuthResponse, error)
}
