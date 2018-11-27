// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package postpb

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
	return fileDescriptor_e114ad14deab1dd1, []int{0}
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
	ID                   string     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
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
	return fileDescriptor_e114ad14deab1dd1, []int{1}
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

func (m *Post) GetID() string {
	if m != nil {
		return m.ID
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
	EpochNanoseconds     int64    `protobuf:"varint,1,opt,name=EpochNanoseconds,proto3" json:"EpochNanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{2}
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

func (m *Timestamp) GetEpochNanoseconds() int64 {
	if m != nil {
		return m.EpochNanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "post.model.User")
	proto.RegisterType((*Post)(nil), "post.model.Post")
	proto.RegisterType((*Timestamp)(nil), "post.model.Timestamp")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0x2f, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0xb0, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x94, 0xb4,
	0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x98, 0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x98,
	0xc0, 0x22, 0x60, 0xb6, 0x52, 0x2b, 0x23, 0x17, 0x4b, 0x40, 0x7e, 0x71, 0x09, 0x36, 0xc5, 0x21,
	0xa9, 0x15, 0x25, 0x30, 0xc5, 0x20, 0xb6, 0x90, 0x0a, 0xc4, 0x60, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x01, 0x3d, 0x84, 0x9d, 0x7a, 0x20, 0xf1, 0x20, 0x88, 0xb5, 0xc6, 0x5c, 0x9c, 0x21,
	0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x2c, 0x60, 0xa5, 0xa2, 0xc8, 0x4a, 0xe1,
	0x92, 0x41, 0x08, 0x75, 0x4a, 0xe6, 0x48, 0x9a, 0x84, 0xb4, 0xb8, 0x04, 0x5c, 0x0b, 0xf2, 0x93,
	0x33, 0xfc, 0x12, 0xf3, 0xf2, 0x8b, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x8a, 0xc1, 0x2e, 0x63, 0x0e,
	0xc2, 0x10, 0x77, 0xe2, 0x88, 0x62, 0x03, 0x99, 0x5d, 0x90, 0x94, 0xc4, 0x06, 0x0e, 0x09, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x6c, 0xda, 0x46, 0x17, 0x01, 0x00, 0x00,
}
