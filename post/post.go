package post

import (
	"github.com/google/uuid"
)

type Post struct {
	postID uuid.UUID
	userID uuid.UUID
	text   string
}
