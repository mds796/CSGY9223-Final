package post

import (
	"context"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"time"
)

type StubService struct {
	PostCache      map[string]*postpb.Post
	UserPostsCache map[string][]*postpb.Post
}

func CreateStub() *StubService {
	stub := new(StubService)
	stub.PostCache = make(map[string]*postpb.Post)
	stub.UserPostsCache = make(map[string][]*postpb.Post)
	return stub
}

func (stub *StubService) Create(ctx context.Context, request *postpb.CreateRequest) (*postpb.CreateResponse, error) {
	if request.Post.Text == "" {
		return nil, &EmptyPostTextError{Text: request.Post.Text}
	}

	// Generate post
	post := &postpb.Post{ID: uuid.New().String(), Text: request.Post.Text, User: request.User, Timestamp: generateTimestamp()}

	// Cache it in posts and user cache
	stub.PostCache[post.ID] = post
	stub.UserPostsCache[request.User.ID] = prepend(stub.UserPostsCache[request.User.ID], post)

	return &postpb.CreateResponse{Post: post}, nil
}

func generateTimestamp() *postpb.Timestamp {
	return &postpb.Timestamp{EpochSeconds: time.Now().Unix()}
}

func prepend(slice []*postpb.Post, obj *postpb.Post) []*postpb.Post {
	return append([]*postpb.Post{obj}, slice...)
}

// func (stub *StubService) View(ctx context.Context, request *postpb.ViewRequest) (*postpb.ViewResponse, error) {
// 	post, ok := stub.PostCache[request.Post.ID]

// 	if !ok {
// 		return ViewPostResponse{}, &InvalidPostIDError{PostID: request.Post.ID}
// 	}

// 	return ViewPostResponse{Text: post.Text, Timestamp: post.Timestamp}, nil
// }

// func (stub *StubService) List(request ListPostsRequest) (ListPostsResponse, error) {
// 	postIDs, _ := stub.UserPostsCache[request.UserID]
// 	return ListPostsResponse{PostIDs: postIDs}, nil
// }
