package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_ToggleFollow(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	rr := followUser(t, service, recorder)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}
}

func followUser(t *testing.T, service *HttpService, recorder *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/follow", strings.NewReader("name=fake123&follow=true"))
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ToggleFollow)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	return rr
}

func addCookies(recorder *httptest.ResponseRecorder, req *http.Request) {
	registerCookies := recorder.Header()["Set-Cookie"]

	for i := range registerCookies {
		req.Header.Add("Cookie", registerCookies[i])
	}
}

func TestHttpService_ToggleFollow_NotFollowed(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	recorder := registerUserWithName(t, service, "fake234")

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/follow", strings.NewReader("name=fake123&follow=false"))
	if err != nil {
		t.Fatal(err)
	}

	addCookies(recorder, req)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ToggleFollow)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}
}

func TestHttpService_ToggleFollow_UserNotValid(t *testing.T) {
	service := CreateService()

	recorder := registerUserWithName(t, service, "fake234")
	rr := followUser(t, service, recorder)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}
}

func TestHttpService_ToggleFollow_NotLoggedIn(t *testing.T) {
	service := CreateService()

	registerUserWithName(t, service, "fake123")
	registerUserWithName(t, service, "fake234")

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/follow", strings.NewReader("name=fake123&follow=true"))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ToggleFollow)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}
}
