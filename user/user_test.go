package user

import (
	"testing"
)

func TestUserCreateStandard(t *testing.T) {
	service := CreateStub()
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

func TestUserViewStandard(t *testing.T) {
	service := CreateStub()
	create_request := CreateUserRequest{Username: "mksavic"}
	create_response, _ := service.Create(create_request)

	request := ViewUserRequest{Uuid: create_response.Uuid}
	_, err := service.View(request)

	if err != nil {
		t.Fail()
	}
}

func TestUserViewDoesNotExist(t *testing.T) {
	service := CreateStub()
	request := ViewUserRequest{Uuid: "123456"}
	_, err := service.View(request)

	if err == nil {
		t.Fail()
	}
}

func TestUserSearchStandard(t *testing.T) {
	service := CreateStub()
	create_request := CreateUserRequest{Username: "mksavic"}
	service.Create(create_request)

	request := SearchUserRequest{Query: "sav"}
	response, _ := service.Search(request)

	if len(response.Uuids) != 1 {
		t.Fail()
	}
}

func TestUserSearchDoesNotExist(t *testing.T) {
	service := CreateStub()
	request := SearchUserRequest{Query: "sav"}
	response, _ := service.Search(request)

	if len(response.Uuids) > 0 {
		t.Fail()
	}
}

func TestUserSearchMulti(t *testing.T) {
	service := CreateStub()
	create_request := CreateUserRequest{Username: "mksavic"}
	service.Create(create_request)
	create_request = CreateUserRequest{Username: "mds796"}
	service.Create(create_request)
	create_request = CreateUserRequest{Username: "mvp307"}
	service.Create(create_request)

	request := SearchUserRequest{Query: "s"}
	response, _ := service.Search(request)

	if len(response.Uuids) != 2 {
		t.Fail()
	}
}