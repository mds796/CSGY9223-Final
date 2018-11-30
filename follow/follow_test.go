package follow

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"testing"
)

func createFollowService() (*StubClient, *user.StubClient) {
	userService := user.NewStubClient(user.CreateStub())
	return &StubClient{service: CreateStub(userService)}, userService
}

func createUsers(userService *user.StubClient, usernames ...string) []string {
	uuids := []string{}
	for _, username := range usernames {
		uuids = append(uuids, createUser(userService, username))
	}
	return uuids
}

func createUser(userService *user.StubClient, username string) string {
	createUserRequest := &userpb.CreateUserRequest{Username: username}
	createUserResponse, _ := userService.Create(context.Background(), createUserRequest)
	return createUserResponse.UID
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

func searchFollowedUsers(followService *StubClient, followerUserID string, query string) (*followpb.SearchResponse, error) {
	return followService.Search(
		context.Background(),
		&followpb.SearchRequest{
			User:  &followpb.User{ID: followerUserID},
			Query: query,
		})
}

func TestFollowDoesNotReturnError(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := createUser(userService, "fake234")
	_, err := followUser(followService, followerID, followedID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFollowDoesNotDuplicateConnections(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := createUser(userService, "fake234")
	followUser(followService, followerID, followedID)
	followUser(followService, followerID, followedID)

	viewResponse, err := viewFollowedUsers(followService, followerID)
	if err != nil || len(viewResponse.Users) != 1 {
		t.Fatal()
	}
}

func TestFollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService, userService := createFollowService()
	followerID := uuid.New().String()
	followedID := createUser(userService, "fake234")
	_, err := followUser(followService, followerID, followedID)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestFollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := uuid.New().String()
	_, err := followUser(followService, followerID, followedID)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowAfterFollowing(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := createUser(userService, "fake234")
	followUser(followService, followerID, followedID)
	_, err := unfollowUser(followService, followerID, followedID)
	if err != nil {
		t.Fatal()
	}
}

func TestUnfollowRemovesCorrectConnection(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedIDs := createUsers(userService, "fake234", "fake345", "fake456")

	for _, f := range followedIDs {
		followUser(followService, followerID, f)
	}

	unfollowUser(followService, followerID, followedIDs[1])

	expectedFollowed := []string{followedIDs[0], followedIDs[2]}
	viewResponse, _ := viewFollowedUsers(followService, followerID)

	if len(expectedFollowed) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i, user := range viewResponse.Users {
		if user.ID != expectedFollowed[i] {
			t.Fatal()
		}
	}
}

func TestUnfollowReturnsErrorForUnknownFollowerUserID(t *testing.T) {
	followService, userService := createFollowService()
	followerID := uuid.New().String()
	followedID := createUser(userService, "fake234")
	_, err := unfollowUser(followService, followerID, followedID)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestUnfollowReturnsErrorForUnknownFollowedUserID(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := uuid.New().String()
	_, err := unfollowUser(followService, followerID, followedID)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestViewReturnsFollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedIDs := createUsers(userService, "fake234", "fake345", "fake456")
	for _, f := range followedIDs {
		followUser(followService, followerID, f)
	}

	viewResponse, err := viewFollowedUsers(followService, followerID)
	if err != nil || len(followedIDs) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i, user := range viewResponse.Users {
		if user.ID != followedIDs[i] {
			t.Fatalf("Received unexpected user ID - received '%v' expected '%v'", user.ID, followedIDs[i])
		}

		if !user.Followed {
			t.Fatalf("Received user %v is marked as not followedID", user.ID)
		}
	}
}

func TestViewReturnsCorrectFollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedIDs := createUsers(userService, "fake234", "fake345")
	notFollowedID := createUser(userService, "fake456") // should not appear in response

	followUser(followService, followerID, followedIDs[0])
	followUser(followService, followerID, followedIDs[1])

	viewResponse, err := viewFollowedUsers(followService, followerID)
	if err != nil || len(followedIDs) != len(viewResponse.Users) {
		t.Fatal()
	}

	for i := range followedIDs {
		if viewResponse.Users[i].ID != followedIDs[i] {
			t.Fatal()
		}
	}

	for _, user := range viewResponse.Users {
		if user.ID == notFollowedID {
			t.Fatal()
		}
	}
}

func TestViewDoesNotReturnUnfollowedUsers(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedID := createUser(userService, "fake234")
	followUser(followService, followerID, followedID)
	unfollowUser(followService, followerID, followedID)

	viewResponse, err := viewFollowedUsers(followService, followerID)
	if err != nil || len(viewResponse.Users) > 0 {
		t.Fatal()
	}
}

func TestViewReturnsErrorForUnknownUserID(t *testing.T) {
	followService, _ := createFollowService()
	followerID := uuid.New().String()
	_, err := viewFollowedUsers(followService, followerID)
	_, ok := err.(*InvalidUserIDError)
	if !ok {
		t.Fatal()
	}
}

func TestSearchReturnsAllUsersWithFollowedFlag(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedIDs := createUsers(userService, "fake234", "fake789")
	notFollowedIDs := createUsers(userService, "fake456", "fake345")

	status := make(map[string]bool)
	for _, user := range followedIDs {
		followUser(followService, followerID, user)
		status[user] = true
	}
	for _, user := range notFollowedIDs {
		status[user] = false
	}

	searchResponse, err := searchFollowedUsers(followService, followerID, "")
	if err != nil || len(status) != len(searchResponse.Users) {
		t.Fatalf("Search returned an unexpected number of users: %v", len(searchResponse.Users))
	}

	for _, user := range searchResponse.Users {
		if user.Followed != status[user.ID] {
			t.Fatalf("Search returned wrong following status: expected '%v' received '%v'", status[user.ID], user.Followed)
		}
	}
}

func TestSearchReturnsFilteredUsersUsingQuery(t *testing.T) {
	followService, userService := createFollowService()
	followerID := createUser(userService, "fake123")
	followedIDs := createUsers(userService, "thisUserIsNice", "notThisOne")
	notFollowedID := createUsers(userService, "maybeThis", "thisUserIsWeird")

	for _, user := range followedIDs {
		followUser(followService, followerID, user)
	}

	searchResponse, err := searchFollowedUsers(followService, followerID, "User")
	expected := make(map[string]bool)
	expected[followedIDs[0]] = true
	expected[notFollowedID[1]] = false

	if err != nil || len(expected) != len(searchResponse.Users) {
		t.Fatalf("Search returned an unexpected number of users: %v", len(searchResponse.Users))
	}

	for _, user := range searchResponse.Users {
		if user.Followed != expected[user.ID] {
			t.Fatalf("Search returned wrong following status for user %v: expected '%v' received '%v'", user.ID, expected[user.ID], user.Followed)
		}
	}
}
