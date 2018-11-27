package feed

import (
	"context"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/user"
	"google.golang.org/grpc"
	"log"
	"sort"
)

type StubService struct {
	Post   postpb.PostClient
	Follow follow.Service
	User   user.Service
}

func (s StubService) View(ctx context.Context, request *feedpb.ViewRequest) (*feedpb.ViewResponse, error) {
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

func (s StubService) ListPosts(followed []*user.ViewUserResponse) ([]*feedpb.Post, error) {
	posts := make([]*feedpb.Post, 0, len(followed))

	for i := range followed {
		postsForUser := s.PostsForUser(followed[i].Uuid, followed[i].Username)
		posts = append(posts, postsForUser...)
	}

	return posts, nil
}

func (s StubService) PostsForUser(userId string, username string) []*feedpb.Post {

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

func (s StubService) ListFollowed(userId string) ([]*user.ViewUserResponse, error) {
	viewResponse, err := s.Follow.View(follow.ViewRequest{UserID: userId})
	if err != nil {
		return nil, err
	}

	userIds := make([]*user.ViewUserResponse, 0, len(viewResponse.UserIDs))

	for i := range viewResponse.UserIDs {
		response, err := s.User.View(user.ViewUserRequest{UserID: viewResponse.UserIDs[i]})

		if err != nil {
			log.Printf("Encountered an error viewing user %v.\n", viewResponse.UserIDs[i])
		} else {
			userIds = append(userIds, &response)
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
