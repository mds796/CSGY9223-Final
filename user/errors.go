package user

import (
	"fmt"
)

type CreateUserError struct {
	Username string
}

type ViewUserError struct {
	Username string
}

func (e *CreateUserError) Error() string {
	return fmt.Sprintf("[USER]: Invalid username %s.", e.Username)
}

func (e *ViewUserError) Error() string {
	return fmt.Sprintf("[USER]: User %s does not exist.", e.Username)
}
