package post

import (
	"github.com/google/uuid"
)

type Post struct {
	postID uuid.UUID
	User   string
	Text   string
}
