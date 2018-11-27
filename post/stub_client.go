package post

import (
	"context"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"google.golang.org/grpc"
)

type StubClient struct {
	service postpb.PostServer
}

func (s StubClient) Create(ctx context.Context, in *postpb.CreateRequest, opts ...grpc.CallOption) (*postpb.CreateResponse, error) {
	return s.service.Create(ctx, in)
}

func (s StubClient) View(ctx context.Context, in *postpb.ViewRequest, opts ...grpc.CallOption) (*postpb.ViewResponse, error) {
	return s.service.View(ctx, in)
}
