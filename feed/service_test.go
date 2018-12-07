package feed

import (
	"context"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"testing"
)

func createFeed() (*StubClient, userpb.UserClient, postpb.PostClient, followpb.FollowClient) {
	userService := user.NewStubClient()
	postService := post.NewStubClient()
	followService := follow.NewStubClient(follow.CreateStub(userService))

	return &StubClient{service: &StubService{User: userService, Post: postService, Follow: followService}}, userService, postService, followService
}

func TestStubService_View_WithNonExistingUser(t *testing.T) {
	client, _, _, _ := createFeed()

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: "does-not-exit", Name: "fake007"}})

	if err == nil {
		t.Fatalf("Expected to receive an error, instead got '%v' as response.\n", response)
	}
}

func TestStubService_View_WithUserEmptyFeed(t *testing.T) {
	client, u, _, _ := createFeed()
	response, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})

	_, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: response.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive an empty response, instead got '%v' as an error.\n", err)
	}
}

func TestStubService_View_WithUserSelfPost(t *testing.T) {
	client, u, p, _ := createFeed()
	userResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})
	_, _ = u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake234"})

	message := "Hello, World!"
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.UID}, Post: &postpb.Post{Text: message}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 1 || response.Feed.Posts[0].User.Name != "fake123" || response.Feed.Posts[0].Text != message {
		t.Fatalf("Received unexpected feed response: %v.\n", response)
	}
}

func TestStubService_View_WithUserPostNoFollow(t *testing.T) {
	client, u, p, _ := createFeed()
	userResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake234"})

	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.UID}, Post: &postpb.Post{Text: "Hello, World!"}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected an empty feed.\n", response)
	}
}

func TestStubService_View_WithUserFollowedPost(t *testing.T) {
	client, u, p, f := createFeed()

	userResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(context.Background(), &followpb.FollowRequest{FollowerUser: &followpb.User{ID: userResponse.UID}, FollowedUser: &followpb.User{ID: otherUserResponse.UID}})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	message := "Hello, World!"
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.UID}, Post: &postpb.Post{Text: message}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 1 || response.Feed.Posts[0].User.Name != "fake234" || response.Feed.Posts[0].Text != message {
		t.Log("Received unexpected feed response:")

		for _, p := range response.Feed.Posts {
			t.Log(p)
		}

		t.Fail()
	}
}

func TestStubService_View_WithUsersFollowedNoPost(t *testing.T) {
	client, u, _, f := createFeed()

	userResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(context.Background(), &followpb.FollowRequest{FollowerUser: &followpb.User{ID: userResponse.UID}, FollowedUser: &followpb.User{ID: otherUserResponse.UID}})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected empty feed.\n", response)
	}
}

func TestStubService_View_ListPostsByReverseCreateOrder(t *testing.T) {
	client, u, p, f := createFeed()

	userResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(context.Background(), &userpb.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(context.Background(), &followpb.FollowRequest{FollowerUser: &followpb.User{ID: userResponse.UID}, FollowedUser: &followpb.User{ID: otherUserResponse.UID}})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.UID}, Post: &postpb.Post{Text: "post 1"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.UID}, Post: &postpb.Post{Text: "post 2"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.UID}, Post: &postpb.Post{Text: "post 3"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.UID}, Post: &postpb.Post{Text: "post 4"}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.UID, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	expectedUsers := []string{"fake234", "fake123", "fake234", "fake123"}
	expectedTexts := []string{"post 4", "post 3", "post 2", "post 1"}

	for i, post := range response.Feed.Posts {
		if post.User.Name != expectedUsers[i] || post.Text != expectedTexts[i] {
			t.Errorf("Received unexpected post response, <'%v', '%v'> instead of <'%v', '%v'>\n", post.User.Name, post.Text, expectedUsers[i], expectedTexts[i])
		}
	}
}
