package user

import (
	"fmt"
)

type CreateUserError struct {
	Username string
}

type ViewUserError struct {
	Uuid string
}

func (e *CreateUserError) Error() string {
	return fmt.Sprintf("[USER]: Username %s already exists.", e.Username)
}

func (e *ViewUserError) Error() string {
	return fmt.Sprintf("[USER]: UUID %s does not exist.", e.Uuid)
}
