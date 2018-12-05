package storage

type RaftStorage struct {
}

func CreateRaftStorage() *RaftStorage {
	s := &RaftStorage{}
	return s
}

func (s *RaftStorage) Get(key string) ([]byte, error) {
	return []byte(""), nil
}

func (s *RaftStorage) Put(key string, value []byte) error {
	return nil
}

func (s *RaftStorage) Iterate() map[string][]byte {
	return map[string][]byte{}
}
