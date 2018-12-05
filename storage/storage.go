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

func CreateStorage(storageType StorageType) Storage {
	if storageType == RAFT {
		return CreateRaftStorage()
	} else {
		return CreateStubStorage()
	}
}
