package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"google.golang.org/grpc"
	"strings"
)

const MIN_USERNAME = 6

type Service struct {
	UIDCache      storage.Storage // (UID, username)
	UsernameCache storage.Storage // (username, UID)
}

func CreateStub(storageType storage.StorageType) *Service {
	service := new(Service)
	service.UIDCache = storage.CreateStorage(storageType, "user/uid_cache")
	service.UsernameCache = storage.CreateStorage(storageType, "user/username_cache")
	return service
}

func (s *Service) Create(ctx context.Context, request *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// ensure this username meets the minimum requirements
	if len(request.Username) < MIN_USERNAME {
		return &userpb.CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// ensure this username doesn't already exist
	if _, err := s.UsernameCache.Get(request.Username); err == nil {
		return &userpb.CreateUserResponse{}, &CreateUserError{request.Username}
	}

	// generate the UID
	newUID := uuid.New().String()

	// add the user
	s.UIDCache.Put(newUID, []byte(request.Username))
	s.UsernameCache.Put(request.Username, []byte(newUID))

	// create the response
	response := &userpb.CreateUserResponse{UID: newUID}
	return response, nil
}

func (s *Service) View(ctx context.Context, request *userpb.ViewUserRequest) (*userpb.ViewUserResponse, error) {
	if id, err := s.UsernameCache.Get(request.Username); err == nil {
		// username exists
		return &userpb.ViewUserResponse{UID: string(id), Username: request.Username}, nil
	} else if username, err := s.UIDCache.Get(request.UID); err == nil {
		return &userpb.ViewUserResponse{UID: string(request.UID), Username: string(username)}, nil
	} else {
		// username doesn't exist
		return &userpb.ViewUserResponse{}, &ViewUserError{request.Username}
	}
}

func (s *Service) Search(ctx context.Context, request *userpb.SearchUserRequest) (*userpb.SearchUserResponse, error) {
	// find UIDs that match given query
	var usernames []string
	var userIds []string
	for username, userId := range s.UsernameCache.Iterate() {
		if strings.Contains(username, request.Query) {
			usernames = append(usernames, username)
			userIds = append(userIds, string(userId))
		}
	}

	// create the response
	response := &userpb.SearchUserResponse{Usernames: usernames, UIDs: userIds}
	return response, nil
}

type StubClient struct {
	service userpb.UserServer
}

func (s StubClient) Create(ctx context.Context, in *userpb.CreateUserRequest, opts ...grpc.CallOption) (*userpb.CreateUserResponse, error) {
	return s.service.Create(ctx, in)
}

func (s StubClient) View(ctx context.Context, in *userpb.ViewUserRequest, opts ...grpc.CallOption) (*userpb.ViewUserResponse, error) {
	return s.service.View(ctx, in)
}

func (s StubClient) Search(ctx context.Context, in *userpb.SearchUserRequest, opts ...grpc.CallOption) (*userpb.SearchUserResponse, error) {
	return s.service.Search(ctx, in)
}
