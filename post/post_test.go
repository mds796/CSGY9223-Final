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
