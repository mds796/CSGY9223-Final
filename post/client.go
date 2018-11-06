package post

type CreatePostRequest struct {
	UserID string
	Text   string
}

type CreatePostResponse struct {
	PostID string
}

type Service interface {
	Create(request CreatePostRequest) (CreatePostResponse, error)
}
