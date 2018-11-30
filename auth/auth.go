package auth

import (
	"github.com/mds796/CSGY9223-Final/auth/authpb"
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
	authpb.RegisterAuthServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Auth now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	userService := user.NewStubClient(user.NewStubServer())
	return &RpcService{config: config, service: NewStubServer(userService)}
}

func NewStubServer(userService userpb.UserClient) *StubService {
	return &StubService{UserService: userService}
}

func NewClient(target string) (authpb.AuthClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return authpb.NewAuthClient(conn), nil

}

func NewStubClient(server authpb.AuthServer) *StubClient {
	return &StubClient{service: server}
}
