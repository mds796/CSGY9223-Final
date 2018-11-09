package web

import (
	"net/http"
	"testing"
)

func TestHttpService_Start(t *testing.T) {
	service := StartService()
	defer service.Stop()

	response, err := http.Get("http://localhost:9999/index.html")

	if err != nil {
		t.Fatalf("The server did not start as expected: %v\n", err)
	} else if response.StatusCode != 200 {
		t.Fatalf("Could not fetch index page from server. Status: %d.\n", response.StatusCode)
	}
}

func TestHttpService_RegisterUser(t *testing.T) {
	service := StartService()
	defer service.Stop()

	response, err := http.Get("http://localhost:9999/register")

	if err != nil {
		t.Fatalf("The server did not start as expected: %v\n", err)
	} else if response.StatusCode != 200 {
		t.Fatalf("Could not fetch index page from server. Status: %d.\n", response.StatusCode)
	}
}

func StartService() Service {
	service := New("localhost", 9999, "../static/")
	go service.Start()

	return service
}

func TestHttpService_Stop(t *testing.T) {
	service := StartService()
	service.Stop()

	if _, err := http.Get("http://localhost:9999/index.html"); err == nil {
		t.Fatalf("The server did not stop as expected.\n")
	}
}
