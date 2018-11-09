package user

import (
	"github.com/google/uuid"
	"strings"
)

const MIN_USERNAME = 6

type StubService struct {
	UuidCache     map[string]string // (UUID, username)
	UsernameCache map[string]string // (username, UUID)
}

func CreateStub() Service {
	stub := new(StubService)
	stub.UuidCache = make(map[string]string)
	stub.UsernameCache = make(map[string]string)
	return stub
}

func (s *StubService) Create(request CreateUserRequest) (CreateUserResponse, error) {
	// ensure this username meets the minimum requirements
	if len(request.Username) < MIN_USERNAME {
		return CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// ensure this username doesn't already exist
	if _, ok := s.UsernameCache[request.Username]; ok {
		return CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// generate the uuid
	newUuid := uuid.New().String()

	// add the user
	s.UuidCache[newUuid] = request.Username
	s.UsernameCache[request.Username] = newUuid

	// create the response
	response := CreateUserResponse{Uuid: newUuid}
	return response, nil
}

func (s *StubService) View(request ViewUserRequest) (ViewUserResponse, error) {
	if uuid, ok := s.UsernameCache[request.Username]; ok {
		// username exists
		return ViewUserResponse{Uuid: uuid}, nil
	} else {
		// username doesn't exist
		return ViewUserResponse{}, &ViewUserError{request.Username}
	}
}

func (s *StubService) Search(request SearchUserRequest) (SearchUserResponse, error) {
	// find uuids that match given query
	var usernames []string
	for username, _ := range s.UsernameCache {
		if strings.Contains(username, request.Query) {
			usernames = append(usernames, username)
		}
	}

	// create the response
	response := SearchUserResponse{Usernames: usernames}
	return response, nil
}
