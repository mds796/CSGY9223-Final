package post

import "testing"

import "github.com/google/uuid"

func TestCreate(t *testing.T) {
	service := CreateStub()
	request := CreatePostRequest{UserID: uuid.New().String(), Text: "testing"}
	response, err := service.Create(request)

	if err != nil {
		t.Errorf("Error when creating post with user ID '%v' and data '%v'", request.UserID, request.Text)
	}

	_, err = uuid.Parse(response.PostID)

	if err != nil {
		t.Errorf("Invalid post ID in post/create response")
	}
}
