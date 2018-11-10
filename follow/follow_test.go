package follow

import "testing"

import "github.com/google/uuid"

func TestFollowDoesNotReturnError(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := uuid.New().String()
	_, err := service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	if err != nil {
		t.Fail()
	}
}

func TestFollowDoesNotDuplicateConnections(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := uuid.New().String()
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(viewResponse.UserIDs) != 1 {
		t.Fail()
	}
}

func TestUnfollowAfterFollowing(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := uuid.New().String()
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	_, err := service.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	if err != nil {
		t.Fail()
	}
}

func TestViewReturnsEmptyListForUnknownUserID(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(viewResponse.UserIDs) > 0 {
		t.Fail()
	}
}

func TestViewReturnsFollowedUsers(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}

	for i := 0; i < len(followed); i++ {
		service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[i]})
	}

	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(followed) != len(viewResponse.UserIDs) {
		t.Fail()
	}

	for i := 0; i < len(followed); i++ {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fail()
		}
	}
}

func TestViewReturnsCorrectFollowedUsers(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := []string{uuid.New().String(), uuid.New().String()}
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[0]})
	service.Follow(FollowRequest{FollowerUserID: uuid.New().String(), FollowedUserID: uuid.New().String()}) // should not appear in response
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(followed) != len(viewResponse.UserIDs) {
		t.Fail()
	}

	for i := 0; i < len(followed); i++ {
		if viewResponse.UserIDs[i] != followed[i] {
			t.Fail()
		}
	}
}

func TestViewDoesNotReturnUnfollowedUsers(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := uuid.New().String()
	service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed})
	service.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed})

	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(viewResponse.UserIDs) > 0 {
		t.Fail()
	}
}

func TestUnfollowRemovesCorrectConnection(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	followed := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}

	for i := 0; i < len(followed); i++ {
		service.Follow(FollowRequest{FollowerUserID: follower, FollowedUserID: followed[i]})
	}

	service.Unfollow(UnfollowRequest{FollowerUserID: follower, FollowedUserID: followed[1]})

	expectedFollowed := []string{followed[0], followed[2]}
	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(expectedFollowed) != len(viewResponse.UserIDs) {
		t.Fail()
	}

	for i := 0; i < len(expectedFollowed); i++ {
		if viewResponse.UserIDs[i] != expectedFollowed[i] {
			t.Fail()
		}
	}
}
