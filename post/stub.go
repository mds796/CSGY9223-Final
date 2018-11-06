package post

import "github.com/google/uuid"

type StubService struct {
	Cache map[string]Post
}

func CreateStub() Service {
	stub := new(StubService)
	stub.Cache = make(map[string]Post)
	return stub
}
func (stub *StubService) Create(request CreatePostRequest) (CreatePostResponse, error) {
	// Process request
	userID, _ := uuid.Parse(request.UserID)

	// Store in the stubbed cache
	postID := uuid.New()
	stub.Cache[postID.String()] = Post{postID: postID, userID: userID, text: request.Text}

	// Create response
	response := CreatePostResponse{PostID: postID.String()}

	return response, nil
}
