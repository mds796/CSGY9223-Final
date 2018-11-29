package user

import (
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

func TestUserCreateBasic(t *testing.T) {
	service := &StubClient{service: CreateStub()}

	request := CreateUserRequest{Username: "mksavic"}
	_, err := service.Create(request)

	if err != nil {
		t.Fail()
	}
}

func TestUserCreateExists(t *testing.T) {
	service := CreateStub()

	request := CreateUserRequest{Username: "mksavic"}
	service.Create(request)

	_, err := service.Create(request)

	if err == nil {
		t.Fail()
	}
}

func TestUserViewBasic(t *testing.T) {
	service := CreateStub()

	createRequest := CreateUserRequest{Username: "mksavic"}
	service.Create(createRequest)

	request := ViewUserRequest{Username: "mksavic"}
	_, err := service.View(request)

	if err != nil {
		t.Fail()
	}
}

func TestUserViewDoesNotExist(t *testing.T) {
	service := CreateStub()

	request := ViewUserRequest{Username: "mksavic"}
	_, err := service.View(request)

	if err == nil {
		t.Fail()
	}
}

func TestUserSearchBasic(t *testing.T) {
	service := CreateStub()

	createRequest := CreateUserRequest{Username: "mksavic"}
	service.Create(createRequest)

	request := SearchUserRequest{Query: "sav"}
	response, _ := service.Search(request)

	if len(response.Usernames) != 1 {
		t.Fail()
	}

	if !contains(response.Usernames, "mksavic") {
		t.Fail()
	}
}

func TestUserSearchDoesNotExist(t *testing.T) {
	service := CreateStub()

	request := SearchUserRequest{Query: "sav"}
	response, _ := service.Search(request)

	if len(response.Usernames) > 0 {
		t.Fail()
	}
}

func TestUserSearchMulti(t *testing.T) {
	service := CreateStub()

	createRequest := CreateUserRequest{Username: "mksavic"}
	service.Create(createRequest)

	createRequest = CreateUserRequest{Username: "mds796"}
	service.Create(createRequest)

	createRequest = CreateUserRequest{Username: "mvp307"}
	service.Create(createRequest)

	request := SearchUserRequest{Query: "s"}
	response, _ := service.Search(request)

	if len(response.Usernames) != 2 {
		t.Fail()
	}

	if !contains(response.Usernames, "mksavic") || !contains(response.Usernames, "mds796") {
		t.Fail()
	}
}
