// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Todo struct {
	Id                   *TodoIdentifier      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() *TodoIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type TodoCreate struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoCreate) Reset()         { *m = TodoCreate{} }
func (m *TodoCreate) String() string { return proto.CompactTextString(m) }
func (*TodoCreate) ProtoMessage()    {}
func (*TodoCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *TodoCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoCreate.Unmarshal(m, b)
}
func (m *TodoCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoCreate.Marshal(b, m, deterministic)
}
func (m *TodoCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoCreate.Merge(m, src)
}
func (m *TodoCreate) XXX_Size() int {
	return xxx_messageInfo_TodoCreate.Size(m)
}
func (m *TodoCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoCreate.DiscardUnknown(m)
}

var xxx_messageInfo_TodoCreate proto.InternalMessageInfo

func (m *TodoCreate) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TodoCreate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type TodoList struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoList) Reset()         { *m = TodoList{} }
func (m *TodoList) String() string { return proto.CompactTextString(m) }
func (*TodoList) ProtoMessage()    {}
func (*TodoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *TodoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoList.Unmarshal(m, b)
}
func (m *TodoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoList.Marshal(b, m, deterministic)
}
func (m *TodoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoList.Merge(m, src)
}
func (m *TodoList) XXX_Size() int {
	return xxx_messageInfo_TodoList.Size(m)
}
func (m *TodoList) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoList.DiscardUnknown(m)
}

var xxx_messageInfo_TodoList proto.InternalMessageInfo

func (m *TodoList) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type TodoIdentifier struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoIdentifier) Reset()         { *m = TodoIdentifier{} }
func (m *TodoIdentifier) String() string { return proto.CompactTextString(m) }
func (*TodoIdentifier) ProtoMessage()    {}
func (*TodoIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *TodoIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoIdentifier.Unmarshal(m, b)
}
func (m *TodoIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoIdentifier.Marshal(b, m, deterministic)
}
func (m *TodoIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoIdentifier.Merge(m, src)
}
func (m *TodoIdentifier) XXX_Size() int {
	return xxx_messageInfo_TodoIdentifier.Size(m)
}
func (m *TodoIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_TodoIdentifier proto.InternalMessageInfo

func (m *TodoIdentifier) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Todo)(nil), "pb.Todo")
	proto.RegisterType((*TodoCreate)(nil), "pb.TodoCreate")
	proto.RegisterType((*TodoList)(nil), "pb.TodoList")
	proto.RegisterType((*TodoIdentifier)(nil), "pb.TodoIdentifier")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x86, 0x4d, 0x8c, 0x41, 0x27, 0xe2, 0x61, 0x29, 0x25, 0xa4, 0x60, 0xc3, 0x62, 0x41, 0x3c,
	0x44, 0xb0, 0xa7, 0x1e, 0x4b, 0x2d, 0xa5, 0xd0, 0x53, 0xb4, 0xe7, 0x92, 0xb8, 0xa3, 0x2c, 0xa8,
	0x1b, 0x92, 0x49, 0xa1, 0x0f, 0xd3, 0xd7, 0xea, 0xf3, 0x94, 0xdd, 0x24, 0xb6, 0x56, 0xe9, 0x2d,
	0xf3, 0xcf, 0xcf, 0xcc, 0xf7, 0x4f, 0x16, 0x80, 0x94, 0x50, 0x51, 0x96, 0x2b, 0x52, 0xcc, 0xce,
	0xd2, 0xe0, 0x6a, 0xa3, 0xd4, 0x66, 0x8b, 0x53, 0xa3, 0xa4, 0xe5, 0x7a, 0x8a, 0xbb, 0x8c, 0x3e,
	0x2a, 0x43, 0x70, 0xfd, 0xb7, 0x49, 0x72, 0x87, 0x05, 0x25, 0xbb, 0xac, 0x36, 0xf4, 0x0b, 0x4a,
	0xa8, 0x2c, 0xaa, 0x8a, 0x7f, 0x5a, 0xe0, 0x2c, 0x95, 0x50, 0x8c, 0x83, 0x2d, 0x85, 0x6f, 0x85,
	0xd6, 0xd8, 0x9b, 0xb1, 0x28, 0x4b, 0x23, 0xad, 0x3e, 0x0b, 0xdc, 0x93, 0x5c, 0x4b, 0xcc, 0x63,
	0x5b, 0x0a, 0x76, 0x01, 0x1d, 0x92, 0xb4, 0x45, 0xdf, 0x0e, 0xad, 0x71, 0x2f, 0xae, 0x0a, 0x16,
	0x82, 0x27, 0xb0, 0x58, 0xe5, 0x32, 0x23, 0xa9, 0xf6, 0x7e, 0xdb, 0xf4, 0x7e, 0x4b, 0xec, 0x0e,
	0x60, 0x95, 0x63, 0x42, 0x28, 0xde, 0x12, 0xf2, 0x1d, 0xb3, 0x23, 0x88, 0x2a, 0xd0, 0xa8, 0x01,
	0x8d, 0x96, 0x0d, 0x68, 0xdc, 0xab, 0xdd, 0xf7, 0xc4, 0xe7, 0x00, 0x1a, 0xe4, 0xc1, 0x08, 0x3f,
	0x00, 0xd6, 0x3f, 0x00, 0xf6, 0x09, 0x00, 0x9f, 0x40, 0x57, 0x4f, 0x79, 0x91, 0x05, 0xb1, 0x21,
	0x74, 0xf4, 0x3d, 0x0b, 0xdf, 0x0a, 0xdb, 0x63, 0x6f, 0xd6, 0x6d, 0xb2, 0xc6, 0x95, 0xcc, 0x47,
	0x30, 0x38, 0x8e, 0xce, 0x18, 0x38, 0x65, 0x59, 0x1f, 0xa7, 0x17, 0x9b, 0xef, 0xd9, 0x97, 0x05,
	0x9e, 0xb6, 0x2d, 0x30, 0x7f, 0x97, 0x2b, 0x64, 0x11, 0x38, 0x66, 0xfa, 0xe5, 0x49, 0xac, 0x47,
	0xfd, 0x73, 0x82, 0x7e, 0xb3, 0x46, 0xbb, 0x78, 0x8b, 0xdd, 0x40, 0xfb, 0x09, 0x89, 0x9d, 0xb9,
	0x74, 0x70, 0x20, 0xe2, 0x2d, 0x36, 0x02, 0xb7, 0x8e, 0x3e, 0x68, 0xd4, 0xaa, 0x3e, 0x72, 0x0d,
	0xc1, 0x7d, 0xcd, 0x84, 0x76, 0x1d, 0xd4, 0xa3, 0xfe, 0x04, 0xdc, 0x39, 0x6e, 0x91, 0xf0, 0xec,
	0x3e, 0xd0, 0xda, 0xc2, 0x3c, 0x0a, 0xde, 0x4a, 0x5d, 0x03, 0x7e, 0xfb, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x16, 0xeb, 0x1a, 0x75, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error)
	Get(ctx context.Context, in *TodoIdentifier, opts ...grpc.CallOption) (*Todo, error)
	Create(ctx context.Context, in *TodoCreate, opts ...grpc.CallOption) (*Todo, error)
	// id, created_at ignored
	Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	Delete(ctx context.Context, in *TodoIdentifier, opts ...grpc.CallOption) (*Status, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, "/pb.TodoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Get(ctx context.Context, in *TodoIdentifier, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/pb.TodoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Create(ctx context.Context, in *TodoCreate, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/pb.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/pb.TodoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *TodoIdentifier, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pb.TodoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	List(context.Context, *empty.Empty) (*TodoList, error)
	Get(context.Context, *TodoIdentifier) (*Todo, error)
	Create(context.Context, *TodoCreate) (*Todo, error)
	// id, created_at ignored
	Update(context.Context, *Todo) (*Todo, error)
	Delete(context.Context, *TodoIdentifier) (*Status, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) List(ctx context.Context, req *empty.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTodoServiceServer) Get(ctx context.Context, req *TodoIdentifier) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTodoServiceServer) Create(ctx context.Context, req *TodoCreate) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTodoServiceServer) Update(ctx context.Context, req *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTodoServiceServer) Delete(ctx context.Context, req *TodoIdentifier) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Get(ctx, req.(*TodoIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*TodoCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*TodoIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TodoService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TodoService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
