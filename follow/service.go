package follow

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"google.golang.org/grpc"
)

type Service struct {
	User userpb.UserClient

	// Storing connections in adjacency list allows to follow and unfollow in
	// O(n) time and retrieve list of followers in O(1).
	//
	// Using a hash set (map[string]bool), we can achieve follow and unfollow in
	// O(1) average time, however, retrieving the list of followers will require
	// to iterate through the hash set keys in O(n) time.

	// FollowingGraph map[string][]*followpb.User
	FollowingGraph storage.Storage
}

func CreateService(storageConfig storage.StorageConfig, userService userpb.UserClient) *Service {
	service := new(Service)
	service.User = userService
	service.FollowingGraph = storage.CreateStorage(storageConfig, "follow/following_graph")
	return service
}

func (service *Service) Follow(ctx context.Context, request *followpb.FollowRequest) (*followpb.FollowResponse, error) {
	// Validate user IDs
	ok := service.validateUserID(ctx, request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = service.validateUserID(ctx, request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Avoid duplicated connections
	followedBytes, _ := service.FollowingGraph.Get(request.FollowerUser.ID)
	followed := &followpb.Users{}
	proto.Unmarshal(followedBytes, followed)

	newConnection := true
	for _, f := range followed.Users {
		if f.ID == request.FollowedUser.ID {
			newConnection = false
		}
	}

	// Add followed user from follow graph
	if newConnection {
		user := &followpb.User{ID: request.FollowedUser.ID, Followed: true}

		followed = &followpb.Users{Users: append(followed.Users, user)}
		followedBytes, _ = proto.Marshal(followed)
		service.FollowingGraph.Put(request.FollowerUser.ID, followedBytes)
	}

	return &followpb.FollowResponse{}, nil
}

func (service *Service) Unfollow(ctx context.Context, request *followpb.UnfollowRequest) (*followpb.UnfollowResponse, error) {
	// Validate user IDs
	ok := service.validateUserID(ctx, request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = service.validateUserID(ctx, request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Remove followed user from follow graph
	followedBytes, _ := service.FollowingGraph.Get(request.FollowerUser.ID)
	followed := &followpb.Users{}
	proto.Unmarshal(followedBytes, followed)

	for i := 0; i < len(followed.Users); i++ {
		if followed.Users[i].ID == request.FollowedUser.ID {
			followed.Users = append(followed.Users[:i], followed.Users[i+1:]...)
		}
	}

	followedBytes, _ = proto.Marshal(followed)
	service.FollowingGraph.Put(request.FollowerUser.ID, followedBytes)
	return &followpb.UnfollowResponse{}, nil
}

func (service *Service) View(ctx context.Context, request *followpb.ViewRequest) (*followpb.ViewResponse, error) {

	// Validate user ID
	ok := service.validateUserID(ctx, request.User.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.User.ID}
	}

	// Return user's adjacency list
	followedBytes, _ := service.FollowingGraph.Get(request.User.ID)
	followed := &followpb.Users{}
	proto.Unmarshal(followedBytes, followed)
	return &followpb.ViewResponse{Users: followed.Users}, nil
}

func (service *Service) Search(ctx context.Context, request *followpb.SearchRequest) (*followpb.SearchResponse, error) {

	// // Validate user ID
	ok := service.validateUserID(ctx, request.User.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.User.ID}
	}

	userResponse, err := service.User.Search(ctx, &userpb.SearchUserRequest{Query: request.Query})
	if err != nil {
		return nil, err
	}

	response := []*followpb.User{}

	// Caution: sub-optimal O(n^2) search
	for _, userID := range userResponse.UIDs {
		if userID != request.User.ID {
			followed := false

			usersBytes, _ := service.FollowingGraph.Get(request.User.ID)
			users := &followpb.Users{}
			proto.Unmarshal(usersBytes, users)

			for _, followedUser := range users.Users {
				if followedUser.ID == userID {
					followed = true
					break
				}
			}

			response = append(response, &followpb.User{ID: userID, Followed: followed})
		}
	}

	return &followpb.SearchResponse{Users: response}, nil
}

func (service *Service) validateUserID(ctx context.Context, userID string) bool {
	_, err := service.User.View(ctx, &userpb.ViewUserRequest{UID: userID})
	if err != nil {
		return false
	}
	return true
}

type StubClient struct {
	service followpb.FollowServer
}

func (s StubClient) Follow(ctx context.Context, in *followpb.FollowRequest, opts ...grpc.CallOption) (*followpb.FollowResponse, error) {
	return s.service.Follow(ctx, in)
}

func (s StubClient) Unfollow(ctx context.Context, in *followpb.UnfollowRequest, opts ...grpc.CallOption) (*followpb.UnfollowResponse, error) {
	return s.service.Unfollow(ctx, in)
}

func (s StubClient) View(ctx context.Context, in *followpb.ViewRequest, opts ...grpc.CallOption) (*followpb.ViewResponse, error) {
	return s.service.View(ctx, in)
}

func (s StubClient) Search(ctx context.Context, in *followpb.SearchRequest, opts ...grpc.CallOption) (*followpb.SearchResponse, error) {
	return s.service.Search(ctx, in)
}
