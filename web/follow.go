package web

import (
	"encoding/json"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/user"
	"io/ioutil"
	"log"
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
	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		follow := new(Follow)
		json.Unmarshal(bytes, follow)
		log.Printf("%v", follow)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
								{ 
									"follows":[
					                	{"name": "fake123", "followed": true},
					                	{"name": "fake234", "followed": false}
					            	]
								}
								`))
}
