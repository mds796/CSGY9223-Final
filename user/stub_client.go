package user

import (
	"context"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"google.golang.org/grpc"
)

type StubClient struct {
	service userpb.UserServer
}

func (s StubClient) Create(ctx context.Context, in *userpb.CreateUserRequest, opts ...grpc.CallOption) (*userpb.CreateUserResponse, error) {
	return s.service.Create(ctx, in)
}

func (s StubClient) View(ctx context.Context, in *userpb.ViewUserRequest, opts ...grpc.CallOption) (*userpb.ViewUserResponse, error) {
	return s.service.View(ctx, in)
}

func (s StubClient) Search(ctx context.Context, in *userpb.SearchUserRequest, opts ...grpc.CallOption) (*userpb.SearchUserResponse, error) {
	return s.service.Search(ctx, in)
}
