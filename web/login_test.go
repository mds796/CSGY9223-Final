package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_LoginUser(t *testing.T) {
	service := CreateService()

	registerUser(t, service)
	rr := loginUser(t, service)

	verifyLogin(rr, t)
}

func loginUser(t *testing.T, service *HttpService) *httptest.ResponseRecorder {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/login", strings.NewReader("username=fake123&password=1234567890"))
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.LogInUser)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	return rr
}

func verifyLogin(rr *httptest.ResponseRecorder, t *testing.T) {
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}
	// Check the response body is what we expect.
	cookies := rr.Header()["Set-Cookie"]
	if len(cookies) != 3 {
		t.Errorf("Handler did not set the auth cookies correctly %v\n", rr.Header())
	}
	if !contains("error=; Expires=Thu, 01 Jan 1970 00:00:00 GMT", cookies) {
		t.Errorf("Handler did not remove the error cookie correctly %v\n", rr.Header())
	}
	if !match("username=fake123; Expires=", cookies) {
		t.Errorf("Handler did not set the username cookie correctly %v\n", rr.Header())
	}
	if !match("fake123=", cookies) {
		t.Errorf("Handler did not set the auth token cookie correctly %v\n", rr.Header())
	}
}

func TestHttpService_LoginUser_DoesNotExist(t *testing.T) {
	service := CreateService()

	rr := loginUser(t, service)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusTemporaryRedirect)
	}

	// Check the response cookie is what we expect.
	cookies := rr.Header()["Set-Cookie"]
	if len(cookies) != 1 || !match("error=\"Invalid login request.\"; Expires=", cookies) {
		t.Errorf("Handler did not set the error cookie correctly %v\n", rr.Header())
	}
}
