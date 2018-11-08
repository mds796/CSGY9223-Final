package user

import (
	"github.com/google/uuid"
	"strings"
)

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
	// ensure this username doesn't already exist
	if _, ok := s.UsernameCache[request.Username]; ok {
		return CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// generate the uuid
	new_uuid := uuid.New().String()

	// add the user
	s.UuidCache[new_uuid] = request.Username
	s.UsernameCache[request.Username] = new_uuid

	// create the response
	response := CreateUserResponse{Uuid: new_uuid}
	return response, nil
}

func (s *StubService) View(request ViewUserRequest) (ViewUserResponse, error) {
	if username, ok := s.UuidCache[request.Uuid]; ok {
		// uuid exists
		response := ViewUserResponse{Username: username}
		return response, nil
	} else {
		// uuid doesn't exist
		return ViewUserResponse{}, &ViewUserError{request.Uuid}
	}
}

func (s *StubService) Search(request SearchUserRequest) (SearchUserResponse, error) {
	// find uuids that match given query
	var uuids []string
	for uuid, username := range s.UuidCache {
		if strings.Contains(username, request.Query) {
			uuids = append(uuids, uuid)
		}
	}

	// create the response
	response := SearchUserResponse{Uuids: uuids}
	return response, nil
}
