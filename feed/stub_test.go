package feed

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/user"
	"testing"
)

func TestStubService_View(t *testing.T) {
	service := createFeed()
	_, err := service.View(&ViewRequest{UserID: "does-not-exit"})

	if err != nil {
		t.Fatalf("Expected to receive an empty response, instead got '%v' as an error.\n", err)
	}
}

func TestStubService_View_WithUserEmptyFeed(t *testing.T) {
	service := createFeed()
	response, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})

	_, err := service.View(&ViewRequest{UserID: response.Uuid, Username: "fake123"})

	if err != nil {
		t.Fatalf("Expected to receive an empty response, instead got '%v' as an error.\n", err)
	}
}

func TestStubService_View_WithUserSelfPost(t *testing.T) {
	service := createFeed()
	userResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})
	_, _ = service.User.Create(user.CreateUserRequest{Username: "fake234"})

	message := "Hello, World!"
	_, _ = service.Post.Create(post.CreatePostRequest{Text: message, UserID: userResponse.Uuid})

	response, err := service.View(&ViewRequest{UserID: userResponse.Uuid, Username: "fake123"})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Posts) != 1 || response.Posts[0].From != "fake123" || response.Posts[0].Text != message {
		t.Fatalf("Received unexpected feed response: %v.\n", response)
	}
}

func TestStubService_View_WithUserPostNoFollow(t *testing.T) {
	service := createFeed()
	userResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake234"})

	_, _ = service.Post.Create(post.CreatePostRequest{Text: "Hello, World!", UserID: otherUserResponse.Uuid})

	response, err := service.View(&ViewRequest{UserID: userResponse.Uuid})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected an empty feed.\n", response)
	}
}

func TestStubService_View_WithUserFollowedPost(t *testing.T) {
	service := createFeed()
	userResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake234"})
	service.Follow.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})

	message := "Hello, World!"
	_, _ = service.Post.Create(post.CreatePostRequest{Text: message, UserID: otherUserResponse.Uuid})

	response, err := service.View(&ViewRequest{UserID: userResponse.Uuid})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Posts) != 1 || response.Posts[0].From != "fake234" || response.Posts[0].Text != message {
		t.Log("Received unexpected feed response:")
		for _, post := range response.Posts {
			t.Log(post)
		}
		t.Fail()
	}
}

func TestStubService_View_WithUsersFollowedNoPost(t *testing.T) {
	service := createFeed()
	userResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake234"})
	service.Follow.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})

	response, err := service.View(&ViewRequest{UserID: userResponse.Uuid})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected empty feed.\n", response)
	}
}

func TestStubService_View_ListPostsByReverseCreateOrder(t *testing.T) {
	service := createFeed()
	userResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := service.User.Create(user.CreateUserRequest{Username: "fake234"})
	service.Follow.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})

	_, _ = service.Post.Create(post.CreatePostRequest{Text: "post 1", UserID: userResponse.Uuid})
	_, _ = service.Post.Create(post.CreatePostRequest{Text: "post 2", UserID: otherUserResponse.Uuid})
	_, _ = service.Post.Create(post.CreatePostRequest{Text: "post 3", UserID: userResponse.Uuid})
	_, _ = service.Post.Create(post.CreatePostRequest{Text: "post 4", UserID: otherUserResponse.Uuid})

	response, err := service.View(&ViewRequest{UserID: userResponse.Uuid, Username: "fake123"})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if response.Posts[0].From != "fake234" || response.Posts[0].Text != "post 4" ||
		response.Posts[1].From != "fake123" || response.Posts[1].Text != "post 3" ||
		response.Posts[2].From != "fake234" || response.Posts[2].Text != "post 2" ||
		response.Posts[3].From != "fake123" || response.Posts[3].Text != "post 1" {
		t.Log("Received unexpected feed response:")
		for _, post := range response.Posts {
			t.Log(post)
		}
		t.Fail()
	}
}

func createFeed() *StubService {
	return &StubService{User: user.CreateStub(), Post: post.CreateStub(), Follow: follow.CreateStub()}
}
