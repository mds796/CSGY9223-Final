package post

import (
	"github.com/google/uuid"
	"testing"
)

func TestValidCreatePostDoesNotReturnsError(t *testing.T) {
	service := CreateStub()
	response, err := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})

	if err != nil {
		t.Fail()
	}

	_, err = uuid.Parse(response.PostID)

	if err != nil {
		t.Fail()
	}
}

func TestCreatePostReturnsValidUUID(t *testing.T) {
	service := CreateStub()
	response, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
	_, err := uuid.Parse(response.PostID)

	if err != nil {
		t.Fail()
	}
}

func TestCreatePostReturnsErrorWithEmptyText(t *testing.T) {
	service := CreateStub()
	_, err := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: ""})
	_, ok := err.(*EmptyPostTextError)

	if !ok {
		t.Fail()
	}
}

func TestViewReturnsTextAfterCreatingPost(t *testing.T) {
	service := CreateStub()
	createResponse, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
	viewResponse, _ := service.View(ViewPostRequest{PostID: createResponse.PostID})

	if viewResponse.Text != "testing" {
		t.Fail()
	}
}

func TestViewReturnsTextFromCorrectPost(t *testing.T) {
	service := CreateStub()
	service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
	createResponse, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing more"})
	viewResponse, _ := service.View(ViewPostRequest{PostID: createResponse.PostID})

	if viewResponse.Text != "testing more" {
		t.Fail()
	}
}

func TestCreatePostsHaveIncreasingTimestamps(t *testing.T) {
	service := CreateStub()
	createResponse1, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 1"})
	createResponse2, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 2"})

	viewResponse1, _ := service.View(ViewPostRequest{PostID: createResponse1.PostID})
	viewResponse2, _ := service.View(ViewPostRequest{PostID: createResponse2.PostID})

	if viewResponse2.Timestamp.Before(viewResponse1.Timestamp) || viewResponse2.Timestamp.Equal(viewResponse1.Timestamp) {
		t.Fail()
	}
}

func TestViewReturnsErrorWithInvalidPostID(t *testing.T) {
	service := CreateStub()
	_, err := service.View(ViewPostRequest{PostID: "123"})
	_, ok := err.(*InvalidPostIDError)

	if !ok {
		t.Fail()
	}
}

func TestListReturnsAllPostsFromUserInReverseOrder(t *testing.T) {
	service := CreateStub()
	userID := uuid.New().String()
	service.Create(CreatePostRequest{UserID: userID, Text: "post 1"})
	service.Create(CreatePostRequest{UserID: userID, Text: "post 2"})
	service.Create(CreatePostRequest{UserID: userID, Text: "post 3"})
	listResponse, _ := service.List(ListPostsRequest{UserID: userID})

	viewResponse, _ := service.View(ViewPostRequest{PostID: listResponse.PostIDs[0]})
	if viewResponse.Text != "post 3" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
	if viewResponse.Text != "post 2" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[2]})
	if viewResponse.Text != "post 1" {
		t.Fail()
	}
}

func TestListReturnsPostsFromCorrectUser(t *testing.T) {
	service := CreateStub()
	userID := uuid.New().String()
	service.Create(CreatePostRequest{UserID: userID, Text: "post 1"})
	service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "post 2"})
	service.Create(CreatePostRequest{UserID: userID, Text: "post 3"})
	listResponse, _ := service.List(ListPostsRequest{UserID: userID})

	viewResponse, _ := service.View(ViewPostRequest{PostID: listResponse.PostIDs[0]})
	if viewResponse.Text != "post 3" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
	if viewResponse.Text != "post 1" {
		t.Fail()
	}
}

func TestListReturnsEmptyPostsListWithUnknownUserID(t *testing.T) {
	service := CreateStub()
	userID := uuid.New().String()
	listResponse, err := service.List(ListPostsRequest{UserID: userID})

	if err != nil {
		t.Fail()
	}

	if len(listResponse.PostIDs) != 0 {
		t.Fail()
	}
}
