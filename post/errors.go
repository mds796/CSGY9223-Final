package post

import (
	"fmt"
)

type EmptyPostTextError struct {
	Text string
}

type InvalidPostIDError struct {
	PostID string
}

func (e *EmptyPostTextError) Error() string {
	return fmt.Sprintf("[POST]: Invalid post text \"%s\".", e.Text)
}

func (e *InvalidPostIDError) Error() string {
	return fmt.Sprintf("[POST]: Invalid post ID %s.", e.PostID)
}
