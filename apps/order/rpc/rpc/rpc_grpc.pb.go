// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: rpc.proto

package rpc

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

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Rpc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Rpc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Rpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
