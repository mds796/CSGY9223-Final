package follow

import (
	"context"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/user"
)

type StubService struct {
	User user.Service

	// Storing connections in adjacency list allows to follow and unfollow in
	// O(n) time and retrieve list of followers in O(1).
	//
	// Using a hash set (map[string]bool), we can achieve follow and unfollow in
	// O(1) average time, however, retrieving the list of followers will require
	// to iterate through the hash set keys in O(n) time.
	FollowingGraph map[string][]string
}

func CreateStub(userService user.Service) *StubService {
	stub := new(StubService)
	stub.User = userService
	stub.FollowingGraph = make(map[string][]string)
	return stub
}

func (stub *StubService) Follow(ctx context.Context, request *followpb.FollowRequest) (*followpb.FollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = stub.validateUserID(request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Avoid duplicated connections
	followed := stub.FollowingGraph[request.FollowerUser.ID]
	newConnection := true
	for _, f := range followed {
		if f == request.FollowedUser.ID {
			newConnection = false
		}
	}

	// Add followed user from follow graph
	if newConnection {
		stub.FollowingGraph[request.FollowerUser.ID] = append(stub.FollowingGraph[request.FollowerUser.ID], request.FollowedUser.ID)
	}

	return &followpb.FollowResponse{}, nil
}

func (stub *StubService) Unfollow(ctx context.Context, request *followpb.UnfollowRequest) (*followpb.UnfollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(request.FollowerUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowerUser.ID}
	}

	ok = stub.validateUserID(request.FollowedUser.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.FollowedUser.ID}
	}

	// Remove followed user from follow graph
	followed := stub.FollowingGraph[request.FollowerUser.ID]
	for i := 0; i < len(followed); i++ {
		if followed[i] == request.FollowedUser.ID {
			followed = append(followed[:i], followed[i+1:]...)
		}
	}
	stub.FollowingGraph[request.FollowerUser.ID] = followed
	return &followpb.UnfollowResponse{}, nil
}

func (stub *StubService) View(ctx context.Context, request *followpb.ViewRequest) (*followpb.ViewResponse, error) {

	// Validate user ID
	ok := stub.validateUserID(request.User.ID)
	if !ok {
		return nil, &InvalidUserIDError{UserID: request.User.ID}
	}

	// Return user's adjacency list
	userIDs := stub.FollowingGraph[request.User.ID]
	users := []*followpb.User{}
	for _, userID := range userIDs {
		users = append(users, &followpb.User{ID: userID})
	}

	return &followpb.ViewResponse{Users: users}, nil
}

func (stub *StubService) validateUserID(userID string) bool {
	_, err := stub.User.View(user.ViewUserRequest{UserID: userID})
	if err != nil {
		return false
	}
	return true
}
