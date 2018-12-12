package storage

type StubStorage struct {
	storage map[string][]byte
}

func CreateStubStorage() *StubStorage {
	s := &StubStorage{}
	s.storage = map[string][]byte{}
	return s
}

func (s *StubStorage) Get(key string) ([]byte, error) {
	if value, ok := s.storage[key]; ok {
		return value, nil
	}
	return []byte(""), &InvalidKeyError{Key: key}
}

func (s *StubStorage) Put(key string, value []byte) error {
	s.storage[key] = value
	return nil
}

func (s *StubStorage) Delete(key string) error {
	if _, ok := s.storage[key]; ok {
		delete(s.storage, key)
		return nil
	}
	return &InvalidKeyError{Key: key}
}

func (s *StubStorage) Iterate() map[string][]byte {
	return s.storage
}
