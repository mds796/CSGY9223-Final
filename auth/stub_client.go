package auth

import (
	"context"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"google.golang.org/grpc"
)

type StubClient struct {
	service authpb.AuthServer
}

func (s StubClient) Register(ctx context.Context, in *authpb.RegisterAuthRequest, opts ...grpc.CallOption) (*authpb.RegisterAuthResponse, error) {
	return s.service.Register(ctx, in)
}

func (s StubClient) Login(ctx context.Context, in *authpb.LoginAuthRequest, opts ...grpc.CallOption) (*authpb.LoginAuthResponse, error) {
	return s.service.Login(ctx, in)
}

func (s StubClient) Verify(ctx context.Context, in *authpb.VerifyAuthRequest, opts ...grpc.CallOption) (*authpb.VerifyAuthResponse, error) {
	return s.service.Verify(ctx, in)
}

func (s StubClient) Logout(ctx context.Context, in *authpb.LogoutAuthRequest, opts ...grpc.CallOption) (*authpb.LogoutAuthResponse, error) {
	return s.service.Logout(ctx, in)
}
