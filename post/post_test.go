package post

import "testing"

import "github.com/google/uuid"

func TestCreatePostBasic(t *testing.T) {
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

func TestViewReturnTextAfterCreatingPost(t *testing.T) {
	service := CreateStub()
	createResponse, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
	viewResponse, _ := service.View(CreateViewRequest{PostID: createResponse.PostID})

	if viewResponse.Text != "testing" {
		t.Fail()
	}
}

func TestViewReturnTextFromCorrectPost(t *testing.T) {
	service := CreateStub()
	service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing"})
	createResponse, _ := service.Create(CreatePostRequest{UserID: uuid.New().String(), Text: "testing more"})
	viewResponse, _ := service.View(CreateViewRequest{PostID: createResponse.PostID})

	if viewResponse.Text != "testing more" {
		t.Fail()
	}
}

func TestViewReturnErrorWithInvalidPostID(t *testing.T) {
	service := CreateStub()
	_, err := service.View(CreateViewRequest{PostID: "123"})
	_, ok := err.(*InvalidPostIDError)

	if !ok {
		t.Fail()
	}
}
