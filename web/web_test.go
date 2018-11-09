package web

import (
	"net/http"
	"strings"
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

	response, err := http.Post(
		"http://localhost:9999/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("username=fake123&password=1234567890&password2=1234567890"))

	if err != nil {
		t.Fatalf("Could not register service: %v\n", err)
	} else if response.StatusCode != 307 {
		t.Fatalf("Could not fetch index page from server. Status: %d.\n", response.StatusCode)
	}
}

func TestHttpService_RegisterUser_InvalidPassword(t *testing.T) {
	service := StartService()
	defer service.Stop()

	response, err := http.Post(
		"http://localhost:9999/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("username=fake123&password=1234567890&password2=123"))

	if err != nil {
		t.Fatalf("Could not register service: %v\n", err)
	} else if response.StatusCode != 307 {
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
