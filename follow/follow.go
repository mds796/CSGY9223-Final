// TODO: Remove duplicated code (feed, follow and post re-implement this)

package follow

import (
	"github.com/mds796/CSGY9223-Final/follow/followpb"
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
	followpb.RegisterFollowServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("Follow now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	userService, err := user.NewClient(config.UserTarget())
	if err != nil {
		log.Fatal(err)
	}

	return &RpcService{config: config, service: CreateService(storage.RAFT, userService)}
}

func NewStubServer(userService userpb.UserClient) *Service {
	return CreateService(storage.STUB, userService)
}

func NewClient(target string) (followpb.FollowClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return followpb.NewFollowClient(conn), nil

}

func NewStubClient(userService userpb.UserClient) *StubClient {
	return &StubClient{service: NewStubServer(userService)}
}
