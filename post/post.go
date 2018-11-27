package post

import (
	"github.com/mds796/CSGY9223-Final/post/postpb"
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
	postpb.RegisterPostServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Post now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	return &RpcService{config: config, service: NewStubServer()}
}

func NewStubServer() *StubService {
	return &StubService{}
}

func NewClient(target string) (postpb.PostClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return postpb.NewPostClient(conn), nil

}

func NewStubClient(server postpb.PostServer) *StubClient {
	return &StubClient{service: server}
}
