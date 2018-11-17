package feed

import (
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/user"
	"google.golang.org/grpc"
)

func NewStubServer(postService post.Service, userService user.Service, followService follow.Service) *StubService {
	return &StubService{Post: postService, Follow: followService, User: userService}

}

func NewClient(target string) (feedpb.FeedClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return feedpb.NewFeedClient(conn), nil

}

func NewStubClient(server feedpb.FeedServer) *StubClient {
	return &StubClient{service: server}

}
