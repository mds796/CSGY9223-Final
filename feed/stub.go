package feed

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/user"
	"log"
)

type StubService struct {
	Post   post.Service
	Follow follow.Service
	User   user.Service
}

func (s StubService) View(request *ViewRequest) (*ViewResponse, error) {
	followed, err := s.ListFollowed(request.UserID)
	if err != nil {
		return nil, err
	}

	posts, err := s.ListPosts(followed)
	if err != nil {
		return nil, err
	}

	postsForUser := s.PostsForUser(request.UserID, request.Username)
	posts = append(posts, postsForUser...)

	return &ViewResponse{Posts: posts}, nil
}

func (s StubService) ListPosts(followed []*user.ViewUserResponse) ([]*Post, error) {
	posts := make([]*Post, 0, len(followed))

	for i := range followed {
		postsForUser := s.PostsForUser(followed[i].Uuid, followed[i].Username)
		posts = append(posts, postsForUser...)
	}

	return posts, nil
}

func (s StubService) PostsForUser(userId string, username string) []*Post {

	response, err := s.Post.List(post.ListPostsRequest{UserID: userId})
	if err != nil {
		log.Printf("Encountered an error listing posts for user %v.\n", userId)
		return nil
	}

	posts := make([]*Post, 0, len(response.PostIDs))

	for j := range response.PostIDs {
		postResponse, err := s.Post.View(post.ViewPostRequest{PostID: response.PostIDs[j]})
		if err != nil {
			log.Printf("Encountered an error viewing post %v.\n", response.PostIDs[j])
		} else {
			posts = append(posts, &Post{From: username, Text: postResponse.Text})
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
