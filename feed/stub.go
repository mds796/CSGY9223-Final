package feed

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
)

type StubService struct {
	Post   post.Service
	Follow follow.Service
}

func (s StubService) View(request ViewRequest) (ViewResponse, error) {
}
