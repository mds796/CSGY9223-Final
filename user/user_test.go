package user

import (
	"context"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"testing"
)

func contains(s []string, e string) bool {
	for _, val := range s {
		if val == e {
			return true
		}
	}
	return false
}

func createUserService() *StubClient {
	return &StubClient{service: CreateStub()}
}

func sendCreateUserRequest(client *StubClient, request *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return client.Create(context.Background(), request)
}

func sendViewUserRequest(client *StubClient, request *userpb.ViewUserRequest) (*userpb.ViewUserResponse, error) {
	return client.View(context.Background(), request)
}

func sendSearchUserRequest(client *StubClient, request *userpb.SearchUserRequest) (*userpb.SearchUserResponse, error) {
	return client.Search(context.Background(), request)
}

func TestUserCreateBasic(t *testing.T) {
	client := createUserService()

	createRequest := &userpb.CreateUserRequest{Username: "mksavic"}
	_, err := sendCreateUserRequest(client, createRequest)

	if err != nil {
		t.Fail()
	}
}

func TestUserCreateExists(t *testing.T) {
	client := createUserService()

	createRequest := &userpb.CreateUserRequest{Username: "mksavic"}
	sendCreateUserRequest(client, createRequest)

	_, err := sendCreateUserRequest(client, createRequest)

	if err == nil {
		t.Fail()
	}
}

func TestUserViewBasic(t *testing.T) {
	client := createUserService()

	createRequest := &userpb.CreateUserRequest{Username: "mksavic"}
	sendCreateUserRequest(client, createRequest)

	viewRequest := &userpb.ViewUserRequest{Username: "mksavic"}
	_, err := sendViewUserRequest(client, viewRequest)

	if err != nil {
		t.Fail()
	}
}

func TestUserViewDoesNotExist(t *testing.T) {
	client := createUserService()

	viewRequest := &userpb.ViewUserRequest{Username: "mksavic"}
	_, err := sendViewUserRequest(client, viewRequest)

	if err == nil {
		t.Fail()
	}
}

func TestUserSearchBasic(t *testing.T) {
	client := createUserService()

	createRequest := &userpb.CreateUserRequest{Username: "mksavic"}
	sendCreateUserRequest(client, createRequest)

	searchRequest := &userpb.SearchUserRequest{Query: "sav"}
	response, _ := sendSearchUserRequest(client, searchRequest)

	if len(response.Usernames) != 1 {
		t.Fail()
	}

	if !contains(response.Usernames, "mksavic") {
		t.Fail()
	}
}

func TestUserSearchDoesNotExist(t *testing.T) {
	client := createUserService()

	searchRequest := &userpb.SearchUserRequest{Query: "sav"}
	response, _ := sendSearchUserRequest(client, searchRequest)

	if len(response.Usernames) > 0 {
		t.Fail()
	}
}

func TestUserSearchMulti(t *testing.T) {
	client := createUserService()

	createRequest := &userpb.CreateUserRequest{Username: "mksavic"}
	sendCreateUserRequest(client, createRequest)

	createRequest = &userpb.CreateUserRequest{Username: "mds796"}
	sendCreateUserRequest(client, createRequest)

	createRequest = &userpb.CreateUserRequest{Username: "mvp307"}
	sendCreateUserRequest(client, createRequest)

	searchRequest := &userpb.SearchUserRequest{Query: "s"}
	response, _ := sendSearchUserRequest(client, searchRequest)

	if len(response.Usernames) != 2 {
		t.Fail()
	}

	if !contains(response.Usernames, "mksavic") || !contains(response.Usernames, "mds796") {
		t.Fail()
	}
}
