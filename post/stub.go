package post

import "github.com/google/uuid"

type StubService struct {
	PostCache      map[string]Post
	UserPostsCache map[string][]string
}

func CreateStub() Service {
	stub := new(StubService)
	stub.PostCache = make(map[string]Post)
	stub.UserPostsCache = make(map[string][]string)
	return stub
}

func (stub *StubService) Create(request CreatePostRequest) (CreatePostResponse, error) {
	if request.Text == "" {
		return CreatePostResponse{}, &EmptyPostTextError{Text: request.Text}
	}

	// Store in the stubbed cache
	postID := uuid.New().String()
	stub.PostCache[postID] = Post{PostID: postID, User: request.UserID, Text: request.Text}

	// Store in user posts cache as well
	stub.UserPostsCache[request.UserID] = append(stub.UserPostsCache[request.UserID], postID)

	return CreatePostResponse{PostID: postID}, nil
}

func (stub *StubService) View(request ViewPostRequest) (ViewPostResponse, error) {
	post, ok := stub.PostCache[request.PostID]

	if !ok {
		return ViewPostResponse{}, &InvalidPostIDError{PostID: request.PostID}
	}

	return ViewPostResponse{Text: post.Text}, nil
}

func (stub *StubService) List(request ListPostsRequest) (ListPostsResponse, error) {
	postIDs, _ := stub.UserPostsCache[request.UserID]
	return ListPostsResponse{PostIDs: postIDs}, nil
}
