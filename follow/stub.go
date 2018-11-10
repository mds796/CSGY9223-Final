package follow

type StubService struct {
	FollowingGraph map[string][]string
}

func CreateStub() Service {
	stub := new(StubService)
	stub.FollowingGraph = make(map[string][]string)
	return stub
}

func (stub *StubService) Follow(request FollowRequest) (FollowResponse, error) {
	stub.FollowingGraph[request.FollowerUserID] = append(stub.FollowingGraph[request.FollowerUserID], request.FollowedUserID)
	return FollowResponse{}, nil
}

func (stub *StubService) View(request ViewRequest) (ViewResponse, error) {
	userIDs := stub.FollowingGraph[request.UserID]
	return ViewResponse{UserIDs: userIDs}, nil
}
