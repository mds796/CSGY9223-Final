package web

import (
	"encoding/json"
	"github.com/mds796/CSGY9223-Final/feed"
	"github.com/mds796/CSGY9223-Final/post"
	"io/ioutil"
	"log"
	"net/http"
)

func (srv *HttpService) FetchFeed(w http.ResponseWriter, r *http.Request) {
	posts, err := srv.listFeedPosts(r)

	if err == nil {
		w.Write(posts)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (srv *HttpService) listFeedPosts(r *http.Request) ([]byte, error) {
	response, err := srv.verifyToken(r)
	if err != nil {
		return nil, err
	}

	viewResponse, err := srv.FeedService.View(&feed.ViewRequest{UserID: response.UserID, Username: response.Username})
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(viewResponse)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (srv *HttpService) MakePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := srv.createPost(r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (srv *HttpService) createPost(r *http.Request) error {
	response, err := srv.verifyToken(r)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	request := post.CreatePostRequest{UserID: response.UserID, Text: string(bytes)}

	_, err = srv.PostService.Create(request)

	return err
}
