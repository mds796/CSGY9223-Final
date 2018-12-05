package auth

import (
	"bytes"
	"context"
	"crypto/sha256"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"log"
	"net/http"
	"time"
)

type StubService struct {
	UserService   userpb.UserClient
	PasswordCache storage.Storage // (UUID, sha256(password))
	StatusCache   storage.Storage // (UUID, status)
	CookieCache   storage.Storage // (username, cookie)
}

func DecodeCookie(cookie string) *http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}
	return request.Cookies()[0]
}

func CreateStub(storageType storage.StorageType, userService userpb.UserClient) *StubService {
	stub := new(StubService)
	stub.UserService = userService
	stub.PasswordCache = storage.CreateStorage(storageType)
	stub.StatusCache = storage.CreateStorage(storageType)
	stub.CookieCache = storage.CreateStorage(storageType)
	return stub
}

func (s *StubService) Register(ctx context.Context, request *authpb.RegisterAuthRequest) (*authpb.RegisterAuthResponse, error) {
	// create the user in user service
	createUserRequest := &userpb.CreateUserRequest{Username: request.Username}
	createUserResponse, err := s.UserService.Create(ctx, createUserRequest)

	// something went wrong in user service
	if err != nil {
		log.Printf("[AUTH] %v", err)
		return &authpb.RegisterAuthResponse{}, &RegisterAuthError{request.Username}
	}

	// register the user
	h := sha256.New()
	h.Write([]byte(request.Password))
	s.PasswordCache.Put(createUserResponse.UID, h.Sum(nil))
	s.StatusCache.Put(createUserResponse.UID, []byte("LOGGED_IN"))
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := (&http.Cookie{Name: request.Username, Value: createUserResponse.UID, Expires: expiration}).String()
	s.CookieCache.Put(request.Username, []byte(cookie))
	return &authpb.RegisterAuthResponse{Cookie: cookie}, nil
}

func (s *StubService) Login(ctx context.Context, request *authpb.LoginAuthRequest) (*authpb.LoginAuthResponse, error) {
	// view the user in user service
	viewUserRequest := &userpb.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(ctx, viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return &authpb.LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}

	// check user password
	h := sha256.New()
	h.Write([]byte(request.Password))
	password, _ := s.PasswordCache.Get(viewUserResponse.UID)
	if bytes.Equal(password, h.Sum(nil)) {
		// login the user, their current status is irrelevant
		s.StatusCache.Put(viewUserResponse.UID, []byte("LOGGED_IN"))
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := (&http.Cookie{Name: request.Username, Value: viewUserResponse.UID, Expires: expiration}).String()
		s.CookieCache.Put(request.Username, []byte(cookie))
		return &authpb.LoginAuthResponse{Cookie: cookie}, nil
	} else {
		return &authpb.LoginAuthResponse{}, &LoginAuthError{request.Username, request.Password}
	}
}

func (s *StubService) Verify(ctx context.Context, request *authpb.VerifyAuthRequest) (*authpb.VerifyAuthResponse, error) {
	// check if cookie is assigned to a username
	for username, cookie := range s.CookieCache.Iterate() {
		savedCookie := DecodeCookie(string(cookie))
		requestCookie := DecodeCookie(request.Cookie)
		if savedCookie.Name == requestCookie.Name && savedCookie.Path == requestCookie.Path && savedCookie.Value == requestCookie.Value {
			response := &authpb.VerifyAuthResponse{Username: username, UID: savedCookie.Value}
			return response, nil
		}
	}

	// cookie is not known to us
	return &authpb.VerifyAuthResponse{}, &VerifyAuthError{request.Cookie}
}

func (s *StubService) Logout(ctx context.Context, request *authpb.LogoutAuthRequest) (*authpb.LogoutAuthResponse, error) {
	// view the user in user service
	viewUserRequest := &userpb.ViewUserRequest{Username: request.Username}
	viewUserResponse, err := s.UserService.View(ctx, viewUserRequest)

	// something went wrong in user service
	if err != nil {
		return &authpb.LogoutAuthResponse{}, &LogoutAuthError{request.Username}
	}

	// logout the user, their current status is irrelevant
	s.StatusCache.Put(viewUserResponse.UID, []byte("LOGGED_OUT"))
	s.CookieCache.Delete(request.Username)
	return &authpb.LogoutAuthResponse{}, nil
}
