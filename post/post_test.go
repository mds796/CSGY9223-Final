package post

import "testing"

import "github.com/google/uuid"

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

func TestViewReturnsErrorWithInvalidPostID(t *testing.T) {
	service := CreateStub()
	_, err := service.View(ViewPostRequest{PostID: "123"})
	_, ok := err.(*InvalidPostIDError)

	if !ok {
		t.Fail()
	}
}

func TestListReturnsAllPostsFromUser(t *testing.T) {
	service := CreateStub()
	userID := uuid.New().String()
	service.Create(CreatePostRequest{UserID: userID, Text: "post 1"})
	service.Create(CreatePostRequest{UserID: userID, Text: "post 2"})
	service.Create(CreatePostRequest{UserID: userID, Text: "post 3"})
	listResponse, _ := service.List(ListPostsRequest{UserID: userID})

	viewResponse, _ := service.View(ViewPostRequest{PostID: listResponse.PostIDs[0]})
	if viewResponse.Text != "post 1" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
	if viewResponse.Text != "post 2" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[2]})
	if viewResponse.Text != "post 3" {
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
	if viewResponse.Text != "post 1" {
		t.Fail()
	}

	viewResponse, _ = service.View(ViewPostRequest{PostID: listResponse.PostIDs[1]})
	if viewResponse.Text != "post 3" {
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
