package auth

import (
	"fmt"
)

type RegisterAuthError struct {
	Username string
}

type LoginAuthError struct {
	Username string
	Password string
}

type LogoutAuthError struct {
	Username string
}

func (e *RegisterAuthError) Error() string {
	return fmt.Sprintf("[AUTH]: Username %s already exists.", e.Username)
}

func (e *LoginAuthError) Error() string {
	return fmt.Sprintf("[AUTH]: Invalid login request for %s:%s.", e.Username, e.Password)
}

func (e *LogoutAuthError) Error() string {
	return fmt.Sprintf("[AUTH]: Invalid logout request for %s.", e.Username)
}
