// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: rpc/credentials/logins.proto

package credentials

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Logins_CreateLogin_FullMethodName  = "/credentials.Logins/CreateLogin"
	Logins_GetLogin_FullMethodName     = "/credentials.Logins/GetLogin"
	Logins_GetLoginMany_FullMethodName = "/credentials.Logins/GetLoginMany"
	Logins_UpsertLogin_FullMethodName  = "/credentials.Logins/UpsertLogin"
	Logins_DeleteLogin_FullMethodName  = "/credentials.Logins/DeleteLogin"
)

// LoginsClient is the client API for Logins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginsClient interface {
	CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...grpc.CallOption) (*CreateLoginResponse, error)
	GetLogin(ctx context.Context, in *ReadLoginRequest, opts ...grpc.CallOption) (*ReadLoginResponse, error)
	GetLoginMany(ctx context.Context, in *ReadLoginRequest, opts ...grpc.CallOption) (*ReadLoginResponse, error)
	UpsertLogin(ctx context.Context, in *UpsertLoginRequest, opts ...grpc.CallOption) (*UpsertLoginResponse, error)
	DeleteLogin(ctx context.Context, in *DeleteLoginRequest, opts ...grpc.CallOption) (*DeleteLoginResponse, error)
}

type loginsClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginsClient(cc grpc.ClientConnInterface) LoginsClient {
	return &loginsClient{cc}
}

func (c *loginsClient) CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...grpc.CallOption) (*CreateLoginResponse, error) {
	out := new(CreateLoginResponse)
	err := c.cc.Invoke(ctx, Logins_CreateLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginsClient) GetLogin(ctx context.Context, in *ReadLoginRequest, opts ...grpc.CallOption) (*ReadLoginResponse, error) {
	out := new(ReadLoginResponse)
	err := c.cc.Invoke(ctx, Logins_GetLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginsClient) GetLoginMany(ctx context.Context, in *ReadLoginRequest, opts ...grpc.CallOption) (*ReadLoginResponse, error) {
	out := new(ReadLoginResponse)
	err := c.cc.Invoke(ctx, Logins_GetLoginMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginsClient) UpsertLogin(ctx context.Context, in *UpsertLoginRequest, opts ...grpc.CallOption) (*UpsertLoginResponse, error) {
	out := new(UpsertLoginResponse)
	err := c.cc.Invoke(ctx, Logins_UpsertLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginsClient) DeleteLogin(ctx context.Context, in *DeleteLoginRequest, opts ...grpc.CallOption) (*DeleteLoginResponse, error) {
	out := new(DeleteLoginResponse)
	err := c.cc.Invoke(ctx, Logins_DeleteLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginsServer is the server API for Logins service.
// All implementations must embed UnimplementedLoginsServer
// for forward compatibility
type LoginsServer interface {
	CreateLogin(context.Context, *CreateLoginRequest) (*CreateLoginResponse, error)
	GetLogin(context.Context, *ReadLoginRequest) (*ReadLoginResponse, error)
	GetLoginMany(context.Context, *ReadLoginRequest) (*ReadLoginResponse, error)
	UpsertLogin(context.Context, *UpsertLoginRequest) (*UpsertLoginResponse, error)
	DeleteLogin(context.Context, *DeleteLoginRequest) (*DeleteLoginResponse, error)
	mustEmbedUnimplementedLoginsServer()
}

// UnimplementedLoginsServer must be embedded to have forward compatible implementations.
type UnimplementedLoginsServer struct {
}

func (UnimplementedLoginsServer) CreateLogin(context.Context, *CreateLoginRequest) (*CreateLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogin not implemented")
}
func (UnimplementedLoginsServer) GetLogin(context.Context, *ReadLoginRequest) (*ReadLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogin not implemented")
}
func (UnimplementedLoginsServer) GetLoginMany(context.Context, *ReadLoginRequest) (*ReadLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginMany not implemented")
}
func (UnimplementedLoginsServer) UpsertLogin(context.Context, *UpsertLoginRequest) (*UpsertLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertLogin not implemented")
}
func (UnimplementedLoginsServer) DeleteLogin(context.Context, *DeleteLoginRequest) (*DeleteLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogin not implemented")
}
func (UnimplementedLoginsServer) mustEmbedUnimplementedLoginsServer() {}

// UnsafeLoginsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginsServer will
// result in compilation errors.
type UnsafeLoginsServer interface {
	mustEmbedUnimplementedLoginsServer()
}

func RegisterLoginsServer(s grpc.ServiceRegistrar, srv LoginsServer) {
	s.RegisterService(&Logins_ServiceDesc, srv)
}

func _Logins_CreateLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginsServer).CreateLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logins_CreateLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginsServer).CreateLogin(ctx, req.(*CreateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logins_GetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginsServer).GetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logins_GetLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginsServer).GetLogin(ctx, req.(*ReadLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logins_GetLoginMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginsServer).GetLoginMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logins_GetLoginMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginsServer).GetLoginMany(ctx, req.(*ReadLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logins_UpsertLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginsServer).UpsertLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logins_UpsertLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginsServer).UpsertLogin(ctx, req.(*UpsertLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logins_DeleteLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginsServer).DeleteLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logins_DeleteLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginsServer).DeleteLogin(ctx, req.(*DeleteLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logins_ServiceDesc is the grpc.ServiceDesc for Logins service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logins_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "credentials.Logins",
	HandlerType: (*LoginsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLogin",
			Handler:    _Logins_CreateLogin_Handler,
		},
		{
			MethodName: "GetLogin",
			Handler:    _Logins_GetLogin_Handler,
		},
		{
			MethodName: "GetLoginMany",
			Handler:    _Logins_GetLoginMany_Handler,
		},
		{
			MethodName: "UpsertLogin",
			Handler:    _Logins_UpsertLogin_Handler,
		},
		{
			MethodName: "DeleteLogin",
			Handler:    _Logins_DeleteLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/credentials/logins.proto",
}
