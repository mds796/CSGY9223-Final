package post

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
	Text string
}

type Service interface {
	Create(request CreatePostRequest) (CreatePostResponse, error)
	View(request ViewPostRequest) (ViewPostResponse, error)
}
