// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: rpc/hosts/hosts.proto

package hosts

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
	Hosts_CreateHost_FullMethodName  = "/hosts.Hosts/CreateHost"
	Hosts_GetHost_FullMethodName     = "/hosts.Hosts/GetHost"
	Hosts_GetHostMany_FullMethodName = "/hosts.Hosts/GetHostMany"
	Hosts_UpsertHost_FullMethodName  = "/hosts.Hosts/UpsertHost"
	Hosts_DeleteHost_FullMethodName  = "/hosts.Hosts/DeleteHost"
)

// HostsClient is the client API for Hosts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostsClient interface {
	CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostResponse, error)
	GetHost(ctx context.Context, in *ReadHostRequest, opts ...grpc.CallOption) (*ReadHostResponse, error)
	GetHostMany(ctx context.Context, in *ReadHostRequest, opts ...grpc.CallOption) (*ReadHostResponse, error)
	UpsertHost(ctx context.Context, in *UpsertHostRequest, opts ...grpc.CallOption) (*UpsertHostResponse, error)
	DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error)
}

type hostsClient struct {
	cc grpc.ClientConnInterface
}

func NewHostsClient(cc grpc.ClientConnInterface) HostsClient {
	return &hostsClient{cc}
}

func (c *hostsClient) CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostResponse, error) {
	out := new(CreateHostResponse)
	err := c.cc.Invoke(ctx, Hosts_CreateHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsClient) GetHost(ctx context.Context, in *ReadHostRequest, opts ...grpc.CallOption) (*ReadHostResponse, error) {
	out := new(ReadHostResponse)
	err := c.cc.Invoke(ctx, Hosts_GetHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsClient) GetHostMany(ctx context.Context, in *ReadHostRequest, opts ...grpc.CallOption) (*ReadHostResponse, error) {
	out := new(ReadHostResponse)
	err := c.cc.Invoke(ctx, Hosts_GetHostMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsClient) UpsertHost(ctx context.Context, in *UpsertHostRequest, opts ...grpc.CallOption) (*UpsertHostResponse, error) {
	out := new(UpsertHostResponse)
	err := c.cc.Invoke(ctx, Hosts_UpsertHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsClient) DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostResponse, error) {
	out := new(DeleteHostResponse)
	err := c.cc.Invoke(ctx, Hosts_DeleteHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostsServer is the server API for Hosts service.
// All implementations must embed UnimplementedHostsServer
// for forward compatibility
type HostsServer interface {
	CreateHost(context.Context, *CreateHostRequest) (*CreateHostResponse, error)
	GetHost(context.Context, *ReadHostRequest) (*ReadHostResponse, error)
	GetHostMany(context.Context, *ReadHostRequest) (*ReadHostResponse, error)
	UpsertHost(context.Context, *UpsertHostRequest) (*UpsertHostResponse, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*DeleteHostResponse, error)
	mustEmbedUnimplementedHostsServer()
}

// UnimplementedHostsServer must be embedded to have forward compatible implementations.
type UnimplementedHostsServer struct {
}

func (UnimplementedHostsServer) CreateHost(context.Context, *CreateHostRequest) (*CreateHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHost not implemented")
}
func (UnimplementedHostsServer) GetHost(context.Context, *ReadHostRequest) (*ReadHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHost not implemented")
}
func (UnimplementedHostsServer) GetHostMany(context.Context, *ReadHostRequest) (*ReadHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostMany not implemented")
}
func (UnimplementedHostsServer) UpsertHost(context.Context, *UpsertHostRequest) (*UpsertHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertHost not implemented")
}
func (UnimplementedHostsServer) DeleteHost(context.Context, *DeleteHostRequest) (*DeleteHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHost not implemented")
}
func (UnimplementedHostsServer) mustEmbedUnimplementedHostsServer() {}

// UnsafeHostsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostsServer will
// result in compilation errors.
type UnsafeHostsServer interface {
	mustEmbedUnimplementedHostsServer()
}

func RegisterHostsServer(s grpc.ServiceRegistrar, srv HostsServer) {
	s.RegisterService(&Hosts_ServiceDesc, srv)
}

func _Hosts_CreateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).CreateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_CreateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).CreateHost(ctx, req.(*CreateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hosts_GetHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).GetHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_GetHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).GetHost(ctx, req.(*ReadHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hosts_GetHostMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).GetHostMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_GetHostMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).GetHostMany(ctx, req.(*ReadHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hosts_UpsertHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).UpsertHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_UpsertHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).UpsertHost(ctx, req.(*UpsertHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hosts_DeleteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).DeleteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hosts_DeleteHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).DeleteHost(ctx, req.(*DeleteHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hosts_ServiceDesc is the grpc.ServiceDesc for Hosts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hosts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hosts.Hosts",
	HandlerType: (*HostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHost",
			Handler:    _Hosts_CreateHost_Handler,
		},
		{
			MethodName: "GetHost",
			Handler:    _Hosts_GetHost_Handler,
		},
		{
			MethodName: "GetHostMany",
			Handler:    _Hosts_GetHostMany_Handler,
		},
		{
			MethodName: "UpsertHost",
			Handler:    _Hosts_UpsertHost_Handler,
		},
		{
			MethodName: "DeleteHost",
			Handler:    _Hosts_DeleteHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/hosts/hosts.proto",
}
