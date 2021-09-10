// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.5
// source: api/tool.proto

package api

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

// ToolServiceClient is the client API for ToolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToolServiceClient interface {
	Hash(ctx context.Context, in *HashReq, opts ...grpc.CallOption) (*HashRes, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type toolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToolServiceClient(cc grpc.ClientConnInterface) ToolServiceClient {
	return &toolServiceClient{cc}
}

func (c *toolServiceClient) Hash(ctx context.Context, in *HashReq, opts ...grpc.CallOption) (*HashRes, error) {
	out := new(HashRes)
	err := c.cc.Invoke(ctx, "/api.ToolService/Hash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.ToolService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToolServiceServer is the server API for ToolService service.
// All implementations must embed UnimplementedToolServiceServer
// for forward compatibility
type ToolServiceServer interface {
	Hash(context.Context, *HashReq) (*HashRes, error)
	Ping(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedToolServiceServer()
}

// UnimplementedToolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToolServiceServer struct {
}

func (UnimplementedToolServiceServer) Hash(context.Context, *HashReq) (*HashRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash not implemented")
}
func (UnimplementedToolServiceServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedToolServiceServer) mustEmbedUnimplementedToolServiceServer() {}

// UnsafeToolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToolServiceServer will
// result in compilation errors.
type UnsafeToolServiceServer interface {
	mustEmbedUnimplementedToolServiceServer()
}

func RegisterToolServiceServer(s grpc.ServiceRegistrar, srv ToolServiceServer) {
	s.RegisterService(&ToolService_ServiceDesc, srv)
}

func _ToolService_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToolService/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).Hash(ctx, req.(*HashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToolService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ToolService_ServiceDesc is the grpc.ServiceDesc for ToolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ToolService",
	HandlerType: (*ToolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _ToolService_Hash_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ToolService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tool.proto",
}
