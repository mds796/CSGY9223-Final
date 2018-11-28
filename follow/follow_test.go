package follow

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func createFollowService() (*StubClient, user.Service) {
	userService := user.CreateStub()
	return &StubClient{service: CreateStub(userService)}, userService
}

func createUsers(userService user.Service, usernames ...string) []string {
	uuids := []string{}
	for _, username := range usernames {
		uuids = append(uuids, createUser(userService, username))
	}
	return uuids
}

func createUser(userService user.Service, username string) string {
	userResponse, _ := userService.Create(user.CreateUserRequest{Username: username})
	return userResponse.Uuid
}

func followUser(followService *StubClient, followerUserID string, followedUserID string) (*followpb.FollowResponse, error) {
	return followService.Follow(
		context.Background(),
		&followpb.FollowRequest{
			FollowerUser: &followpb.User{ID: followerUserID},
			FollowedUser: &followpb.User{ID: followedUserID},
		})
}

func unfollowUser(followService *StubClient, followerUserID string, followedUserID string) (*followpb.UnfollowResponse, error) {
	return followService.Unfollow(
		context.Background(),
		&followpb.UnfollowRequest{
			FollowerUser: &followpb.User{ID: followerUserID},
			FollowedUser: &followpb.User{ID: followedUserID},
		})
}

func viewFollowedUsers(followService *StubClient, followerUserID string) (*followpb.ViewResponse, error) {
	return followService.View(
		context.Background(),
		&followpb.ViewRequest{
			User: &followpb.User{ID: followerUserID},
		})
}

func TestFollowDoesNotReturnError(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUser(userService, "fake234")
	_, err := followUser(followService, follower, followed)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFollowDoesNotDuplicateConnections(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUser(userService, "fake234")
	followUser(followService, follower, followed)
	followUser(followService, follower, followed)

	viewResponse, err := viewFollowedUsers(followService, follower)
	if err != nil || len(viewResponse.Users) != 1 {
		t.Fatal()
	}
}

func TestUnfollowAfterFollowing(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUser(userService, "fake234")
	followUser(followService, follower, followed)
	_, err := unfollowUser(followService, follower, followed)
	if err != nil {
		t.Fatal()
	}
}

func TestViewReturnsFollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUsers(userService, "fake234", "fake345", "fake456")
	for _, f := range followed {
		followUser(followService, follower, f)
	}

	viewResponse, err := viewFollowedUsers(followService, follower)
	if err != nil || len(followed) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i := range followed {
		if viewResponse.Users[i].ID != followed[i] {
			t.Fatal()
		}
	}
}

func TestViewReturnsCorrectFollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUsers(userService, "fake234", "fake345")
	notFollowed := createUser(userService, "fake456") // should not appear in response

	followUser(followService, follower, followed[0])
	followUser(followService, follower, followed[1])

	viewResponse, err := viewFollowedUsers(followService, follower)
	if err != nil || len(followed) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i := range followed {
		if viewResponse.Users[i].ID != followed[i] {
			t.Fatal()
		}
	}

	for _, user := range viewResponse.Users {
		if user.ID == notFollowed {
			t.Fatal()
		}
	}
}

func TestViewDoesNotReturnUnfollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUser(userService, "fake234")
	followUser(followService, follower, followed)
	unfollowUser(followService, follower, followed)

	viewResponse, err := viewFollowedUsers(followService, follower)
	if err != nil || len(viewResponse.Users) > 0 {
		t.Fatal()
	}
}

func TestUnfollowRemovesCorrectConnection(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := createUsers(userService, "fake234", "fake345", "fake456")

	for _, f := range followed {
		followUser(followService, follower, f)
	}

	unfollowUser(followService, follower, followed[1])

	expectedFollowed := []string{followed[0], followed[2]}
	viewResponse, _ := viewFollowedUsers(followService, follower)

	if len(expectedFollowed) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i := range expectedFollowed {
		if viewResponse.Users[i].ID != expectedFollowed[i] {
			t.Fatal()
		}
	}
}

func TestFollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService, userService := createFollowService()
	follower := uuid.New().String()
	followed := createUser(userService, "fake234")
	_, err := followUser(followService, follower, followed)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestFollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := uuid.New().String()
	_, err := followUser(followService, follower, followed)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService, userService := createFollowService()
	follower := uuid.New().String()
	followed := createUser(userService, "fake234")
	_, err := unfollowUser(followService, follower, followed)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService, userService := createFollowService()
	follower := createUser(userService, "fake123")
	followed := uuid.New().String()
	_, err := unfollowUser(followService, follower, followed)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestViewReturnsErrorForUnknownUserID(t *testing.T) {
	followService, _ := createFollowService()
	follower := uuid.New().String()
	_, err := viewFollowedUsers(followService, follower)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}
