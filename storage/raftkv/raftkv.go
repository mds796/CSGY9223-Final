package raftkv

import (
	"context"
	"github.com/mds796/CSGY9223-Final/storage/raftkv/raftkvpb"
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
	raftkvpb.RegisterRaftKVServer(srv, s.service)

	lis, err := net.Listen("tcp", s.config.Target())
	if err != nil {
		return err
	}

	log.Printf("RaftKV now listening on %v.\n", s.config.Target())

	return srv.Serve(lis)
}

func New(config *Config) *RpcService {
	standaloneCluster := config.JoinHost == ""

	r := &RpcService{
		config:  config,
		service: CreateService(config.NodeID, config.RaftTarget(), standaloneCluster),
	}

	if standaloneCluster {
		log.Printf("RaftKV node %v started as standalone cluster", config.NodeID)
	} else {
		log.Printf("RaftKV node %v trying to join cluster at %v", config.NodeID, config.JoinTarget())
		c, err := NewClient(config.JoinTarget())
		if err != nil {
			panic("Could not open connection with RaftKV cluster")
		}
		c.Join(context.Background(), &raftkvpb.JoinRequest{NodeID: config.NodeID, Address: config.RaftTarget()})
	}

	return r
}

func NewClient(target string) (raftkvpb.RaftKVClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return raftkvpb.NewRaftKVClient(conn), nil

}
