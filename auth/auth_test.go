package auth

import (
	"testing"
)

func TestAuthRegisterStandard(t *testing.T) {
	service := CreateStub()
	request := RegisterAuthRequest{Uuid: "123456", Password: "abc123"}
	_, err := service.Register(request)

	if err != nil {
		t.Fail()
	}
}

func TestAuthRegisterExists(t *testing.T) {
	service := CreateStub()
	request := RegisterAuthRequest{Uuid: "123456", Password: "abc123"}
	service.Register(request)
	_, err := service.Register(request)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginStandard(t *testing.T) {
	service := CreateStub()
	register_request := RegisterAuthRequest{Uuid: "123456", Password: "abc123"}
	service.Register(register_request)

	request := LoginAuthRequest{Uuid: "123456", Password: "abc123"}
	_, err := service.Login(request)

	if err != nil {
		t.Fail()
	}
}

func TestAuthLoginDoesNotExist(t *testing.T) {
	service := CreateStub()
	request := LoginAuthRequest{Uuid: "123456", Password: "abc123"}
	_, err := service.Login(request)

	if err == nil {
		t.Fail()
	}
}

func TestAuthLoginPasswordIncorrect(t *testing.T) {
	service := CreateStub()
	register_request := RegisterAuthRequest{Uuid: "123456", Password: "abc123"}
	service.Register(register_request)

	request := LoginAuthRequest{Uuid: "123456", Password: "123abc"}
	_, err := service.Login(request)

	if err == nil {
		t.Fail()
	}
}
