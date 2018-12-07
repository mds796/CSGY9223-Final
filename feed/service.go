package feed

import (
	"context"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"google.golang.org/grpc"
	"log"
	"sort"
)

type Service struct {
	Post   postpb.PostClient
	Follow followpb.FollowClient
	User   userpb.UserClient
}

func CreateService(postService postpb.PostClient, userService userpb.UserClient, followService followpb.FollowClient) *Service {
	service := new(Service)
	service.Post = postService
	service.Follow = followService
	service.User = userService
	return service
}

func (s Service) View(ctx context.Context, request *feedpb.ViewRequest) (*feedpb.ViewResponse, error) {
	followed, err := s.ListFollowed(request.User.ID)
	if err != nil {
		return nil, err
	}

	posts, err := s.ListPosts(followed)
	if err != nil {
		return nil, err
	}

	postsForUser := s.PostsForUser(request.User.ID, request.User.Name)
	posts = append(posts, postsForUser...)

	// Sort posts by timestamp
	// Caution: using timestamps depending on the computer's clock for ordering
	// posts won't work in a distributed system.
	// Must use a Lamport clock (monotonically increasing integer with consensus protocol)
	// to safely provide total ordering even with distributed processing.
	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Timestamp.EpochNanoseconds < posts[i].Timestamp.EpochNanoseconds
	})

	return &feedpb.ViewResponse{Feed: &feedpb.Feed{Posts: posts}}, nil
}

func (s Service) ListPosts(followed []*userpb.ViewUserResponse) ([]*feedpb.Post, error) {
	posts := make([]*feedpb.Post, 0, len(followed))

	for i := range followed {
		postsForUser := s.PostsForUser(followed[i].UID, followed[i].Username)
		posts = append(posts, postsForUser...)
	}

	return posts, nil
}

func (s Service) PostsForUser(userId string, username string) []*feedpb.Post {

	response, err := s.Post.List(context.Background(), &postpb.ListRequest{User: &postpb.User{ID: userId}})
	if err != nil {
		log.Printf("Encountered an error listing posts for user %v.\n", userId)
		return nil
	}

	posts := make([]*feedpb.Post, 0, len(response.Posts))

	for j := range response.Posts {
		postResponse, err := s.Post.View(context.Background(), &postpb.ViewRequest{Post: &postpb.Post{ID: response.Posts[j].ID}})
		if err != nil {
			log.Printf("Encountered an error viewing post %v.\n", response.Posts[j].ID)
		} else {
			postedBy := &feedpb.User{Name: username, ID: userId}
			postedAt := &feedpb.Timestamp{EpochNanoseconds: postResponse.Post.Timestamp.EpochNanoseconds}
			posts = append(posts, &feedpb.Post{User: postedBy, Text: postResponse.Post.Text, Timestamp: postedAt})
		}
	}

	return posts
}

func (s Service) ListFollowed(userId string) ([]*userpb.ViewUserResponse, error) {
	viewResponse, err := s.Follow.View(context.Background(), &followpb.ViewRequest{User: &followpb.User{ID: userId}})
	if err != nil {
		return nil, err
	}

	userIds := make([]*userpb.ViewUserResponse, 0, len(viewResponse.Users))

	for i := range viewResponse.Users {
		response, err := s.User.View(context.Background(), &userpb.ViewUserRequest{UID: viewResponse.Users[i].ID})

		if err != nil {
			log.Printf("Encountered an error viewing user %v.\n", viewResponse.Users[i].ID)
		} else {
			userIds = append(userIds, response)
		}
	}

	return userIds, nil
}

type StubClient struct {
	service feedpb.FeedServer
}

func (s StubClient) View(ctx context.Context, in *feedpb.ViewRequest, opts ...grpc.CallOption) (*feedpb.ViewResponse, error) {
	return s.service.View(ctx, in)
}
