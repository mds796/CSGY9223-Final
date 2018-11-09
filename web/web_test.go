package web

import (
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestHttpService_Start(t *testing.T) {
	service := StartService()

	response, err := http.Get(service.Address() + "/index.html")

	if err != nil {
		t.Fatalf("The server did not start as expected: %v\n", err)
	} else if response.StatusCode != 200 {
		t.Fatalf("Could not fetch index page from server. Status: %d.\n", response.StatusCode)
	}
}

func TestHttpService_RegisterUser(t *testing.T) {
	service := StartService()

	response, err := http.Post(
		service.Address()+"/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("username=fake123&password=1234567890&password2=1234567890"))

	if err != nil {
		t.Fatalf("Could not register service: %v\n", err)
	}

	cookie := findCookie(response, "fake123")

	if response.StatusCode != 200 || cookie == nil {
		t.Fatalf("Could not fetch index page from server. Cookie: %v, Status: %d.\n", cookie, response.StatusCode)
	}
}

func findCookie(response *http.Response, name string) *http.Cookie {
	cookies := response.Cookies()

	for i := range cookies {
		log.Println(cookies[i])
		if cookies[i].Name == name {
			return cookies[i]
		}
	}

	return nil
}

func TestHttpService_RegisterUser_InvalidPassword(t *testing.T) {
	service := StartService()

	response, err := http.Post(
		service.Address()+"/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("username=fake123&password=1234567890&password2=123"))

	if err != nil {
		t.Fatalf("Could not register service: %v\n", err)
	} else if response.StatusCode != 307 {
		t.Fatalf("Could not fetch index page from server. Status: %d.\n", response.StatusCode)
	}
}

func StartService() Service {
	return New("localhost", 8080, "../static")
}
