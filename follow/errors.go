package follow

import (
	"fmt"
)

type InvalidUserIDError struct {
	UserID string
}

func (e *InvalidUserIDError) Error() string {
	return fmt.Sprintf("[FOLLOW]: Invalid user ID %s.", e.UserID)
}
