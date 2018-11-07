package user

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type StubService struct {
	UsersCache map[string]string
}

func CreateStub() *StubService {
	stub := new(StubService)
	stub.UsersCache = make(map[string]string)
	return stub
}

func (s *StubService) Create(request CreateUserRequest) (CreateUserResponse, error) {
	// ensure this username doesn't already exist
	for _, username := range s.UsersCache {
		if username == request.Username {
			err := errors.New("[USER]: Username already exists.")
			return CreateUserResponse{}, err
		}
	}

	// create the User object
	user := User{uuid: uuid.New().String(), username: request.Username}

	// add the user
	s.UsersCache[user.uuid] = user.username

	// create the response
	response := CreateUserResponse{Uuid: user.uuid}

	return response, nil
}

func (s *StubService) View(request ViewUserRequest) (ViewUserResponse, error) {
	if username, ok := s.UsersCache[request.Uuid]; ok {
		// uuid exists
		response := ViewUserResponse{Username: username}
		return response, nil
	} else {
		// uuid doesn't exist
		err := errors.New("[USER]: UUID does not exist.")
		return ViewUserResponse{}, err
	}
}

func (s *StubService) Search(request SearchUserRequest) (SearchUserResponse, error) {
	// find uuids that match given query
	var uuids []string
	for uuid, username := range s.UsersCache {
		if strings.Contains(username, request.Query) {
			uuids = append(uuids, uuid)
		}
	}

	// create the response
	response := SearchUserResponse{Uuids: uuids}

	return response, nil
}
