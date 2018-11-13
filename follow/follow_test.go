package follow

import (
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func TestFollowDoesNotReturnError(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUser(followService, "fake234")
	_, err := followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	if err != nil {
		t.Fatal()
	}
}

func TestFollowDoesNotDuplicateConnections(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUser(followService, "fake234")
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(viewResponse.UserIDs) != 1 {
		t.Fatal()
	}
}

func TestUnfollowAfterFollowing(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUser(followService, "fake234")
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, err := followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	if err != nil {
		t.Fatal()
	}
}

func TestViewReturnsFollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUsers(followService, "fake234", "fake345", "fake456")
	for _, f := range followed {
		followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: f})
	}

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(followed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := range followed {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fatal()
		}
	}
}

func TestViewReturnsCorrectFollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUsers(followService, "fake234", "fake345")
	notFollowed := createUser(followService, "fake456") // should not appear in response

	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[0]})
	followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	viewResponse, err := followService.View(ViewRequest{UserID: follower})
	if err != nil || len(followed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := range followed {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fatal()
		}
	}

	for _, uuid := range viewResponse.UserIDs {
		if uuid == notFollowed {
			t.Fatal()
		}
	}
}

func TestViewDoesNotReturnUnfollowedUsers(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := createUser(followService, "fake234")
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
	followed := createUsers(followService, "fake234", "fake345", "fake456")

	for _, f := range followed {
		followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: f})
	}

	followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	expectedFollowed := []string{followed[0], followed[2]}
	viewResponse, _ := followService.View(ViewRequest{UserID: follower})

	if len(expectedFollowed) != len(viewResponse.UserIDs) {
		t.Fatal()
	}

	for i := range expectedFollowed {
		if viewResponse.UserIDs[i] != expectedFollowed[i] {
			t.Fatal()
		}
	}
}

func TestFollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	followed := createUser(followService, "fake234")
	_, err := followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestFollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := uuid.New().String()
	_, err := followService.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService := createFeed()
	follower := uuid.New().String()
	followed := createUser(followService, "fake234")
	_, err := followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService := createFeed()
	follower := createUser(followService, "fake123")
	followed := uuid.New().String()
	_, err := followService.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, ok := err.(*InvalidUserIDError)
	if !ok {
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

func createFeed() *StubService {
	return &StubService{UserService: user.CreateStub(), FollowingGraph: make(map[string][]string)}
}

func createUsers(followService *StubService, usernames ...string) []string {
	uuids := []string{}
	for _, username := range usernames {
		uuids = append(uuids, createUser(followService, username))
	}
	return uuids
}

func createUser(followService *StubService, username string) string {
	userResponse, _ := followService.UserService.Create(user.CreateUserRequest{Username: username})
	return userResponse.Uuid
}
