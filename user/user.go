package user

import (
	"github.com/mds796/CSGY9223-Final/storage"
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
	userpb.RegisterUserServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("User now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	return &RpcService{config: config, service: CreateStub(storage.RAFT)}
}

func NewStubServer() *Service {
	return CreateStub(storage.STUB)
}

func NewClient(target string) (userpb.UserClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return userpb.NewUserClient(conn), nil

}

func NewStubClient() *StubClient {
	return &StubClient{service: NewStubServer()}
}
