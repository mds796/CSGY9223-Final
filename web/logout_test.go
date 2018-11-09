package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_LogoutUser(t *testing.T) {
	service := CreateService()

	registerRr := registerUser(t, service)

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	registerCookies := registerRr.HeaderMap["Set-Cookie"]
	for i := range registerCookies {
		req.Header.Add("Cookie", registerCookies[i])
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.LogOutUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	verifyLogout(rr, t)
}

func verifyLogout(rr *httptest.ResponseRecorder, t *testing.T) {
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}
	// Check the response body is what we expect.
	cookies := rr.HeaderMap["Set-Cookie"]
	if len(cookies) != 3 {
		t.Errorf("Handler did not set the auth cookies correctly %v\n", rr.HeaderMap)
	}
	if !contains("error=; Expires=Thu, 01 Jan 1970 00:00:00 GMT", cookies) {
		t.Errorf("Handler did not remove the error cookie correctly %v\n", rr.HeaderMap)
	}
	if !contains("error=; Expires=Thu, 01 Jan 1970 00:00:00 GMT", cookies) {
		t.Errorf("Handler did not remove the username cookie correctly %v\n", rr.HeaderMap)
	}
	if !contains("fake123=; Expires=Thu, 01 Jan 1970 00:00:00 GMT", cookies) {
		t.Errorf("Handler did not remove the auth token cookie correctly %v\n", rr.HeaderMap)
	}
}

func TestHttpService_LogoutUser_NotLoggedIn(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/login", strings.NewReader("username=fake123&password=1234567890"))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.LogOutUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusTemporaryRedirect)
	}

	// Check the response cookie is what we expect.
	cookies := rr.HeaderMap["Set-Cookie"]
	if len(cookies) != 1 || !match("error=\"Invalid logout request.\"; Expires=", cookies) {
		t.Errorf("Handler did not set the error cookie correctly %v\n", rr.HeaderMap)
	}
}
