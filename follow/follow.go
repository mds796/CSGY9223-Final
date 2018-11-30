// TODO: Remove duplicated code (feed, follow and post re-implement this)

package follow

import (
	"github.com/mds796/CSGY9223-Final/follow/followpb"
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
	followpb.RegisterFollowServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Follow now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	userService := user.NewStubClient(user.CreateStub())
	return &RpcService{config: config, service: NewStubServer(userService)}
}

func NewStubServer(userService userpb.UserClient) *StubService {
	return &StubService{User: userService}
}

func NewClient(target string) (followpb.FollowClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return followpb.NewFollowClient(conn), nil

}

func NewStubClient(server followpb.FollowServer) *StubClient {
	return &StubClient{service: server}
}
