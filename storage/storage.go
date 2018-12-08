package storage

type StorageType int

const (
	STUB StorageType = iota
	RAFT StorageType = iota
)

type StorageConfig struct {
	StorageType StorageType
	Hosts       []string
	Namespace   string
}

type Storage interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Iterate() map[string][]byte
}

// storageType: chooses the Raft or Stub implementation of Storage
// namespace: provides logical separation for services in a key-value data store
func CreateStorage(config StorageConfig, namespace string) Storage {
	switch config.StorageType {
	case RAFT:
		return CreateRaftStorage(config, namespace)
	default:
		return CreateStubStorage()
	}
}
