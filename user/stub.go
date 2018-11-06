package user

import "github.com/google/uuid"

type StubService struct {
	Cache map[string]string
}

func CreateStub() Service {
	stub := new(StubService)
	stub.Cache = make(map[string]string)

	return stub
}
func (s *StubService) Create(request CreateUserRequest) (CreateUserResponse, error) {
	response := CreateUserResponse{UserId: uuid.New().String()}

	s.Cache[response.UserId] = User{username: request.Alias, id: response.UserId}

	return response, nil
}
