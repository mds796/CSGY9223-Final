package web

import (
	"encoding/json"
	"github.com/mds796/CSGY9223-Final/feed"
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

	posts := make([]*feed.Post, 1)
	posts[0] = &feed.Post{Text: "Hello, World!", From: "fake234"}
	expected := feed.ViewResponse{Posts: posts}

	data := rr.Body.Bytes()
	var received feed.ViewResponse
	json.Unmarshal(data, &received)

	if len(received.Posts) != len(expected.Posts) {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received, expected)
	}

	for i := range received.Posts {
		if received.Posts[i].Text != expected.Posts[i].Text ||
			received.Posts[i].From != expected.Posts[i].From {
			t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received.Posts[i], expected.Posts[i])
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

	posts := make([]*feed.Post, 0)

	bytes, err := json.Marshal(feed.ViewResponse{Posts: posts})
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

	posts := make([]*feed.Post, 1)
	posts[0] = &feed.Post{Text: "Hello, World!", From: "fake123"}
	expected := feed.ViewResponse{Posts: posts}

	data := rr.Body.Bytes()
	var received feed.ViewResponse
	json.Unmarshal(data, &received)

	if len(received.Posts) != len(expected.Posts) {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", received, expected)
	}

	for i := range received.Posts {
		if received.Posts[i].Text != expected.Posts[i].Text ||
			received.Posts[i].From != expected.Posts[i].From {
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
