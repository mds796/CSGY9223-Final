package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpService_ListFollows(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	followUser(t, service, recorder)

	req, err := http.NewRequest("GET", "/follows?query=fake", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ListFollows)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	follows := make([]Follow, 1)
	follows[0].Name = "fake123"
	follows[0].Follow = true

	bytes, err := json.Marshal(follows)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(bytes)

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", rr.Body.String(), expected)
	}
}

func TestHttpService_ListFollows_NotFollowing(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	req, err := http.NewRequest("GET", "/follows?query=fake", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ListFollows)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	follows := make([]Follow, 1)
	follows[0].Name = "fake123"

	bytes, err := json.Marshal(follows)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(bytes)

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", rr.Body.String(), expected)
	}
}

func TestHttpService_ListFollows_EmptyQuery(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	followUser(t, service, recorder)

	req, err := http.NewRequest("GET", "/follows?query=", nil)
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ListFollows)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	expected := "[]"

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", rr.Body.String(), expected)
	}
}
