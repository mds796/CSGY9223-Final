package web

import (
	"net/http"
	"net/http/httptest"
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
