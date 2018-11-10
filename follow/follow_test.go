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

func TestViewReturnsEmptyListForUnknownUserID(t *testing.T) {
	service := CreateStub()
	follower := uuid.New().String()
	viewResponse, _ := service.View(ViewRequest{UserID: follower})

	if len(viewResponse.UserIDs) > 0 {
		t.Fail()
	}
}
