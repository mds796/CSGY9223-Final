package storage

type StorageType int

const (
	RAFT StorageType = iota // RAFT == 0
	STUB StorageType = iota // STUB == 1
)

type Storage interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Iterate() map[string][]byte
}

// storageType: chooses the Raft or Stub implementation of Storage
// namespace: provides logical separation for services in a key-value data store
func CreateStorage(storageType StorageType, namespace string) Storage {
	if storageType == RAFT {
		return CreateRaftStorage(namespace)
	} else {
		return CreateStubStorage(namespace)
	}
}
