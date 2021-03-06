// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package postpb

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

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Post                 *Post    `protobuf:"bytes,2,opt,name=Post,proto3" json:"Post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreateResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ViewRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewRequest) Reset()         { *m = ViewRequest{} }
func (m *ViewRequest) String() string { return proto.CompactTextString(m) }
func (*ViewRequest) ProtoMessage()    {}
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
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

func (m *ViewRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ViewResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewResponse) Reset()         { *m = ViewResponse{} }
func (m *ViewResponse) String() string { return proto.CompactTextString(m) }
func (*ViewResponse) ProtoMessage()    {}
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
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

func (m *ViewResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ListRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "post.service.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "post.service.CreateResponse")
	proto.RegisterType((*ViewRequest)(nil), "post.service.ViewRequest")
	proto.RegisterType((*ViewResponse)(nil), "post.service.ViewResponse")
	proto.RegisterType((*ListRequest)(nil), "post.service.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "post.service.ListResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0xc8, 0x2f, 0x2e, 0xd1,
	0x83, 0x8a, 0x49, 0x71, 0x81, 0x79, 0x60, 0x19, 0xa5, 0x68, 0x2e, 0x5e, 0xe7, 0xa2, 0xd4, 0xc4,
	0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x15, 0x2e, 0x96, 0xd0, 0xe2, 0xd4,
	0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x01, 0x3d, 0xb0, 0xda, 0xdc, 0xfc, 0x94, 0xd4,
	0x1c, 0x3d, 0x90, 0x78, 0x10, 0x58, 0x16, 0xa4, 0x2a, 0x20, 0xbf, 0xb8, 0x44, 0x82, 0x09, 0x53,
	0x15, 0x48, 0x3c, 0x08, 0x2c, 0xab, 0x64, 0xc6, 0xc5, 0x07, 0x33, 0xbc, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x15, 0xae, 0x8f, 0x11, 0xaf, 0x3e, 0x63, 0x2e, 0xee, 0xb0, 0xcc, 0xd4, 0x72, 0x24, 0x27,
	0x11, 0xa1, 0xc9, 0x84, 0x8b, 0x07, 0xa2, 0x89, 0x54, 0xab, 0x7c, 0x32, 0x8b, 0x4b, 0x48, 0xf2,
	0xbd, 0x92, 0x19, 0x17, 0x0f, 0x44, 0x13, 0xd4, 0x2a, 0x35, 0x2e, 0x56, 0x90, 0x61, 0xc5, 0x12,
	0x8c, 0x0a, 0xcc, 0x58, 0xed, 0x82, 0x48, 0x1b, 0x9d, 0x64, 0x84, 0xb8, 0x49, 0xc8, 0x99, 0x8b,
	0x0d, 0x12, 0x30, 0x42, 0xd2, 0x7a, 0xc8, 0x51, 0xa3, 0x87, 0x12, 0x17, 0x52, 0x32, 0xd8, 0x25,
	0xa1, 0xb6, 0xda, 0x72, 0xb1, 0x80, 0x3c, 0x2c, 0x24, 0x89, 0xaa, 0x0a, 0x29, 0xe4, 0xa4, 0xa4,
	0xb0, 0x49, 0x21, 0xb4, 0x83, 0x3c, 0x81, 0xae, 0x1d, 0x29, 0x34, 0xd0, 0xb5, 0x23, 0xfb, 0xd9,
	0x89, 0x23, 0x8a, 0x0d, 0x24, 0x59, 0x90, 0x94, 0xc4, 0x06, 0x4e, 0x49, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x58, 0xbb, 0x2c, 0x74, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type postClient struct {
	cc *grpc.ClientConn
}

func NewPostClient(cc *grpc.ClientConn) PostClient {
	return &postClient{cc}
}

func (c *postClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/post.service.Post/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, "/post.service.Post/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/post.service.Post/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
type PostServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	View(context.Context, *ViewRequest) (*ViewResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterPostServer(s *grpc.Server, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

func _Post_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.service.Post/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.service.Post/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).View(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.service.Post/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.service.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Post_Create_Handler,
		},
		{
			MethodName: "View",
			Handler:    _Post_View_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Post_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
