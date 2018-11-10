package web

import (
	"encoding/json"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/user"
	"net/http"
)

func (srv *HttpService) ToggleFollow(w http.ResponseWriter, r *http.Request) {
	err := srv.toggleFollowStatus(r)

	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (srv *HttpService) toggleFollowStatus(r *http.Request) error {
	response, err := srv.verifyToken(r)
	if err != nil {
		return err
	}

	values, err := getParameters(r.Body)
	if err != nil {
		return err
	}

	followedUser, err := getKey(values, "name")
	if err != nil {
		return err
	}

	followStatus, err := getKey(values, "follow")
	if err != nil {
		return err
	}

	userResponse, err := srv.UserService.View(user.ViewUserRequest{Username: followedUser})
	if err != nil {
		return err
	}

	if followStatus == "true" {
		_, err = srv.FollowService.Follow(follow.FollowRequest{FollowerUserID: response.UserID, FollowedUserID: userResponse.Uuid})
	} else {
		_, err = srv.FollowService.Unfollow(follow.UnfollowRequest{FollowerUserID: response.UserID, FollowedUserID: userResponse.Uuid})
	}

	return err
}

func (srv *HttpService) ListFollows(w http.ResponseWriter, r *http.Request) {
	follows, err := srv.listUsersWithFollowStatus(r)

	if err == nil {
		w.Write(follows)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (srv *HttpService) listUsersWithFollowStatus(r *http.Request) ([]byte, error) {
	response, err := srv.verifyToken(r)
	if err != nil {
		return nil, err
	}

	values, err := getParameters(r.Body)
	if err != nil {
		return nil, err
	}

	query, err := getKey(values, "query")
	if err != nil {
		return nil, err
	}

	userResponse, err := srv.UserService.Search(user.SearchUserRequest{Query: query})
	if err != nil {
		return nil, err
	}

	viewResponse, err := srv.FollowService.View(follow.ViewRequest{UserID: response.UserID})
	if err != nil {
		return nil, err
	}

	followsCache := make(map[string]Follow, len(userResponse.UserIDs))
	follows := make([]Follow, 0, len(userResponse.UserIDs))

	for i := range userResponse.UserIDs {
		id := userResponse.UserIDs[i]
		name := userResponse.Usernames[i]
		data := Follow{name: name, follow: false}

		followsCache[id] = data
		follows = append(follows, data)
	}

	for i := range viewResponse.UserIDs {
		data, ok := followsCache[viewResponse.UserIDs[i]]
		if ok {
			data.follow = true
		}
	}

	bytes, err := json.Marshal(follows)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Follow is a data transfer object
type Follow struct {
	name   string
	follow bool
}
