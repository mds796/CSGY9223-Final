package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"strings"
)

const MIN_USERNAME = 6

type StubService struct {
	UuidCache     map[string]string // (UUID, username)
	UsernameCache map[string]string // (username, UUID)
}

func CreateStub() *StubService {
	stub := new(StubService)
	stub.UuidCache = make(map[string]string)
	stub.UsernameCache = make(map[string]string)
	return stub
}

func (s *StubService) Create(ctx context.Context, request *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// ensure this username meets the minimum requirements
	if len(request.Username) < MIN_USERNAME {
		return &userpb.CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// ensure this username doesn't already exist
	if _, ok := s.UsernameCache[request.Username]; ok {
		return &userpb.CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// generate the uuid
	newUuid := uuid.New().String()

	// add the user
	s.UuidCache[newUuid] = request.Username
	s.UsernameCache[request.Username] = newUuid

	// create the response
	response := &userpb.CreateUserResponse{UID: newUuid}
	return response, nil
}

func (s *StubService) View(ctx context.Context, request *userpb.ViewUserRequest) (*userpb.ViewUserResponse, error) {
	if id, ok := s.UsernameCache[request.Username]; ok {
		// username exists
		return &userpb.ViewUserResponse{UID: id, Username: request.Username}, nil
	} else if username, ok := s.UuidCache[request.UID]; ok {
		return &userpb.ViewUserResponse{UID: request.UID, Username: username}, nil
	} else {
		// username doesn't exist
		return &userpb.ViewUserResponse{}, &ViewUserError{request.Username}
	}
}

func (s *StubService) Search(ctx context.Context, request *userpb.SearchUserRequest) (*userpb.SearchUserResponse, error) {
	// find uuids that match given query
	var usernames []string
	var userIds []string
	for username, userId := range s.UsernameCache {
		if strings.Contains(username, request.Query) {
			usernames = append(usernames, username)
			userIds = append(userIds, userId)
		}
	}

	// create the response
	response := &userpb.SearchUserResponse{Usernames: usernames, UIDs: userIds}
	return response, nil
}
