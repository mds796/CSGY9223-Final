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
	Namespace string
}

func CreateRaftStorage(config StorageConfig, ns string) *RaftStorage {
	clients := []raftkvpb.RaftKVClient{}
	for _, host := range config.Hosts {
		client, err := raftkv.NewClient(host)
		if err != nil {
			log.Printf("[RAFTKV] Could not connect to Raft node '%v' when creating a cluster", err)
		} else {
			clients = append(clients, client)
		}
	}

	if len(clients) == 0 {
		panic("[RAFT] Could not connect to any Raft node when creating a cluster")
	}

	return &RaftStorage{
		Clients:   clients,
		Leader:    clients[0],
		Namespace: ns + "/",
	}
}

func (s *RaftStorage) Get(key string) ([]byte, error) {
	response, err := s.Leader.Get(
		s.context(),
		&raftkvpb.GetRequest{Key: s.keyWithNamespace(key)},
	)

	if err == nil {
		if len(response.Value) == 0 {
			return response.Value, &InvalidKeyError{Key: key}
		} else {
			return response.Value, nil
		}
	}

	// Retry in other cluster followers in case the leader has changed
	if err != nil {
		for _, client := range s.Clients {
			response, err := client.Get(
				s.context(),
				&raftkvpb.GetRequest{Key: s.keyWithNamespace(key)},
			)

			if err == nil {
				s.Leader = client
				log.Printf("[RAFT] Updated Raft leader %v", s.Leader)
				return response.Value, nil
			}
		}
	}

	return []byte{}, err
}

func (s *RaftStorage) Put(key string, value []byte) error {
	_, err := s.Leader.Put(
		s.context(),
		&raftkvpb.PutRequest{Key: s.keyWithNamespace(key), Value: value},
	)

	// Retry in other cluster followers in case the leader has changed
	if err != nil {
		for _, client := range s.Clients {
			_, err := client.Put(
				s.context(),
				&raftkvpb.PutRequest{Key: s.keyWithNamespace(key), Value: value},
			)

			if err == nil {
				s.Leader = client
				log.Printf("[RAFT] Updated Raft leader %v", s.Leader)
				return nil
			}
		}
	}

	return err
}

func (s *RaftStorage) Delete(key string) error {
	_, err := s.Leader.Delete(
		s.context(),
		&raftkvpb.DeleteRequest{Key: s.keyWithNamespace(key)},
	)

	// Retry in other cluster followers in case the leader has changed
	if err != nil {
		for _, client := range s.Clients {
			_, err := client.Delete(
				s.context(),
				&raftkvpb.DeleteRequest{Key: s.keyWithNamespace(key)},
			)

			if err == nil {
				s.Leader = client
				log.Printf("[RAFT] Updated Raft leader %v", s.Leader)
				return nil
			}
		}
	}

	return err
}

func (s *RaftStorage) Iterate() map[string][]byte {
	response, err := s.Leader.Iterate(
		s.context(),
		&raftkvpb.IterateRequest{Namespace: s.Namespace},
	)
	result := map[string][]byte{}
	if err == nil {
		for k, v := range response.KV.KV {
			result[s.keyWithoutNamespace(k)] = v
		}
		return result
	}

	// Retry in other cluster followers in case the leader has changed
	if err != nil {
		for _, client := range s.Clients {
			response, err := client.Iterate(
				s.context(),
				&raftkvpb.IterateRequest{Namespace: s.Namespace},
			)

			if err == nil {
				s.Leader = client
				log.Printf("[RAFT] Updated Raft leader %v", s.Leader)

				for k, v := range response.KV.KV {
					result[s.keyWithoutNamespace(k)] = v
				}
				return result
			}
		}
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
