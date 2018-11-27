package post

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"testing"
)

func createPostService() *StubClient {
	return &StubClient{service: CreateStub()}
}

func TestValidCreatePostDoesNotReturnsError(t *testing.T) {
	client := createPostService()

	response, err := client.Create(
		context.Background(),
		&postpb.CreateRequest{
			User: &postpb.User{ID: uuid.New().String()},
			Post: &postpb.Post{Text: "testing"},
		})

	if err != nil {
		t.Fail()
	}

	_, err = uuid.Parse(response.Post.ID)

	if err != nil {
		t.Fail()
	}
}

func TestCreatePostReturnsValidUUID(t *testing.T) {
	client := createPostService()

	response, _ := client.Create(
		context.Background(),
		&postpb.CreateRequest{
			User: &postpb.User{ID: uuid.New().String()},
			Post: &postpb.Post{Text: "testing"},
		})
	_, err := uuid.Parse(response.Post.ID)

	if err != nil {
		t.Fail()
	}
}

func TestCreatePostReturnsErrorWithEmptyText(t *testing.T) {
	client := createPostService()

	_, err := client.Create(
		context.Background(),
		&postpb.CreateRequest{
			User: &postpb.User{ID: uuid.New().String()},
			Post: &postpb.Post{Text: ""},
		})
	_, ok := err.(*EmptyPostTextError)

	if !ok {
		t.Fail()
	}
}

// func TestViewReturnsTextAfterCreatingPost(t *testing.T) {
// 	client := createPostService()
// 	createResponse, _ := client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
// 	viewResponse, _ := client.View(ViewPostRequest{PostID: createResponse.PostID})

// 	if viewResponse.Text != "testing" {
// 		t.Fail()
// 	}
// }

// func TestViewReturnsTextFromCorrectPost(t *testing.T) {
// 	client := createPostService()
// 	client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
// 	createResponse, _ := client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing more"})
// 	viewResponse, _ := client.View(ViewPostRequest{PostID: createResponse.PostID})

// 	if viewResponse.Text != "testing more" {
// 		t.Fail()
// 	}
// }

// func TestCreatePostsHaveIncreasingTimestamps(t *testing.T) {
// 	client := createPostService()
// 	createResponse1, _ := client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 1"})
// 	createResponse2, _ := client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 2"})

// 	viewResponse1, _ := client.View(ViewPostRequest{PostID: createResponse1.PostID})
// 	viewResponse2, _ := client.View(ViewPostRequest{PostID: createResponse2.PostID})

// 	if viewResponse2.Timestamp.Before(viewResponse1.Timestamp) || viewResponse2.Timestamp.Equal(viewResponse1.Timestamp) {
// 		t.Fail()
// 	}
// }

// func TestViewReturnsErrorWithInvalidPostID(t *testing.T) {
// 	client := createPostService()
// 	_, err := client.View(ViewPostRequest{PostID: "123"})
// 	_, ok := err.(*InvalidPostIDError)

// 	if !ok {
// 		t.Fail()
// 	}
// }

// func TestListReturnsAllPostsFromUserInReverseOrder(t *testing.T) {
// 	client := createPostService()
// 	userID := uuid.New().String()
// 	client.Create(CreatePostRequest{UserID: userID, Text: "post 1"})
// 	client.Create(CreatePostRequest{UserID: userID, Text: "post 2"})
// 	client.Create(CreatePostRequest{UserID: userID, Text: "post 3"})
// 	listResponse, _ := client.List(ListPostsRequest{UserID: userID})

// 	viewResponse, _ := client.View(ViewPostRequest{PostID: listResponse.PostIDs[0]})
// 	if viewResponse.Text != "post 3" {
// 		t.Fail()
// 	}

// 	viewResponse, _ = client.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
// 	if viewResponse.Text != "post 2" {
// 		t.Fail()
// 	}

// 	viewResponse, _ = client.View(ViewPostRequest{PostID: listResponse.PostIDs[2]})
// 	if viewResponse.Text != "post 1" {
// 		t.Fail()
// 	}
// }

// func TestListReturnsPostsFromCorrectUser(t *testing.T) {
// 	client := createPostService()
// 	userID := uuid.New().String()
// 	client.Create(CreatePostRequest{UserID: userID, Text: "post 1"})
// 	client.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 2"})
// 	client.Create(CreatePostRequest{UserID: userID, Text: "post 3"})
// 	listResponse, _ := client.List(ListPostsRequest{UserID: userID})

// 	viewResponse, _ := client.View(ViewPostRequest{PostID: listResponse.PostIDs[0]})
// 	if viewResponse.Text != "post 3" {
// 		t.Fail()
// 	}

// 	viewResponse, _ = client.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
// 	if viewResponse.Text != "post 1" {
// 		t.Fail()
// 	}
// }

// func TestListReturnsEmptyPostsListWithUnknownUserID(t *testing.T) {
// 	client := createPostService()
// 	userID := uuid.New().String()
// 	listResponse, err := client.List(ListPostsRequest{UserID: userID})

// 	if err != nil {
// 		t.Fail()
// 	}

// 	if len(listResponse.PostIDs) != 0 {
// 		t.Fail()
// 	}
// }
