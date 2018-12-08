package auth

import (
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/mds796/CSGY9223-Final/user/userpb"
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
	authpb.RegisterAuthServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Auth now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	userService, err := user.NewClient(config.UserTarget())
	if err != nil {
		log.Fatal(err)
	}

	return &RpcService{
		config:  config,
		service: CreateService(storage.StorageConfig{StorageType: storage.RAFT, Hosts: config.StorageHosts}, userService),
	}
}

func NewStubServer(userService userpb.UserClient) *Service {
	return CreateService(storage.StorageConfig{StorageType: storage.STUB}, userService)
}

func NewClient(target string) (authpb.AuthClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return authpb.NewAuthClient(conn), nil

}

func NewStubClient(userService userpb.UserClient) *StubClient {
	return &StubClient{service: NewStubServer(userService)}
}
