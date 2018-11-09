package feed

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
)

type ViewRequest struct {
	Username string
}

type ViewResponse struct {
	Posts []post.Post
}

type Service interface {
	View(request ViewRequest) (ViewResponse, error)
}

func CreateStub(postService post.Service, followService follow.Service) Service {
	return &StubService{Post: postService, Follow: followService}
}
