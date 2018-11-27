package feed

import (
	"context"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/user"
	"strconv"
	"testing"
)

func createFeed() (*StubClient, user.Service, postpb.PostClient, follow.Service) {
	userService := user.CreateStub()
	postService := post.NewStubClient(post.CreateStub())
	followService := follow.CreateStub(userService)

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
	response, _ := u.Create(user.CreateUserRequest{Username: "fake123"})

	_, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: response.Uuid, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive an empty response, instead got '%v' as an error.\n", err)
	}
}

func TestStubService_View_WithUserSelfPost(t *testing.T) {
	client, u, p, _ := createFeed()
	userResponse, _ := u.Create(user.CreateUserRequest{Username: "fake123"})
	_, _ = u.Create(user.CreateUserRequest{Username: "fake234"})

	message := "Hello, World!"
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.Uuid}, Post: &postpb.Post{Text: message}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.Uuid, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 1 || response.Feed.Posts[0].User.Name != "fake123" || response.Feed.Posts[0].Text != message {
		t.Fatalf("Received unexpected feed response: %v.\n", response)
	}
}

func TestStubService_View_WithUserPostNoFollow(t *testing.T) {
	client, u, p, _ := createFeed()
	userResponse, _ := u.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(user.CreateUserRequest{Username: "fake234"})

	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.Uuid}, Post: &postpb.Post{Text: "Hello, World!"}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.Uuid, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected an empty feed.\n", response)
	}
}

func TestStubService_View_WithUserFollowedPost(t *testing.T) {
	client, u, p, f := createFeed()

	userResponse, _ := u.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(user.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	message := "Hello, World!"
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.Uuid}, Post: &postpb.Post{Text: message}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.Uuid, Name: "fake123"}})

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

	userResponse, _ := u.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(user.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.Uuid, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	if len(response.Feed.Posts) != 0 {
		t.Fatalf("Received unexpected feed response: %v. Expected empty feed.\n", response)
	}
}

func TestStubService_View_ListPostsByReverseCreateOrder(t *testing.T) {
	client, u, p, f := createFeed()

	userResponse, _ := u.Create(user.CreateUserRequest{Username: "fake123"})
	otherUserResponse, _ := u.Create(user.CreateUserRequest{Username: "fake234"})
	_, err := f.Follow(follow.FollowRequest{FollowerUserID: userResponse.Uuid, FollowedUserID: otherUserResponse.Uuid})
	if err != nil {
		t.Fatalf("Unable to follow other user.")
	}

	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.Uuid}, Post: &postpb.Post{Text: "post 1"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.Uuid}, Post: &postpb.Post{Text: "post 2"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: userResponse.Uuid}, Post: &postpb.Post{Text: "post 3"}})
	_, _ = p.Create(context.Background(), &postpb.CreateRequest{User: &postpb.User{ID: otherUserResponse.Uuid}, Post: &postpb.Post{Text: "post 4"}})

	response, err := client.View(context.Background(), &feedpb.ViewRequest{User: &feedpb.User{ID: userResponse.Uuid, Name: "fake123"}})

	if err != nil {
		t.Fatalf("Expected to receive a response, instead got '%v' as an error.\n", err)
	}

	order := [4]int{4, 2, 3, 1}
	for i, actual := range response.Feed.Posts {
		expectedText := "post " + strconv.Itoa(order[i])
		expectedUser := "fake234"
		if i >= 2 {
			expectedUser = "fake123"
		}

		if response.Feed.Posts[i].User.Name != expectedUser || response.Feed.Posts[i].Text != expectedText {
			t.Errorf("Received unexpected post response, '%v' != '%v' user and '%v' != '%v' text.\n", actual.User.Name, expectedUser, actual.Text, expectedText)
		}
	}
}
