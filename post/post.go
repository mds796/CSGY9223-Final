package post

import (
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcService struct {
	config  *Config
	service *Service
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
	return &RpcService{
		config:  config,
		service: CreateService(storage.StorageConfig{StorageType: storage.RAFT, Hosts: config.StorageHosts}),
	}
}

func NewStubServer() *Service {
	return CreateService(storage.StorageConfig{StorageType: storage.STUB})
}

func NewClient(target string) (postpb.PostClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return postpb.NewPostClient(conn), nil

}

func NewStubClient() *StubClient {
	return &StubClient{service: NewStubServer()}
}
