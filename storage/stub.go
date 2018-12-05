package storage

type StubStorage struct {
	storage map[string]string
}

func CreateStubStorage() *StubStorage {
	s := &StubStorage{}
	s.storage = map[string]string{}
	return s
}

func (s *StubStorage) Get(key string) (string, error) {
	if value, ok := s.storage[key]; ok {
		return value, nil
	}
	return "", &InvalidKeyError{Key: key}
}

func (s *StubStorage) Put(key string, value string) error {
	s.storage[key] = value
	return nil
}

func (s *StubStorage) Iterate() map[string]string { //(result []string) {
	//for k := range s.storage {
	//	result = append(result, k)
	//}
	//return result
	return s.storage
}
