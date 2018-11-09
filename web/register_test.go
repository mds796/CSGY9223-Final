package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_RegisterUser(t *testing.T) {
	service := CreateService()

	rr := registerUser(t, service)

	verifyLogin(rr, t)
}

func registerUser(t *testing.T, service *HttpService) *httptest.ResponseRecorder {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/register", strings.NewReader("username=fake123&password=1234567890&password2=1234567890"))
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.RegisterUser)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	return rr
}

func TestHttpService_RegisterUser_InvalidPassword(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/register", strings.NewReader("username=fake123&password=1234567890&password2=123"))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.RegisterUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusTemporaryRedirect)
	}

	// Check the response cookie is what we expect.
	cookies := rr.HeaderMap["Set-Cookie"]
	if len(cookies) != 1 || !match("error=\"Invalid register request.\"; Expires=", cookies) {
		t.Errorf("Handler did not set the error cookie correctly %v\n", rr.HeaderMap)
	}
}
