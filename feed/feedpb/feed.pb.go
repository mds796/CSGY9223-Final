// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feed.proto

package feedpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Feed struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Posts                []*Post  `protobuf:"bytes,2,rep,name=Posts,proto3" json:"Posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{0}
}

func (m *Feed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed.Unmarshal(m, b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return xxx_messageInfo_Feed.Size(m)
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Feed) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Post struct {
	PostID               string     `protobuf:"bytes,1,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Text                 string     `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	User                 *User      `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty"`
	Timestamp            *Timestamp `protobuf:"bytes,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{2}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *Post) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Post) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Post) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Timestamp struct {
	EpochSeconds         int64    `protobuf:"varint,1,opt,name=EpochSeconds,proto3" json:"EpochSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{3}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetEpochSeconds() int64 {
	if m != nil {
		return m.EpochSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Feed)(nil), "feed.model.Feed")
	proto.RegisterType((*User)(nil), "feed.model.User")
	proto.RegisterType((*Post)(nil), "feed.model.Post")
	proto.RegisterType((*Timestamp)(nil), "feed.model.Timestamp")
}

func init() { proto.RegisterFile("feed.proto", fileDescriptor_d7a672c1337cb5ac) }

var fileDescriptor_d7a672c1337cb5ac = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4b, 0x4d, 0x4d,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0xb0, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x94, 0x42,
	0xb8, 0x58, 0xdc, 0x52, 0x53, 0x53, 0x84, 0x54, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04, 0xf4, 0x10, 0x4a, 0xf4, 0x40, 0xe2, 0x41, 0x60, 0x59, 0x21,
	0x35, 0x2e, 0xd6, 0x80, 0xfc, 0xe2, 0x92, 0x62, 0x09, 0x26, 0x05, 0x66, 0x74, 0x65, 0x20, 0x89,
	0x20, 0x88, 0xb4, 0x92, 0x16, 0xc4, 0x34, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0xb0, 0x99, 0x9c,
	0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x11,
	0x30, 0x5b, 0xa9, 0x97, 0x91, 0x8b, 0x05, 0xa4, 0x4b, 0x48, 0x8c, 0x8b, 0x0d, 0x44, 0xc3, 0x35,
	0x40, 0x79, 0x20, 0x4d, 0x21, 0xa9, 0x15, 0x25, 0x30, 0x4d, 0x20, 0x36, 0xdc, 0xb9, 0xcc, 0x78,
	0x9d, 0x6b, 0xcc, 0xc5, 0x19, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0x02,
	0x56, 0x2a, 0x8a, 0xac, 0x14, 0x2e, 0x19, 0x84, 0x50, 0xa7, 0xa4, 0x8f, 0xa4, 0x49, 0x48, 0x89,
	0x8b, 0xc7, 0xb5, 0x20, 0x3f, 0x39, 0x23, 0x38, 0x35, 0x39, 0x3f, 0x2f, 0xa5, 0x18, 0xec, 0x32,
	0xe6, 0x20, 0x14, 0x31, 0x27, 0x8e, 0x28, 0x36, 0x90, 0x99, 0x05, 0x49, 0x49, 0x6c, 0xe0, 0xf0,
	0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x54, 0x2a, 0x2d, 0x6d, 0x01, 0x00, 0x00,
}