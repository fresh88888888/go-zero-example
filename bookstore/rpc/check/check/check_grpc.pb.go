// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: check.proto

package check

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
	Check_Ping_FullMethodName = "/check.Check/Ping"
)

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type checkClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckClient(cc grpc.ClientConnInterface) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Check_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
// All implementations must embed UnimplementedCheckServer
// for forward compatibility
type CheckServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCheckServer()
}

// UnimplementedCheckServer must be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (UnimplementedCheckServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCheckServer) mustEmbedUnimplementedCheckServer() {}

// UnsafeCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckServer will
// result in compilation errors.
type UnsafeCheckServer interface {
	mustEmbedUnimplementedCheckServer()
}

func RegisterCheckServer(s grpc.ServiceRegistrar, srv CheckServer) {
	s.RegisterService(&Check_ServiceDesc, srv)
}

func _Check_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Check_ServiceDesc is the grpc.ServiceDesc for Check service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Check_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "check.Check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Check_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "check.proto",
}
