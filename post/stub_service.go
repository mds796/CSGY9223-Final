package post

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"time"
)

type StubService struct {
	PostCache      storage.Storage
	UserPostsCache storage.Storage
}

func CreateStub(storageType storage.StorageType) *StubService {
	stub := new(StubService)
	stub.PostCache = storage.CreateStorage(storageType)
	stub.UserPostsCache = storage.CreateStorage(storageType)
	return stub
}

func (stub *StubService) Create(ctx context.Context, request *postpb.CreateRequest) (*postpb.CreateResponse, error) {
	if request.Post.Text == "" {
		return nil, &EmptyPostTextError{Text: request.Post.Text}
	}

	// Generate post
	post := &postpb.Post{ID: uuid.New().String(), Text: request.Post.Text, User: request.User, Timestamp: generateTimestamp()}

	// Cache it in posts and user cache
	postBytes, _ := proto.Marshal(post)
	stub.PostCache.Put(post.ID, postBytes)
	posts, _ := stub.UserPostsCache.Get(request.User.ID)
	stub.UserPostsCache.Put(request.User.ID, prepend(posts, postBytes))

	return &postpb.CreateResponse{Post: post}, nil
}

func generateTimestamp() *postpb.Timestamp {
	return &postpb.Timestamp{EpochNanoseconds: time.Now().UnixNano()}
}

func prepend(slice []byte, obj []byte) []byte {
	return append(obj, slice...)
}

func (stub *StubService) View(ctx context.Context, request *postpb.ViewRequest) (*postpb.ViewResponse, error) {
	post, err := stub.PostCache.Get(request.Post.ID)

	if err != nil {
		return nil, &InvalidPostIDError{PostID: request.Post.ID}
	}

	deserializedPost := &postpb.Post{}
	proto.Unmarshal(post, deserializedPost)
	return &postpb.ViewResponse{Post: deserializedPost}, nil
}

func (stub *StubService) List(ctx context.Context, request *postpb.ListRequest) (*postpb.ListResponse, error) {
	posts, _ := stub.UserPostsCache.Get(request.User.ID)
	deserializedPosts := []postpb.Post{}
	proto.Unmarshal(posts, deserializedPosts)
	return &postpb.ListResponse{Posts: deserializedPosts}, nil
}
