package web

import (
	"context"
	"encoding/json"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpService_FetchFeed(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	followUser(t, service, recorder)
	makePost(t, recorder, service)
	response, _ := service.UserService.View(context.Background(), &userpb.ViewUserRequest{Username: "fake234"})

	req, err := http.NewRequest("GET", "/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.FetchFeed)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	posts := make([]*feedpb.Post, 1)
	posts[0] = &feedpb.Post{Text: "Hello, World!", User: &feedpb.User{Name: "fake234", ID: response.UID}}
	expected := feedpb.ViewResponse{Feed: &feedpb.Feed{Posts: posts}}

	data := rr.Body.Bytes()
	var received feedpb.Feed
	_ = json.Unmarshal(data, &received)

	if len(received.Posts) != len(expected.Feed.Posts) {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received, expected)
	}

	for i := range received.Posts {
		if received.Posts[i].Text != expected.Feed.Posts[i].Text ||
			received.Posts[i].User.Name != expected.Feed.Posts[i].User.Name {
			t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received.Posts[i], expected.Feed.Posts[i])
		}
	}
}

func TestHttpService_FetchFeedEmpty(t *testing.T) {
	service := CreateService()

	recorder := registerUserWithName(t, service, "fake123")

	req, err := http.NewRequest("GET", "/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.FetchFeed)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	posts := make([]*feedpb.Post, 0)

	bytes, err := json.Marshal(feedpb.Feed{Posts: posts})
	if err != nil {
		t.Fatal(err)
	}
	expected := string(bytes)

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", rr.Body.String(), expected)
	}
}

func TestHttpService_FetchFeed_OtherUserPosts(t *testing.T) {
	service := CreateService()

	otherUserRecorder := registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	followUser(t, service, recorder)
	makePost(t, otherUserRecorder, service)
	response, _ := service.UserService.View(context.Background(), &userpb.ViewUserRequest{Username: "fake123"})

	req, err := http.NewRequest("GET", "/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.FetchFeed)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	posts := make([]*feedpb.Post, 1)
	posts[0] = &feedpb.Post{Text: "Hello, World!", User: &feedpb.User{Name: "fake123", ID: response.UID}}
	expected := &feedpb.Feed{Posts: posts}

	data := rr.Body.Bytes()
	var received feedpb.Feed
	_ = json.Unmarshal(data, &received)

	if len(received.Posts) != len(expected.Posts) {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received, expected)
	}

	for i := range received.Posts {
		if received.Posts[i].Text != expected.Posts[i].Text ||
			received.Posts[i].User.Name != expected.Posts[i].User.Name {
			t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received.Posts[i], expected.Posts[i])
		}
	}
}

func TestHttpService_FetchFeed_InvalidUser(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")

	req, err := http.NewRequest("GET", "/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.FetchFeed)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}
}
