// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import (
	basic "query_rpc/infrastructure/rpc/basic"
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

type User struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserPwd              string   `protobuf:"bytes,2,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserTag              []string `protobuf:"bytes,4,rep,name=user_tag,json=userTag,proto3" json:"user_tag,omitempty"`
	Pass                 bool     `protobuf:"varint,5,opt,name=pass,proto3" json:"pass,omitempty"`
	Woekspace            []int64  `protobuf:"varint,6,rep,packed,name=woekspace,proto3" json:"woekspace,omitempty"`
	History              []int64  `protobuf:"varint,7,rep,packed,name=history,proto3" json:"history,omitempty"`
	Subscribe            []int64  `protobuf:"varint,8,rep,packed,name=subscribe,proto3" json:"subscribe,omitempty"`
	Phone                int64    `protobuf:"varint,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Log                  string   `protobuf:"bytes,11,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
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

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetUserTag() []string {
	if m != nil {
		return m.UserTag
	}
	return nil
}

func (m *User) GetPass() bool {
	if m != nil {
		return m.Pass
	}
	return false
}

func (m *User) GetWoekspace() []int64 {
	if m != nil {
		return m.Woekspace
	}
	return nil
}

func (m *User) GetHistory() []int64 {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *User) GetSubscribe() []int64 {
	if m != nil {
		return m.Subscribe
	}
	return nil
}

func (m *User) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type UserLoginReq struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserPwd              string   `protobuf:"bytes,2,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginReq) Reset()         { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string { return proto.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()    {}
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *UserLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginReq.Unmarshal(m, b)
}
func (m *UserLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginReq.Marshal(b, m, deterministic)
}
func (m *UserLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginReq.Merge(m, src)
}
func (m *UserLoginReq) XXX_Size() int {
	return xxx_messageInfo_UserLoginReq.Size(m)
}
func (m *UserLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginReq proto.InternalMessageInfo

func (m *UserLoginReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserLoginReq) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type UserLoginResp struct {
	UserInfo             *User           `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Resp                 *basic.RespBody `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserLoginResp) Reset()         { *m = UserLoginResp{} }
func (m *UserLoginResp) String() string { return proto.CompactTextString(m) }
func (*UserLoginResp) ProtoMessage()    {}
func (*UserLoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *UserLoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResp.Unmarshal(m, b)
}
func (m *UserLoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResp.Marshal(b, m, deterministic)
}
func (m *UserLoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResp.Merge(m, src)
}
func (m *UserLoginResp) XXX_Size() int {
	return xxx_messageInfo_UserLoginResp.Size(m)
}
func (m *UserLoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResp proto.InternalMessageInfo

func (m *UserLoginResp) GetUserInfo() *User {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *UserLoginResp) GetResp() *basic.RespBody {
	if m != nil {
		return m.Resp
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserLoginReq)(nil), "user.UserLoginReq")
	proto.RegisterType((*UserLoginResp)(nil), "user.UserLoginResp")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0xbd, 0x69, 0xd2, 0x8f, 0x4c, 0xee, 0xa5, 0x65, 0xae, 0xe0, 0x5a, 0x7d, 0x08, 0xf1, 0xc1,
	0x3c, 0xa5, 0x50, 0xc1, 0x1f, 0x50, 0xf0, 0xa1, 0x20, 0x22, 0x51, 0x1f, 0x45, 0x36, 0xc9, 0x36,
	0x5d, 0x6c, 0xb3, 0x6b, 0xb6, 0xb5, 0xf4, 0xaf, 0xf8, 0x6b, 0x65, 0x27, 0xb6, 0x15, 0x7d, 0xf1,
	0x25, 0xcc, 0x9c, 0x33, 0xe7, 0x24, 0x73, 0x26, 0xd0, 0x5f, 0x1b, 0x51, 0x8f, 0xec, 0x23, 0xd1,
	0xb5, 0x5a, 0x29, 0xf4, 0x6c, 0x3d, 0x1c, 0x64, 0xdc, 0xc8, 0x7c, 0x54, 0x0b, 0xa3, 0x1b, 0x3c,
	0x7a, 0x6f, 0x81, 0xf7, 0x68, 0x44, 0x8d, 0xc7, 0xd0, 0xb5, 0x23, 0xcf, 0xb2, 0x60, 0x4e, 0xe8,
	0xc4, 0x6e, 0xda, 0xb1, 0xed, 0xb4, 0xc0, 0x13, 0xe8, 0x11, 0xa1, 0x37, 0x05, 0x6b, 0x85, 0x4e,
	0xec, 0xa7, 0x34, 0x78, 0xb7, 0x29, 0xf0, 0x14, 0x7c, 0xa2, 0x2a, 0xbe, 0x14, 0xcc, 0x25, 0x8e,
	0x66, 0x6f, 0xf9, 0x52, 0xec, 0x75, 0x2b, 0x5e, 0x32, 0x2f, 0x74, 0x77, 0xba, 0x07, 0x5e, 0x22,
	0x82, 0xa7, 0xb9, 0x31, 0xac, 0x1d, 0x3a, 0x71, 0x2f, 0xa5, 0x1a, 0xcf, 0xc0, 0xdf, 0x28, 0xf1,
	0x62, 0x34, 0xcf, 0x05, 0xeb, 0x84, 0x6e, 0xec, 0xa6, 0x07, 0x00, 0x19, 0x74, 0xe7, 0xd2, 0xac,
	0x54, 0xbd, 0x65, 0x5d, 0xe2, 0x76, 0xad, 0xd5, 0x99, 0x75, 0x66, 0xf2, 0x5a, 0x66, 0x82, 0xf5,
	0x1a, 0xdd, 0x1e, 0xc0, 0x23, 0x68, 0xeb, 0xb9, 0xaa, 0x04, 0xf3, 0x69, 0xa7, 0xa6, 0xb1, 0xa8,
	0x58, 0x72, 0xb9, 0x60, 0x40, 0xdf, 0xdc, 0x34, 0x38, 0x00, 0x77, 0xa1, 0x4a, 0x16, 0x10, 0x66,
	0xcb, 0x68, 0x02, 0x7f, 0x6d, 0x36, 0x37, 0xaa, 0x94, 0x55, 0x2a, 0x5e, 0xbf, 0x67, 0xd4, 0xfe,
	0x45, 0x46, 0xd1, 0x13, 0xfc, 0xfb, 0xe2, 0x61, 0x34, 0x5e, 0x7c, 0x86, 0x26, 0xab, 0x99, 0x22,
	0x9b, 0x60, 0x0c, 0x09, 0x5d, 0xca, 0xce, 0x35, 0x01, 0x4e, 0xab, 0x99, 0xc2, 0x73, 0xf0, 0xec,
	0xa1, 0xc8, 0x30, 0x18, 0xf7, 0x13, 0xba, 0x5d, 0x62, 0x3d, 0x26, 0xaa, 0xd8, 0xa6, 0x44, 0x8e,
	0xaf, 0x21, 0xb0, 0xb2, 0x7b, 0x51, 0xbf, 0xc9, 0x5c, 0xe0, 0x15, 0xf8, 0xfb, 0xb7, 0x21, 0x1e,
	0x6c, 0x77, 0x2b, 0x0c, 0xff, 0xff, 0xc0, 0x8c, 0x8e, 0xfe, 0x64, 0x1d, 0xfa, 0x1b, 0x2e, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x83, 0xb3, 0x09, 0x38, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, "/user.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) UserLogin(ctx context.Context, req *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}