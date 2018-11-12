package feed

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/user"
)

type Post struct {
	From string
	Text string
}

type ViewRequest struct {
	UserID   string
	Username string
}

type ViewResponse struct {
	Posts []*Post
}

type Service interface {
	View(request *ViewRequest) (*ViewResponse, error)
}

func CreateStub(postService post.Service, userService user.Service, followService follow.Service) Service {
	return &StubService{Post: postService, Follow: followService, User: userService}
}
