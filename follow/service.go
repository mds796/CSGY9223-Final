package follow

import (
	"context"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
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
	FollowingGraph map[string][]*followpb.User
}

func CreateService(userService userpb.UserClient) *Service {
	stub := new(Service)
	stub.User = userService
	stub.FollowingGraph = make(map[string][]*followpb.User)
	return stub
}

func (stub *Service) Follow(ctx context.Context, request *followpb.FollowRequest) (*followpb.FollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(ctx, request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = stub.validateUserID(ctx, request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Avoid duplicated connections
	followed := stub.FollowingGraph[request.FollowerUser.ID]
	newConnection := true
	for _, f := range followed {
		if f.ID == request.FollowedUser.ID {
			newConnection = false
		}
	}

	// Add followed user from follow graph
	if newConnection {
		stub.FollowingGraph[request.FollowerUser.ID] = append(stub.FollowingGraph[request.FollowerUser.ID], &followpb.User{ID: request.FollowedUser.ID, Followed: true})
	}

	return &followpb.FollowResponse{}, nil
}

func (stub *Service) Unfollow(ctx context.Context, request *followpb.UnfollowRequest) (*followpb.UnfollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(ctx, request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = stub.validateUserID(ctx, request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Remove followed user from follow graph
	followed := stub.FollowingGraph[request.FollowerUser.ID]
	for i := 0; i < len(followed); i++ {
		if followed[i].ID == request.FollowedUser.ID {
			followed = append(followed[:i], followed[i+1:]...)
		}
	}
	stub.FollowingGraph[request.FollowerUser.ID] = followed
	return &followpb.UnfollowResponse{}, nil
}

func (stub *Service) View(ctx context.Context, request *followpb.ViewRequest) (*followpb.ViewResponse, error) {

	// Validate user ID
	ok := stub.validateUserID(ctx, request.User.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.User.ID}
	}

	// Return user's adjacency list
	users := stub.FollowingGraph[request.User.ID]
	return &followpb.ViewResponse{Users: users}, nil
}

func (stub *Service) Search(ctx context.Context, request *followpb.SearchRequest) (*followpb.SearchResponse, error) {

	// // Validate user ID
	ok := stub.validateUserID(ctx, request.User.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.User.ID}
	}

	userResponse, err := stub.User.Search(ctx, &userpb.SearchUserRequest{Query: request.Query})
	if err != nil {
		return nil, err
	}

	response := []*followpb.User{}

	// Caution: sub-optimal O(n^2) search
	for _, userID := range userResponse.UIDs {
		if userID != request.User.ID {
			followed := false

			for _, followedUser := range stub.FollowingGraph[request.User.ID] {
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

func (stub *Service) validateUserID(ctx context.Context, userID string) bool {
	_, err := stub.User.View(ctx, &userpb.ViewUserRequest{UID: userID})
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
