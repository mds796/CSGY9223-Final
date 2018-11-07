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
