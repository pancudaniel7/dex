// Code generated by protoc-gen-go.
// source: api/api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Client
	CreateClientReq
	CreateClientResp
	DeleteClientReq
	DeleteClientResp
	Password
	CreatePasswordReq
	CreatePasswordResp
	UpdatePasswordReq
	UpdatePasswordResp
	DeletePasswordReq
	DeletePasswordResp
	VersionReq
	VersionResp
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Client represents an OAuth2 client.
type Client struct {
	Id           string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Secret       string   `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	RedirectUris []string `protobuf:"bytes,3,rep,name=redirect_uris,json=redirectUris" json:"redirect_uris,omitempty"`
	TrustedPeers []string `protobuf:"bytes,4,rep,name=trusted_peers,json=trustedPeers" json:"trusted_peers,omitempty"`
	Public       bool     `protobuf:"varint,5,opt,name=public" json:"public,omitempty"`
	Name         string   `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	LogoUrl      string   `protobuf:"bytes,7,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CreateClientReq is a request to make a client.
type CreateClientReq struct {
	Client *Client `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientReq) Reset()                    { *m = CreateClientReq{} }
func (m *CreateClientReq) String() string            { return proto.CompactTextString(m) }
func (*CreateClientReq) ProtoMessage()               {}
func (*CreateClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateClientReq) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// CreateClientResp returns the response from creating a client.
type CreateClientResp struct {
	AlreadyExists bool    `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
	Client        *Client `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientResp) Reset()                    { *m = CreateClientResp{} }
func (m *CreateClientResp) String() string            { return proto.CompactTextString(m) }
func (*CreateClientResp) ProtoMessage()               {}
func (*CreateClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateClientResp) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// DeleteClientReq is a request to delete a client.
type DeleteClientReq struct {
	// The ID of the client.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteClientReq) Reset()                    { *m = DeleteClientReq{} }
func (m *DeleteClientReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientReq) ProtoMessage()               {}
func (*DeleteClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// DeleteClientResp determines if the.
type DeleteClientResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeleteClientResp) Reset()                    { *m = DeleteClientResp{} }
func (m *DeleteClientResp) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientResp) ProtoMessage()               {}
func (*DeleteClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Password is an email for password mapping managed by the storage.
type Password struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// Currently we do not accept plain text passwords. Could be an option in the future.
	Hash     []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	UserId   string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Password) Reset()                    { *m = Password{} }
func (m *Password) String() string            { return proto.CompactTextString(m) }
func (*Password) ProtoMessage()               {}
func (*Password) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// CreatePasswordReq is a request to make a password.
type CreatePasswordReq struct {
	Password *Password `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *CreatePasswordReq) Reset()                    { *m = CreatePasswordReq{} }
func (m *CreatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordReq) ProtoMessage()               {}
func (*CreatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreatePasswordReq) GetPassword() *Password {
	if m != nil {
		return m.Password
	}
	return nil
}

// CreatePasswordResp returns the response from creating a password.
type CreatePasswordResp struct {
	AlreadyExists bool `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
}

func (m *CreatePasswordResp) Reset()                    { *m = CreatePasswordResp{} }
func (m *CreatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordResp) ProtoMessage()               {}
func (*CreatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// UpdatePasswordReq is a request to modify an existing password.
type UpdatePasswordReq struct {
	// The email used to lookup the password. This field cannot be modified
	Email       string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	NewHash     []byte `protobuf:"bytes,2,opt,name=new_hash,json=newHash,proto3" json:"new_hash,omitempty"`
	NewUsername string `protobuf:"bytes,3,opt,name=new_username,json=newUsername" json:"new_username,omitempty"`
}

func (m *UpdatePasswordReq) Reset()                    { *m = UpdatePasswordReq{} }
func (m *UpdatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordReq) ProtoMessage()               {}
func (*UpdatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// UpdatePasswordResp returns the response from modifying an existing password.
type UpdatePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *UpdatePasswordResp) Reset()                    { *m = UpdatePasswordResp{} }
func (m *UpdatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResp) ProtoMessage()               {}
func (*UpdatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// DeletePasswordReq is a request to delete a password.
type DeletePasswordReq struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *DeletePasswordReq) Reset()                    { *m = DeletePasswordReq{} }
func (m *DeletePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordReq) ProtoMessage()               {}
func (*DeletePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// DeletePasswordResp returns the response from deleting a password.
type DeletePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeletePasswordResp) Reset()                    { *m = DeletePasswordResp{} }
func (m *DeletePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordResp) ProtoMessage()               {}
func (*DeletePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// VersionReq is a request to fetch version info.
type VersionReq struct {
}

func (m *VersionReq) Reset()                    { *m = VersionReq{} }
func (m *VersionReq) String() string            { return proto.CompactTextString(m) }
func (*VersionReq) ProtoMessage()               {}
func (*VersionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// VersionResp holds the version info of components.
type VersionResp struct {
	// Semantic version of the server.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// Numeric version of the API. It increases everytime a new call is added to the API.
	// Clients should use this info to determine if the server supports specific features.
	Api int32 `protobuf:"varint,2,opt,name=api" json:"api,omitempty"`
}

func (m *VersionResp) Reset()                    { *m = VersionResp{} }
func (m *VersionResp) String() string            { return proto.CompactTextString(m) }
func (*VersionResp) ProtoMessage()               {}
func (*VersionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*CreateClientReq)(nil), "api.CreateClientReq")
	proto.RegisterType((*CreateClientResp)(nil), "api.CreateClientResp")
	proto.RegisterType((*DeleteClientReq)(nil), "api.DeleteClientReq")
	proto.RegisterType((*DeleteClientResp)(nil), "api.DeleteClientResp")
	proto.RegisterType((*Password)(nil), "api.Password")
	proto.RegisterType((*CreatePasswordReq)(nil), "api.CreatePasswordReq")
	proto.RegisterType((*CreatePasswordResp)(nil), "api.CreatePasswordResp")
	proto.RegisterType((*UpdatePasswordReq)(nil), "api.UpdatePasswordReq")
	proto.RegisterType((*UpdatePasswordResp)(nil), "api.UpdatePasswordResp")
	proto.RegisterType((*DeletePasswordReq)(nil), "api.DeletePasswordReq")
	proto.RegisterType((*DeletePasswordResp)(nil), "api.DeletePasswordResp")
	proto.RegisterType((*VersionReq)(nil), "api.VersionReq")
	proto.RegisterType((*VersionResp)(nil), "api.VersionResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Dex service

type DexClient interface {
	// CreateClient creates a client.
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
}

type dexClient struct {
	cc *grpc.ClientConn
}

func NewDexClient(cc *grpc.ClientConn) DexClient {
	return &dexClient{cc}
}

func (c *dexClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error) {
	out := new(CreateClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	out := new(DeleteClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeleteClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error) {
	out := new(CreatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/UpdatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error) {
	out := new(DeletePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeletePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	out := new(VersionResp)
	err := grpc.Invoke(ctx, "/api.Dex/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dex service

type DexServer interface {
	// CreateClient creates a client.
	CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(context.Context, *VersionReq) (*VersionResp, error)
}

func RegisterDexServer(s *grpc.Server, srv DexServer) {
	s.RegisterService(&_Dex_serviceDesc, srv)
}

func _Dex_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreatePassword(ctx, req.(*CreatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeletePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeletePassword(ctx, req.(*DeletePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).GetVersion(ctx, req.(*VersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Dex",
	HandlerType: (*DexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _Dex_CreateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Dex_DeleteClient_Handler,
		},
		{
			MethodName: "CreatePassword",
			Handler:    _Dex_CreatePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Dex_UpdatePassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _Dex_DeletePassword_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Dex_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0x8d, 0x2d, 0x3f, 0xe4, 0xf1, 0x9b, 0xc8, 0x4d, 0x14, 0xdf, 0x8d, 0xc3, 0xa0, 0x80, 0xb3,
	0x49, 0x90, 0x14, 0x68, 0x17, 0x45, 0xbb, 0x71, 0xfa, 0xda, 0x05, 0x02, 0xdc, 0x65, 0x05, 0xc5,
	0x9a, 0x26, 0x04, 0x14, 0x89, 0x21, 0xa9, 0x3a, 0xfd, 0x80, 0x7e, 0x58, 0xff, 0xac, 0x20, 0x45,
	0xbb, 0x92, 0xec, 0xc2, 0xdd, 0xf1, 0x1c, 0xce, 0x9c, 0xd1, 0xcc, 0x19, 0x0a, 0xfa, 0x21, 0x67,
	0x97, 0x21, 0x67, 0x17, 0x5c, 0xa4, 0x2a, 0x25, 0x4e, 0xc8, 0x19, 0xfd, 0x55, 0x83, 0xd6, 0x3c,
	0x66, 0x98, 0x28, 0x32, 0x80, 0x3a, 0x8b, 0xbc, 0xda, 0xb4, 0x36, 0xeb, 0xf8, 0x75, 0x16, 0x91,
	0x23, 0x68, 0x49, 0x5c, 0x0a, 0x54, 0x5e, 0xdd, 0x70, 0x16, 0x91, 0x33, 0xe8, 0x0b, 0x8c, 0x98,
	0xc0, 0xa5, 0x0a, 0x32, 0xc1, 0xa4, 0xe7, 0x4c, 0x9d, 0x59, 0xc7, 0xef, 0xad, 0xc9, 0x85, 0x60,
	0x52, 0x07, 0x29, 0x91, 0x49, 0x85, 0x51, 0xc0, 0x11, 0x85, 0xf4, 0x1a, 0x79, 0x90, 0x25, 0x6f,
	0x35, 0xa7, 0x2b, 0xf0, 0xec, 0x2e, 0x66, 0x4b, 0xaf, 0x39, 0xad, 0xcd, 0x5c, 0xdf, 0x22, 0x42,
	0xa0, 0x91, 0x84, 0x8f, 0xe8, 0xb5, 0x4c, 0x5d, 0x73, 0x26, 0x27, 0xe0, 0xc6, 0xe9, 0x7d, 0x1a,
	0x64, 0x22, 0xf6, 0xda, 0x86, 0x6f, 0x6b, 0xbc, 0x10, 0x31, 0x7d, 0x05, 0xc3, 0xb9, 0xc0, 0x50,
	0x61, 0xde, 0x88, 0x8f, 0x4f, 0xe4, 0x0c, 0x5a, 0x4b, 0x03, 0x4c, 0x3f, 0xdd, 0xeb, 0xee, 0x85,
	0xee, 0xdb, 0xde, 0xdb, 0x2b, 0xfa, 0x15, 0x46, 0xe5, 0x3c, 0xc9, 0xc9, 0x0b, 0x18, 0x84, 0xb1,
	0xc0, 0x30, 0xfa, 0x11, 0xe0, 0x33, 0x93, 0x4a, 0x1a, 0x01, 0xd7, 0xef, 0x5b, 0xf6, 0xbd, 0x21,
	0x0b, 0xfa, 0xf5, 0xbf, 0xeb, 0x9f, 0xc2, 0xf0, 0x06, 0x63, 0x2c, 0x7e, 0x57, 0x65, 0xc6, 0xf4,
	0x12, 0x46, 0xe5, 0x10, 0xc9, 0xc9, 0xff, 0xd0, 0x49, 0x52, 0x15, 0x7c, 0x4b, 0xb3, 0x24, 0xb2,
	0xd5, 0xdd, 0x24, 0x55, 0x1f, 0x34, 0xa6, 0x0c, 0xdc, 0xdb, 0x50, 0xca, 0x55, 0x2a, 0x22, 0x72,
	0x08, 0x4d, 0x7c, 0x0c, 0x59, 0x6c, 0xf5, 0x72, 0xa0, 0x87, 0xf7, 0x10, 0xca, 0x07, 0xf3, 0x61,
	0x3d, 0xdf, 0x9c, 0xc9, 0x04, 0xdc, 0x4c, 0xa2, 0x30, 0x43, 0x75, 0x4c, 0xf0, 0x06, 0x93, 0x63,
	0x68, 0xeb, 0x73, 0xc0, 0x22, 0xaf, 0x91, 0xfb, 0xac, 0xe1, 0xe7, 0x88, 0xbe, 0x83, 0x71, 0x3e,
	0x9e, 0x75, 0x41, 0xdd, 0xc0, 0x39, 0xb8, 0xdc, 0x42, 0x3b, 0xda, 0xbe, 0x69, 0x7d, 0x13, 0xb3,
	0xb9, 0xa6, 0x6f, 0x80, 0x54, 0xf3, 0xff, 0x79, 0xc0, 0xf4, 0x1e, 0xc6, 0x0b, 0x1e, 0x55, 0x8a,
	0xef, 0x6e, 0xf8, 0x04, 0xdc, 0x04, 0x57, 0x41, 0xa1, 0xe9, 0x76, 0x82, 0xab, 0x4f, 0xba, 0xef,
	0x53, 0xe8, 0xe9, 0xab, 0x4a, 0xef, 0xdd, 0x04, 0x57, 0x0b, 0x4b, 0xd1, 0x2b, 0x20, 0xd5, 0x42,
	0xfb, 0x3c, 0x38, 0x87, 0x71, 0x6e, 0xda, 0xde, 0x6f, 0xd3, 0xea, 0xd5, 0xd0, 0x7d, 0xea, 0x3d,
	0x80, 0x2f, 0x28, 0x24, 0x4b, 0x13, 0x1f, 0x9f, 0xe8, 0x6b, 0xe8, 0x6e, 0x90, 0xe4, 0xf9, 0x9b,
	0x14, 0xdf, 0x51, 0xd8, 0x32, 0x16, 0x91, 0x11, 0xe8, 0xd7, 0x6c, 0xda, 0x6f, 0xfa, 0xfa, 0x78,
	0xfd, 0xd3, 0x01, 0xe7, 0x06, 0x9f, 0xc9, 0x5b, 0xe8, 0x15, 0x97, 0x9c, 0x1c, 0xe6, 0x9b, 0x5a,
	0x7e, 0x2f, 0x93, 0xff, 0x76, 0xb0, 0x92, 0xd3, 0x03, 0x9d, 0x5e, 0x5c, 0x50, 0x9b, 0x5e, 0x59,
	0x6b, 0x9b, 0x5e, 0xdd, 0x64, 0x7a, 0x40, 0xe6, 0x30, 0x28, 0xef, 0x00, 0x39, 0x2a, 0x54, 0x2a,
	0xcc, 0x6f, 0x72, 0xbc, 0x93, 0x5f, 0x8b, 0x94, 0x2d, 0xb2, 0x22, 0x5b, 0x0b, 0x62, 0x45, 0xb6,
	0xfd, 0xcc, 0x45, 0xca, 0x4e, 0x58, 0x91, 0x2d, 0x27, 0xad, 0xc8, 0xb6, 0x6d, 0xf4, 0x80, 0x5c,
	0x01, 0x7c, 0x44, 0x65, 0x0d, 0x21, 0x43, 0x13, 0xf8, 0xc7, 0xac, 0xc9, 0xa8, 0x4c, 0xe8, 0x94,
	0xbb, 0x96, 0xf9, 0xd9, 0xbe, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x72, 0xb9, 0x88, 0x7f, 0x7d,
	0x05, 0x00, 0x00,
}
