package post

import "time"

type Post struct {
	PostID    string
	User      string
	Text      string
	Timestamp time.Time
}
