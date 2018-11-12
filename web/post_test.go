package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_MakePost(t *testing.T) {
	service := CreateService()

	recorder := registerUser(t, service)
	rr := makePost(t, recorder, service)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}
}

func makePost(t *testing.T, recorder *httptest.ResponseRecorder, service *HttpService) *httptest.ResponseRecorder {
	text := "Hello, World!"
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/post", strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	addCookies(recorder, req)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.MakePost)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	return rr
}

func TestHttpService_MakePost_InvalidUser(t *testing.T) {
	service := CreateService()

	text := "Hello, World!"

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/post", strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.MakePost)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}
}
