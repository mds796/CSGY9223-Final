package follow

import (
	"github.com/mds796/CSGY9223-Final/user"
)

type StubService struct {
	UserService user.Service

	// Storing connections in adjacency list allows to follow and unfollow in
	// O(n) time and retrieve list of followers in O(1).
	//
	// Using a hash set (map[string]bool), we can achieve follow and unfollow in
	// O(1) average time, however, retrieving the list of followers will require
	// to iterate through the hash set keys in O(n) time.
	FollowingGraph map[string][]string
}

func CreateStub(userService user.Service) Service {
	stub := new(StubService)
	stub.UserService = userService
	stub.FollowingGraph = make(map[string][]string)
	return stub
}

func (stub *StubService) Follow(request FollowRequest) (FollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(request.FollowerUserID)
	if !ok {
		return FollowResponse{}, &InvalidUserIDError{UserID: request.FollowerUserID}
	}

	ok = stub.validateUserID(request.FollowedUserID)
	if !ok {
		return FollowResponse{}, &InvalidUserIDError{UserID: request.FollowedUserID}
	}

	// Avoid duplicated connections
	followed := stub.FollowingGraph[request.FollowerUserID]
	newConnection := true
	for _, f := range followed {
		if f == request.FollowedUserID {
			newConnection = false
		}
	}

	// Add followed user from follow graph
	if newConnection {
		stub.FollowingGraph[request.FollowerUserID] = append(stub.FollowingGraph[request.FollowerUserID], request.FollowedUserID)
	}

	return FollowResponse{}, nil
}

func (stub *StubService) Unfollow(request UnfollowRequest) (UnfollowResponse, error) {
	// Validate user IDs
	ok := stub.validateUserID(request.FollowerUserID)
	if !ok {
		return UnfollowResponse{}, &InvalidUserIDError{UserID: request.FollowerUserID}
	}

	ok = stub.validateUserID(request.FollowedUserID)
	if !ok {
		return UnfollowResponse{}, &InvalidUserIDError{UserID: request.FollowedUserID}
	}

	// Remove followed user from follow graph
	followed := stub.FollowingGraph[request.FollowerUserID]
	for i := 0; i < len(followed); i++ {
		if followed[i] == request.FollowedUserID {
			followed = append(followed[:i], followed[i+1:]...)
		}
	}
	stub.FollowingGraph[request.FollowerUserID] = followed
	return UnfollowResponse{}, nil
}

func (stub *StubService) View(request ViewRequest) (ViewResponse, error) {
	// Validate user ID
	ok := stub.validateUserID(request.UserID)
	if !ok {
		return ViewResponse{}, &InvalidUserIDError{UserID: request.UserID}
	}

	// Return user's adjacency list
	userIDs := stub.FollowingGraph[request.UserID]
	return ViewResponse{UserIDs: userIDs}, nil
}

func (stub *StubService) validateUserID(userID string) bool {
	_, err := stub.UserService.View(user.ViewUserRequest{UserID: userID})
	if err != nil {
		return false
	}
	return true
}
