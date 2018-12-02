package feed

import (
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcService struct {
	config  *Config
	service *StubService
}

func (s *RpcService) Start() error {
	srv := grpc.NewServer()
	feedpb.RegisterFeedServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Feed now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	userService, err := user.NewClient(config.UserTarget())
	if err != nil {
		log.Fatal(err)
	}

	postService, err := post.NewClient(config.PostTarget())
	if err != nil {
		log.Fatal(err)
	}

	followService, err := follow.NewClient(config.FollowTarget())
	if err != nil {
		log.Fatal(err)
	}

	return &RpcService{config: config, service: NewStubServer(postService, userService, followService)}
}

func NewStubServer(postService postpb.PostClient, userService userpb.UserClient, followService followpb.FollowClient) *StubService {
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
