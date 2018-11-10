package post

import "github.com/google/uuid"

type StubService struct {
	PostCache map[string]Post
}

func CreateStub() Service {
	stub := new(StubService)
	stub.PostCache = make(map[string]Post)
	return stub
}

func (stub *StubService) Create(request CreatePostRequest) (CreatePostResponse, error) {
	if request.Text == "" {
		return CreatePostResponse{}, &EmptyPostTextError{Text: request.Text}
	}

	// Store in the stubbed cache
	postID := uuid.New().String()
	stub.PostCache[postID] = Post{PostID: postID, User: request.UserID, Text: request.Text}

	// Create response
	response := CreatePostResponse{PostID: postID}

	return response, nil
}

func (stub *StubService) View(request ViewPostRequest) (ViewPostResponse, error) {
	post, ok := stub.PostCache[request.PostID]

	if !ok {
		return ViewPostResponse{}, &InvalidPostIDError{PostID: request.PostID}
	}

	return ViewPostResponse{Text: post.Text}, nil
}
