package raftkv

import (
	"context"
	"github.com/mds796/CSGY9223-Final/storage/raftkv/raftkvpb"
	"log"
)

type Service struct {
	Raft *RaftKV
}

func CreateService(nodeID string, address string, standaloneCluster bool) *Service {
	service := new(Service)
	service.Raft = CreateRaftKV(nodeID, address)
	err := service.Raft.Open(standaloneCluster)
	if err != nil {
		log.Fatalf("Error creating Raft node: %v", err)
		panic("Cannot create Raft node")
	}
	return service
}

func (s *Service) Get(ctx context.Context, request *raftkvpb.GetRequest) (*raftkvpb.GetResponse, error) {
	value, err := s.Raft.Get(request.Key)
	return &raftkvpb.GetResponse{Key: request.Key, Value: value}, err
}

func (s *Service) Put(ctx context.Context, request *raftkvpb.PutRequest) (*raftkvpb.PutResponse, error) {
	err := s.Raft.Put(request.Key, request.Value)
	if err != nil {
		switch err.(type) {
		case *NotLeaderError:
			return &raftkvpb.PutResponse{IsLeader: false}, err
		}
	}
	return &raftkvpb.PutResponse{IsLeader: true}, err
}

func (s *Service) Delete(ctx context.Context, request *raftkvpb.DeleteRequest) (*raftkvpb.DeleteResponse, error) {
	err := s.Raft.Delete(request.Key)
	return &raftkvpb.DeleteResponse{}, err
}

func (s *Service) Iterate(ctx context.Context, request *raftkvpb.IterateRequest) (*raftkvpb.IterateResponse, error) {
	kv := s.Raft.Iterate(request.Namespace)
	return &raftkvpb.IterateResponse{KV: &raftkvpb.KeyValue{KV: kv}}, nil
}

func (s *Service) Join(ctx context.Context, request *raftkvpb.JoinRequest) (*raftkvpb.JoinResponse, error) {
	err := s.Raft.Join(request.NodeID, request.Address)
	return &raftkvpb.JoinResponse{}, err
}
