package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpService_Static(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ServeStatic)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() == "" {
		t.Errorf("Handler returned empty body.\n")
	}
}

func TestHttpService_StaticNonRoot(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/service-worker.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ServeStatic)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "console.info('Service worker disabled for development, will be generated at build time.');"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got '%v' want '%v'\n", rr.Body.String(), expected)
	}
}

func TestHttpService_RegisterUser(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
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

	if !match("username=fake123; Expires=", cookies) {
		t.Errorf("Handler did not set the username cookie correctly %v\n", rr.HeaderMap)
	}

	if !match("fake123=", cookies) {
		t.Errorf("Handler did not set the auth token cookie correctly %v\n", rr.HeaderMap)
	}
}

func contains(value string, values []string) bool {
	for i := range values {
		if values[i] == value {
			return true
		}
	}

	return false
}

func match(value string, values []string) bool {
	for i := range values {
		if strings.Contains(values[i], value) {
			return true
		}
	}

	return false
}

func TestHttpService_RegisterUser_InvalidPassword(t *testing.T) {
	service := CreateService()

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
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
		t.Errorf("Handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	cookies := rr.HeaderMap["Set-Cookie"]
	if len(cookies) != 1 || !match("error=\"Invalid register request.\"; Expires=", cookies) {
		t.Errorf("Handler did not set the error cookie correctly %v\n", rr.HeaderMap)
	}
}

func CreateService() *HttpService {
	return newService("localhost", 9999, "../static")
}
