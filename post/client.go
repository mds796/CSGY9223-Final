package post

type CreatePostRequest struct {
	UserID string
	Text   string
}

type CreatePostResponse struct {
	PostID string
}

type CreateViewRequest struct {
	PostID string
}

type CreateViewResponse struct {
	Text string
}

type Service interface {
	Create(request CreatePostRequest) (CreatePostResponse, error)
	View(request CreateViewRequest) (CreateViewResponse, error)
}
