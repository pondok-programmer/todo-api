// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserIdentifier struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdentifier) Reset()         { *m = UserIdentifier{} }
func (m *UserIdentifier) String() string { return proto.CompactTextString(m) }
func (*UserIdentifier) ProtoMessage()    {}
func (*UserIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdentifier.Unmarshal(m, b)
}
func (m *UserIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdentifier.Marshal(b, m, deterministic)
}
func (m *UserIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdentifier.Merge(m, src)
}
func (m *UserIdentifier) XXX_Size() int {
	return xxx_messageInfo_UserIdentifier.Size(m)
}
func (m *UserIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdentifier proto.InternalMessageInfo

func (m *UserIdentifier) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserIdentifier)(nil), "pb.UserIdentifier")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe1, 0x62, 0x09, 0x2d,
	0x4e, 0x2d, 0x12, 0x12, 0xe2, 0x62, 0x29, 0x2d, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0xa4, 0xb8, 0x38, 0x40, 0xaa, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0,
	0xe2, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x6e, 0x62, 0x66, 0x8e, 0x04, 0x33, 0x58, 0x02,
	0xc2, 0x51, 0x52, 0xe1, 0xe2, 0x03, 0x99, 0xe6, 0x99, 0x92, 0x9a, 0x57, 0x92, 0x99, 0x96, 0x89,
	0xdd, 0xdc, 0x24, 0x36, 0xb0, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xcf, 0xac,
	0x3d, 0x8c, 0x00, 0x00, 0x00,
}
