package follow

import (
	"context"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"google.golang.org/grpc"
)

type StubClient struct {
	service followpb.FollowServer
}

func (s StubClient) Follow(ctx context.Context, in *followpb.FollowRequest, opts ...grpc.CallOption) (*followpb.FollowResponse, error) {
	return s.service.Follow(ctx, in)
}

func (s StubClient) Unfollow(ctx context.Context, in *followpb.UnfollowRequest, opts ...grpc.CallOption) (*followpb.UnfollowResponse, error) {
	return s.service.Unfollow(ctx, in)
}

func (s StubClient) View(ctx context.Context, in *followpb.ViewRequest, opts ...grpc.CallOption) (*followpb.ViewResponse, error) {
	return s.service.View(ctx, in)
}
