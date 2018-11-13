package post

import "time"

type CreatePostRequest struct {
	UserID string
	Text   string
}

type CreatePostResponse struct {
	PostID string
}

type ViewPostRequest struct {
	PostID string
}

type ViewPostResponse struct {
	Text      string
	Timestamp time.Time
}

type ListPostsRequest struct {
	UserID string
}

type ListPostsResponse struct {
	PostIDs []string
}

type Service interface {
	Create(request CreatePostRequest) (CreatePostResponse, error)
	View(request ViewPostRequest) (ViewPostResponse, error)
	List(request ListPostsRequest) (ListPostsResponse, error)
}
