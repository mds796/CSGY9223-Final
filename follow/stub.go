package follow

type StubService struct {
	// Storing connections in adjacency list allows to follow and unfollow in
	// O(n) time and retrieve list of followers in O(1).
	//
	// Using a hash set (map[string]bool), we can achieve follow and unfollow in
	// O(1) average time, however, retrieving the list of followers will require
	// to iterate through the hash set keys in O(n) time.
	FollowingGraph map[string][]string
}

func CreateStub() Service {
	stub := new(StubService)
	stub.FollowingGraph = make(map[string][]string)
	return stub
}

func (stub *StubService) Follow(request FollowRequest) (FollowResponse, error) {
	followed := stub.FollowingGraph[request.FollowerUserID]

	// Avoid duplicated connections
	newConnection := true
	for i := 0; i < len(followed); i++ {
		if followed[i] == request.FollowedUserID {
			newConnection = false
		}
	}

	if newConnection {
		stub.FollowingGraph[request.FollowerUserID] = append(stub.FollowingGraph[request.FollowerUserID], request.FollowedUserID)
	}

	return FollowResponse{}, nil
}

func (stub *StubService) Unfollow(request UnfollowRequest) (UnfollowResponse, error) {
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
	userIDs := stub.FollowingGraph[request.UserID]
	return ViewResponse{UserIDs: userIDs}, nil
}
