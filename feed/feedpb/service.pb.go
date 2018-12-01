// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package feedpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ViewRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewRequest) Reset()         { *m = ViewRequest{} }
func (m *ViewRequest) String() string { return proto.CompactTextString(m) }
func (*ViewRequest) ProtoMessage()    {}
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewRequest.Unmarshal(m, b)
}
func (m *ViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewRequest.Marshal(b, m, deterministic)
}
func (m *ViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewRequest.Merge(m, src)
}
func (m *ViewRequest) XXX_Size() int {
	return xxx_messageInfo_ViewRequest.Size(m)
}
func (m *ViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewRequest proto.InternalMessageInfo

func (m *ViewRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ViewResponse struct {
	Feed                 *Feed    `protobuf:"bytes,1,opt,name=Feed,proto3" json:"Feed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewResponse) Reset()         { *m = ViewResponse{} }
func (m *ViewResponse) String() string { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()    {}
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ViewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewResponse.Unmarshal(m, b)
}
func (m *ViewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewResponse.Marshal(b, m, deterministic)
}
func (m *ViewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewResponse.Merge(m, src)
}
func (m *ViewResponse) XXX_Size() int {
	return xxx_messageInfo_ViewResponse.Size(m)
}
func (m *ViewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ViewResponse proto.InternalMessageInfo

func (m *ViewResponse) GetFeed() *Feed {
	if m != nil {
		return m.Feed
	}
	return nil
}

func init() {
	proto.RegisterType((*ViewRequest)(nil), "feed.service.ViewRequest")
	proto.RegisterType((*ViewResponse)(nil), "feed.service.ViewResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x4b, 0x4d, 0x4d, 0xd1,
	0x83, 0x8a, 0x49, 0x71, 0x81, 0x79, 0x60, 0x19, 0x25, 0x63, 0x2e, 0xee, 0xb0, 0xcc, 0xd4, 0xf2,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x15, 0x2e, 0x96, 0xd0, 0xe2, 0xd4, 0x22, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x01, 0x3d, 0xb0, 0xca, 0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x3d,
	0x90, 0x78, 0x10, 0x58, 0x56, 0xc9, 0x84, 0x8b, 0x07, 0xa2, 0xa9, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x15, 0xa4, 0xcb, 0x2d, 0x35, 0x35, 0x05, 0x9b, 0x2e, 0x90, 0x78, 0x10, 0x58, 0xd6, 0xc8, 0x15,
	0xa2, 0x4a, 0xc8, 0x96, 0x8b, 0x05, 0xa4, 0x5b, 0x48, 0x52, 0x0f, 0xd9, 0x55, 0x7a, 0x48, 0xce,
	0x90, 0x92, 0xc2, 0x26, 0x05, 0xb1, 0xcc, 0x89, 0x23, 0x8a, 0x0d, 0x24, 0x59, 0x90, 0x94, 0xc4,
	0x06, 0xf6, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xee, 0x0e, 0x75, 0xe4, 0xed, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedClient interface {
	View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
}

type feedClient struct {
	cc *grpc.ClientConn
}

func NewFeedClient(cc *grpc.ClientConn) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, "/feed.service.Feed/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
type FeedServer interface {
	View(context.Context, *ViewRequest) (*ViewResponse, error)
}

func RegisterFeedServer(s *grpc.Server, srv FeedServer) {
	s.RegisterService(&_Feed_serviceDesc, srv)
}

func _Feed_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.service.Feed/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).View(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Feed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feed.service.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "View",
			Handler:    _Feed_View_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
