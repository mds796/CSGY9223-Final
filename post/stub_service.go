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

func CreateService(storageType storage.StorageType) *StubService {
	stub := new(StubService)
	stub.PostCache = storage.CreateStorage(storageType, "post/post_cache")
	stub.UserPostsCache = storage.CreateStorage(storageType, "post/user_posts_cache")
	return stub
}

func (stub *StubService) Create(ctx context.Context, request *postpb.CreateRequest) (*postpb.CreateResponse, error) {
	if request.Post.Text == "" {
		return nil, &EmptyPostTextError{Text: request.Post.Text}
	}

	// create post object
	post := &postpb.Post{ID: uuid.New().String(), Text: request.Post.Text, User: request.User, Timestamp: generateTimestamp()}

	// cache new post
	postBytes, _ := proto.Marshal(post)
	stub.PostCache.Put(post.ID, postBytes)

	// get user's previous posts
	postsBytes, _ := stub.UserPostsCache.Get(request.User.ID)
	posts := &postpb.Posts{}
	proto.Unmarshal(postsBytes, posts)

	// prepend new post to user's previous posts
	updatedPosts := &postpb.Posts{Posts: prepend(posts.Posts, post)}
	updatedPostsBytes, _ := proto.Marshal(updatedPosts)
	stub.UserPostsCache.Put(request.User.ID, updatedPostsBytes)

	return &postpb.CreateResponse{Post: post}, nil
}

func prepend(slice []*postpb.Post, obj *postpb.Post) []*postpb.Post {
	return append([]*postpb.Post{obj}, slice...)
}

func generateTimestamp() *postpb.Timestamp {
	return &postpb.Timestamp{EpochNanoseconds: time.Now().UnixNano()}
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
	deserializedPosts := &postpb.Posts{}
	proto.Unmarshal(posts, deserializedPosts)
	return &postpb.ListResponse{Posts: deserializedPosts.Posts}, nil
}
