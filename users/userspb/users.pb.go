// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users/userspb/users.proto

package userspb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EmptyData struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyData) Reset()         { *m = EmptyData{} }
func (m *EmptyData) String() string { return proto.CompactTextString(m) }
func (*EmptyData) ProtoMessage()    {}
func (*EmptyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{0}
}

func (m *EmptyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyData.Unmarshal(m, b)
}
func (m *EmptyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyData.Marshal(b, m, deterministic)
}
func (m *EmptyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyData.Merge(m, src)
}
func (m *EmptyData) XXX_Size() int {
	return xxx_messageInfo_EmptyData.Size(m)
}
func (m *EmptyData) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyData.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyData proto.InternalMessageInfo

type CreateUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserResponse) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type GetUsersResponse struct {
	Users                []*GetUsersResponse_User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{3}
}

func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(m, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetUsers() []*GetUsersResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUsersResponse_User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse_User) Reset()         { *m = GetUsersResponse_User{} }
func (m *GetUsersResponse_User) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse_User) ProtoMessage()    {}
func (*GetUsersResponse_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{3, 0}
}

func (m *GetUsersResponse_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse_User.Unmarshal(m, b)
}
func (m *GetUsersResponse_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse_User.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse_User.Merge(m, src)
}
func (m *GetUsersResponse_User) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse_User.Size(m)
}
func (m *GetUsersResponse_User) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse_User.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse_User proto.InternalMessageInfo

func (m *GetUsersResponse_User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUsersResponse_User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUsersResponse_User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type GetUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{4}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{5}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserResponse) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type AuthenticateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{6}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateResponse struct {
	Auth                 bool     `protobuf:"varint,1,opt,name=Auth,proto3" json:"Auth,omitempty"`
	Userid               int64    `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0d748b3820d20a, []int{7}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetAuth() bool {
	if m != nil {
		return m.Auth
	}
	return false
}

func (m *AuthenticateResponse) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func init() {
	proto.RegisterType((*EmptyData)(nil), "userspb.EmptyData")
	proto.RegisterType((*CreateUserRequest)(nil), "userspb.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "userspb.CreateUserResponse")
	proto.RegisterType((*GetUsersResponse)(nil), "userspb.GetUsersResponse")
	proto.RegisterType((*GetUsersResponse_User)(nil), "userspb.GetUsersResponse.User")
	proto.RegisterType((*GetUserRequest)(nil), "userspb.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "userspb.GetUserResponse")
	proto.RegisterType((*AuthenticateRequest)(nil), "userspb.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "userspb.AuthenticateResponse")
}

func init() { proto.RegisterFile("users/userspb/users.proto", fileDescriptor_9c0d748b3820d20a) }

var fileDescriptor_9c0d748b3820d20a = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0xc7, 0x1f, 0x5a, 0x1e, 0x5e, 0x06, 0x82, 0x3a, 0x1a, 0x2d, 0xf5, 0x25, 0x64, 0x4f, 0x9c,
	0x6a, 0x82, 0xde, 0x4c, 0x4c, 0x44, 0x0d, 0x07, 0xe5, 0x52, 0x63, 0x4c, 0xb8, 0x15, 0xd8, 0xc4,
	0x1e, 0xa0, 0xb5, 0xbb, 0xd5, 0xf8, 0x45, 0xfc, 0x02, 0x7e, 0x51, 0xbb, 0xdb, 0xdd, 0x15, 0x0a,
	0x1c, 0x4c, 0xb8, 0xb4, 0x3b, 0x33, 0xff, 0xfd, 0xed, 0xec, 0xcc, 0x2c, 0xb4, 0x53, 0x46, 0x13,
	0x76, 0x2e, 0xbf, 0xf1, 0x38, 0xff, 0x7b, 0x71, 0x12, 0xf1, 0x08, 0xab, 0xca, 0x49, 0x1a, 0x50,
	0xbf, 0x9f, 0xc5, 0xfc, 0xf3, 0x2e, 0xe0, 0x01, 0x79, 0x80, 0xbd, 0xdb, 0x84, 0x06, 0x9c, 0x3e,
	0x67, 0x51, 0x9f, 0xbe, 0xa5, 0x94, 0x71, 0x74, 0xa1, 0x26, 0xc4, 0xf3, 0x60, 0x46, 0x9d, 0x52,
	0xa7, 0xd4, 0xad, 0xfb, 0xc6, 0x16, 0xb1, 0x38, 0x60, 0xec, 0x23, 0x4a, 0xa6, 0x4e, 0x39, 0x8f,
	0x69, 0x9b, 0x8c, 0x00, 0x17, 0x61, 0x2c, 0x8e, 0xe6, 0x8c, 0x62, 0x0b, 0xac, 0x70, 0x2a, 0x39,
	0xb6, 0x9f, 0xad, 0x96, 0xe8, 0x56, 0x81, 0xee, 0x40, 0x75, 0x22, 0x09, 0x53, 0xc7, 0x96, 0x1b,
	0xb4, 0x49, 0xbe, 0x4a, 0xb0, 0x3b, 0xa0, 0x5c, 0x90, 0x99, 0x41, 0x5f, 0xc2, 0x7f, 0x79, 0xab,
	0x8c, 0x6e, 0x77, 0x1b, 0xbd, 0x33, 0x4f, 0xdd, 0xd1, 0x2b, 0x2a, 0x3d, 0x99, 0x51, 0x2e, 0x76,
	0x1f, 0xa1, 0x2c, 0xcc, 0x2d, 0x25, 0xd6, 0x81, 0x96, 0x3a, 0x4d, 0x97, 0xaf, 0xc0, 0x25, 0x2f,
	0xb0, 0x63, 0x14, 0x5b, 0xad, 0xc9, 0x10, 0xf6, 0x6f, 0x52, 0xfe, 0x4a, 0xe7, 0x3c, 0x9c, 0x64,
	0x8e, 0xbf, 0xb6, 0xcf, 0x2a, 0xb4, 0xaf, 0x0f, 0x07, 0xcb, 0x38, 0x95, 0x2c, 0x42, 0x59, 0xf8,
	0x25, 0xab, 0xe6, 0xcb, 0x35, 0x1e, 0x42, 0x45, 0x30, 0xc3, 0x9c, 0x62, 0xfb, 0xca, 0xea, 0x7d,
	0x5b, 0xd0, 0x94, 0x95, 0x7f, 0xa2, 0xc9, 0x7b, 0x38, 0xa1, 0x38, 0x00, 0xf8, 0x9d, 0x09, 0x74,
	0x4d, 0x87, 0x56, 0xa6, 0xce, 0x3d, 0x5e, 0x1b, 0xcb, 0x73, 0x20, 0xff, 0xf0, 0x0a, 0x6a, 0xba,
	0xab, 0x88, 0x46, 0x6a, 0x26, 0xd9, 0x6d, 0x6f, 0x6c, 0x7e, 0xb6, 0xf9, 0x1a, 0xaa, 0xca, 0x8b,
	0x47, 0x45, 0x9d, 0x3e, 0xdf, 0x59, 0x0d, 0x98, 0xfd, 0x43, 0x68, 0x2e, 0x96, 0x06, 0x4f, 0x8c,
	0x76, 0x4d, 0x03, 0xdc, 0xd3, 0x0d, 0x51, 0x8d, 0xeb, 0xd7, 0x47, 0xfa, 0x35, 0x8e, 0x2b, 0xf2,
	0x75, 0x5e, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe4, 0x85, 0x96, 0xba, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUsers(ctx context.Context, in *EmptyData, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/userspb.UsersService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUsers(ctx context.Context, in *EmptyData, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/userspb.UsersService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/userspb.UsersService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/userspb.UsersService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUsers(context.Context, *EmptyData) (*GetUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersServiceServer) GetUsers(ctx context.Context, req *EmptyData) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUsersServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersServiceServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.UsersService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.UsersService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUsers(ctx, req.(*EmptyData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.UsersService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.UsersService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userspb.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UsersService_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UsersService_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UsersService_GetUser_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UsersService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/userspb/users.proto",
}
