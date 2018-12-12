package storage

import (
	"context"
	"github.com/mds796/CSGY9223-Final/storage/raftkv"
	"github.com/mds796/CSGY9223-Final/storage/raftkv/raftkvpb"
	"log"
)

type RaftStorage struct {
	Clients   []raftkvpb.RaftKVClient
	Leader    raftkvpb.RaftKVClient
	Followers []raftkvpb.RaftKVClient
	Namespace string
}

func CreateRaftStorage(config StorageConfig, ns string) *RaftStorage {
	clients, err := raftkv.NewClusterClients(config.Hosts)
	if err != nil {
		log.Fatalf("[RAFT] Could not connect to Raft nodes: %v", err)
		panic(err)
	}

	return &RaftStorage{
		Clients:   clients,
		Leader:    clients[0],
		Followers: clients[1:],
		Namespace: ns + "/",
	}
}

func (s *RaftStorage) Get(key string) ([]byte, error) {
	response, err := s.Leader.Get(
		s.context(),
		&raftkvpb.GetRequest{Key: s.keyWithNamespace(key)},
	)

	if err != nil {
		return []byte{}, err
	}
	return response.Value, nil
}

func (s *RaftStorage) Put(key string, value []byte) error {
	_, err := s.Leader.Put(
		s.context(),
		&raftkvpb.PutRequest{Key: s.keyWithNamespace(key), Value: value},
	)
	return err
}

func (s *RaftStorage) Delete(key string) error {
	_, err := s.Leader.Delete(
		s.context(),
		&raftkvpb.DeleteRequest{Key: s.keyWithNamespace(key)},
	)
	return err
}

func (s *RaftStorage) Iterate() map[string][]byte {
	response, _ := s.Leader.Iterate(
		s.context(),
		&raftkvpb.IterateRequest{Namespace: s.Namespace},
	)
	result := map[string][]byte{}
	for k, v := range response.KV.KV {
		result[s.keyWithoutNamespace(k)] = v
	}
	return result
}

func (s *RaftStorage) context() context.Context {
	return context.Background()
}

func (s *RaftStorage) keyWithNamespace(key string) string {
	return s.Namespace + key
}

func (s *RaftStorage) keyWithoutNamespace(key string) string {
	return key[len(s.Namespace):]
}
