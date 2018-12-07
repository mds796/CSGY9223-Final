package storage

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type RaftStorage struct {
	Client    *clientv3.Client
	Namespace string
}

func CreateRaftStorage(ns string) *RaftStorage {
	client, err := clientv3.New(getRaftClusterConfig())
	if err != nil {
		// handle error!
	}
	// defer client.Close()
	return &RaftStorage{Client: client, Namespace: ns + "/"}
}

func getRaftClusterConfig() clientv3.Config {
	return clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	}
}

func (s *RaftStorage) Get(key string) ([]byte, error) {
	response, err := s.Client.Get(s.context(), s.keyWithNamespace(key))
	value := []byte{}
	if err != nil {
		return value, &GetError{Key: s.keyWithNamespace(key)}
	}

	if response.Count < 1 {
		return value, &InvalidKeyError{Key: s.keyWithNamespace(key)}
	}

	return response.Kvs[0].Value, err
}

func (s *RaftStorage) Put(key string, value []byte) error {
	_, err := s.Client.Put(s.context(), s.keyWithNamespace(key), string(value))
	if err != nil {
		return &PutError{Key: s.keyWithNamespace(key)}
	}
	return err
}

func (s *RaftStorage) Delete(key string) error {
	response, err := s.Client.Delete(s.context(), s.keyWithNamespace(key))
	if err != nil {
		return &DeleteError{Key: s.keyWithNamespace(key)}
	}

	if response.Deleted < 1 {
		return &InvalidKeyError{Key: s.keyWithNamespace(key)}
	}

	return nil
}

func (s *RaftStorage) Iterate() map[string][]byte {
	response, err := s.Client.Get(s.context(), s.Namespace, clientv3.WithPrefix())
	result := map[string][]byte{}
	if err == nil {
		log.Println(response.Kvs)
		for _, kv := range response.Kvs {
			result[s.keyWithoutNamespace(string(kv.Key))] = kv.Value
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
