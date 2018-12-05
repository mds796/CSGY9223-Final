package storage

type RaftStorage struct {
}

func CreateRaftStorage() *RaftStorage {
	s := &RaftStorage{}
	return s
}

func (s *RaftStorage) Get(key string) (string, error) {
	return "", nil
}

func (s *RaftStorage) Put(key string, value string) error {
	return nil
}

func (s *RaftStorage) Iterate() map[string]string { //[]string {
	//return []string{}
	return map[string]string{}
}
