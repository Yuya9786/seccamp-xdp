// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: protobuf/scmlb.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ScmLbApi_Health_FullMethodName = "/scmlb.v1.ScmLbApi/Health"
	ScmLbApi_Stat_FullMethodName   = "/scmlb.v1.ScmLbApi/Stat"
)

// ScmLbApiClient is the client API for ScmLbApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScmLbApiClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
}

type scmLbApiClient struct {
	cc grpc.ClientConnInterface
}

func NewScmLbApiClient(cc grpc.ClientConnInterface) ScmLbApiClient {
	return &scmLbApiClient{cc}
}

func (c *scmLbApiClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ScmLbApi_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scmLbApiClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, ScmLbApi_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScmLbApiServer is the server API for ScmLbApi service.
// All implementations must embed UnimplementedScmLbApiServer
// for forward compatibility
type ScmLbApiServer interface {
	Health(context.Context, *HealthRequest) (*emptypb.Empty, error)
	Stat(context.Context, *StatRequest) (*StatResponse, error)
	mustEmbedUnimplementedScmLbApiServer()
}

// UnimplementedScmLbApiServer must be embedded to have forward compatible implementations.
type UnimplementedScmLbApiServer struct {
}

func (UnimplementedScmLbApiServer) Health(context.Context, *HealthRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedScmLbApiServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedScmLbApiServer) mustEmbedUnimplementedScmLbApiServer() {}

// UnsafeScmLbApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScmLbApiServer will
// result in compilation errors.
type UnsafeScmLbApiServer interface {
	mustEmbedUnimplementedScmLbApiServer()
}

func RegisterScmLbApiServer(s grpc.ServiceRegistrar, srv ScmLbApiServer) {
	s.RegisterService(&ScmLbApi_ServiceDesc, srv)
}

func _ScmLbApi_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScmLbApiServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScmLbApi_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScmLbApiServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScmLbApi_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScmLbApiServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScmLbApi_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScmLbApiServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScmLbApi_ServiceDesc is the grpc.ServiceDesc for ScmLbApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScmLbApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scmlb.v1.ScmLbApi",
	HandlerType: (*ScmLbApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _ScmLbApi_Health_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _ScmLbApi_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/scmlb.proto",
}
