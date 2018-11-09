package post

import (
	"github.com/google/uuid"
)

type Post struct {
	PostID uuid.UUID
	User   string
	Text   string
}
