package follow

import (
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func TestFollowDoesNotReturnError(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	followed := uuid.New().String()
	_, err := followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	if err != nil {
		t.Fatal()
	}
}

func TestFollowDoesNotDuplicateConnections(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := uuid.New().String()
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(viewResponse.UserIDs) != 1 {
		t.Fatal()
	}
}

func TestUnfollowAfterFollowing(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	followed := uuid.New().String()
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, err := followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	if err != nil {
		t.Fatal()
	}
}

func TestViewReturnsEmptyListForUnknownUserID(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	viewResponse, _ := followService.View(ViewRequest{UserID: follower})
	if len(viewResponse.UserIDs) > 0 {
		t.Fatal()
	}
}

func TestViewReturnsErrorForUnknownUserID(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	_, err := followService.View(ViewRequest{UserID: follower})
	_, ok := err.(*InvalidUserIDError)

	if !ok {
		t.Fatal()
	}
}

func TestViewReturnsFollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}

	for i := 0; i < len(followed); i++ {
		followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[i]})
	}

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(followed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := 0; i < len(followed); i++ {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fatal()
		}
	}
}

func TestViewReturnsCorrectFollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := []string{uuid.New().String(), uuid.New().String()}
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[0]})
	followService.Follow(FollowRequest{FollowerUserID: uuid.New().String(), FollowedUserID: uuid.New().String()}) // should not appear in response
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(followed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := 0; i < len(followed); i++ {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fatal()
		}
	}
}

func TestViewDoesNotReturnUnfollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := uuid.New().String()
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(viewResponse.UserIDs) > 0 {
		t.Fatal()
	}
}

func TestUnfollowRemovesCorrectConnection(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}

	for i := 0; i < len(followed); i++ {
		followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[i]})
	}

	followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	expectedFollowed := []string{followed[0], followed[2]}
	viewResponse, _ := followService.View(ViewRequest{UserID: follower})

	if len(expectedFollowed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := 0; i < len(expectedFollowed); i++ {
		if viewResponse.UserIDs[i] != expectedFollowed[i] {
			t.Fatal()
		}
	}
}

func createFeed() *StubService {
	return &StubService{UserService: user.CreateStub(), FollowingGraph: make(map[string][]string)}
}

func createUser(followService *StubService, username string) string {
	userResponse, _ := followService.UserService.Create(user.CreateUserRequest{Username: username})
	return userResponse.Uuid
}
