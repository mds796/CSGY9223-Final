package post

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"testing"
)

func createPostService() *StubClient {
	return &StubClient{service: CreateStub(storage.STUB)}
}

func doCreateRequest(client *StubClient, userID string, text string) (*postpb.CreateResponse, error) {
	return client.Create(
		context.Background(),
		&postpb.CreateRequest{
			User: &postpb.User{ID: userID},
			Post: &postpb.Post{Text: text},
		})
}

func doViewRequest(client *StubClient, postID string) (*postpb.ViewResponse, error) {
	return client.View(
		context.Background(),
		&postpb.ViewRequest{
			Post: &postpb.Post{ID: postID},
		})
}

func doListRequest(client *StubClient, userID string) (*postpb.ListResponse, error) {
	return client.List(
		context.Background(),
		&postpb.ListRequest{
			User: &postpb.User{ID: userID},
		})
}

func TestValidCreatePostDoesNotReturnsError(t *testing.T) {
	client := createPostService()
	response, err := doCreateRequest(client, uuid.New().String(), "testing")
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
	response, _ := doCreateRequest(client, uuid.New().String(), "testing")
	_, err := uuid.Parse(response.Post.ID)
	if err != nil {
		t.Fail()
	}
}

func TestCreatePostReturnsErrorWithEmptyText(t *testing.T) {
	client := createPostService()
	_, err := doCreateRequest(client, uuid.New().String(), "")
	_, ok := err.(*EmptyPostTextError)
	if !ok {
		t.Fail()
	}
}

func TestViewReturnsTextAfterCreatingPost(t *testing.T) {
	client := createPostService()
	createResponse, _ := doCreateRequest(client, uuid.New().String(), "testing")
	viewResponse, _ := doViewRequest(client, createResponse.Post.ID)
	if viewResponse.Post.Text != "testing" {
		t.Fail()
	}
}

func TestViewReturnsTextFromCorrectPost(t *testing.T) {
	client := createPostService()
	doCreateRequest(client, uuid.New().String(), "testing")
	createResponse, _ := doCreateRequest(client, uuid.New().String(), "testing more")
	viewResponse, _ := doViewRequest(client, createResponse.Post.ID)
	if viewResponse.Post.Text != "testing more" {
		t.Fail()
	}
}

func TestCreatedPostsHaveIncreasingTimestamps(t *testing.T) {
	client := createPostService()
	createResponse1, _ := doCreateRequest(client, uuid.New().String(), "post 1")
	createResponse2, _ := doCreateRequest(client, uuid.New().String(), "post 2")
	viewResponse1, _ := doViewRequest(client, createResponse1.Post.ID)
	viewResponse2, _ := doViewRequest(client, createResponse2.Post.ID)
	if viewResponse1.Post.Timestamp.EpochNanoseconds >= viewResponse2.Post.Timestamp.EpochNanoseconds {
		t.Fail()
	}
}

func TestViewReturnsErrorWithInvalidPostID(t *testing.T) {
	client := createPostService()
	_, err := doViewRequest(client, "123")
	_, ok := err.(*InvalidPostIDError)
	if !ok {
		t.Fail()
	}
}

func TestListReturnsAllPostsFromUserInReverseOrder(t *testing.T) {
	client := createPostService()

	userID := uuid.New().String()
	doCreateRequest(client, userID, "post 1")
	doCreateRequest(client, userID, "post 2")
	doCreateRequest(client, userID, "post 3")
	listResponse, _ := doListRequest(client, userID)

	viewResponse, _ := doViewRequest(client, listResponse.Posts[0].ID)
	if viewResponse.Post.Text != "post 3" {
		t.Fail()
	}

	viewResponse, _ = doViewRequest(client, listResponse.Posts[1].ID)
	if viewResponse.Post.Text != "post 2" {
		t.Fail()
	}

	viewResponse, _ = doViewRequest(client, listResponse.Posts[2].ID)
	if viewResponse.Post.Text != "post 1" {
		t.Fail()
	}
}

func TestListReturnsPostsFromCorrectUser(t *testing.T) {
	client := createPostService()
	userID := uuid.New().String()

	doCreateRequest(client, userID, "post 1")
	doCreateRequest(client, uuid.New().String(), "post 2")
	doCreateRequest(client, userID, "post 3")
	listResponse, _ := doListRequest(client, userID)

	viewResponse, _ := doViewRequest(client, listResponse.Posts[0].ID)
	if viewResponse.Post.Text != "post 3" {
		t.Fail()
	}

	viewResponse, _ = doViewRequest(client, listResponse.Posts[1].ID)
	if viewResponse.Post.Text != "post 1" {
		t.Fail()
	}
}

func TestListReturnsEmptyPostsListWithUnknownUserID(t *testing.T) {
	client := createPostService()
	listResponse, err := doListRequest(client, uuid.New().String())

	if err != nil {
		t.Fail()
	}

	if len(listResponse.Posts) != 0 {
		t.Fail()
	}
}
