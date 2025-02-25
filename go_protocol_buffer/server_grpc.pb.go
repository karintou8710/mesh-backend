// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/server.proto

package go_protocol_buffer

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	AnonymousSignUp(ctx context.Context, in *AnonymousSignUpRequest, opts ...grpc.CallOption) (*AnonymousSignUpResponse, error)
	CreateShareGroup(ctx context.Context, in *CreateShareGroupRequest, opts ...grpc.CallOption) (*CreateShareGroupResponse, error)
	JoinShareGroup(ctx context.Context, in *JoinShareGroupRequest, opts ...grpc.CallOption) (*JoinShareGroupResponse, error)
	GetCurrentShareGroup(ctx context.Context, in *GetCurrentShareGroupRequest, opts ...grpc.CallOption) (*GetCurrentShareGroupResponse, error)
	UpdatePosition(ctx context.Context, in *UpdatePositionGroupRequest, opts ...grpc.CallOption) (*UpdatePositionGroupResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AnonymousSignUp(ctx context.Context, in *AnonymousSignUpRequest, opts ...grpc.CallOption) (*AnonymousSignUpResponse, error) {
	out := new(AnonymousSignUpResponse)
	err := c.cc.Invoke(ctx, "/Server.Service/AnonymousSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateShareGroup(ctx context.Context, in *CreateShareGroupRequest, opts ...grpc.CallOption) (*CreateShareGroupResponse, error) {
	out := new(CreateShareGroupResponse)
	err := c.cc.Invoke(ctx, "/Server.Service/CreateShareGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JoinShareGroup(ctx context.Context, in *JoinShareGroupRequest, opts ...grpc.CallOption) (*JoinShareGroupResponse, error) {
	out := new(JoinShareGroupResponse)
	err := c.cc.Invoke(ctx, "/Server.Service/JoinShareGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetCurrentShareGroup(ctx context.Context, in *GetCurrentShareGroupRequest, opts ...grpc.CallOption) (*GetCurrentShareGroupResponse, error) {
	out := new(GetCurrentShareGroupResponse)
	err := c.cc.Invoke(ctx, "/Server.Service/GetCurrentShareGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdatePosition(ctx context.Context, in *UpdatePositionGroupRequest, opts ...grpc.CallOption) (*UpdatePositionGroupResponse, error) {
	out := new(UpdatePositionGroupResponse)
	err := c.cc.Invoke(ctx, "/Server.Service/UpdatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	AnonymousSignUp(context.Context, *AnonymousSignUpRequest) (*AnonymousSignUpResponse, error)
	CreateShareGroup(context.Context, *CreateShareGroupRequest) (*CreateShareGroupResponse, error)
	JoinShareGroup(context.Context, *JoinShareGroupRequest) (*JoinShareGroupResponse, error)
	GetCurrentShareGroup(context.Context, *GetCurrentShareGroupRequest) (*GetCurrentShareGroupResponse, error)
	UpdatePosition(context.Context, *UpdatePositionGroupRequest) (*UpdatePositionGroupResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) AnonymousSignUp(context.Context, *AnonymousSignUpRequest) (*AnonymousSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnonymousSignUp not implemented")
}
func (UnimplementedServiceServer) CreateShareGroup(context.Context, *CreateShareGroupRequest) (*CreateShareGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShareGroup not implemented")
}
func (UnimplementedServiceServer) JoinShareGroup(context.Context, *JoinShareGroupRequest) (*JoinShareGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinShareGroup not implemented")
}
func (UnimplementedServiceServer) GetCurrentShareGroup(context.Context, *GetCurrentShareGroupRequest) (*GetCurrentShareGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentShareGroup not implemented")
}
func (UnimplementedServiceServer) UpdatePosition(context.Context, *UpdatePositionGroupRequest) (*UpdatePositionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosition not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AnonymousSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymousSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AnonymousSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.Service/AnonymousSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AnonymousSignUp(ctx, req.(*AnonymousSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateShareGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShareGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateShareGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.Service/CreateShareGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateShareGroup(ctx, req.(*CreateShareGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JoinShareGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinShareGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JoinShareGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.Service/JoinShareGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JoinShareGroup(ctx, req.(*JoinShareGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetCurrentShareGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentShareGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetCurrentShareGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.Service/GetCurrentShareGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetCurrentShareGroup(ctx, req.(*GetCurrentShareGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.Service/UpdatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdatePosition(ctx, req.(*UpdatePositionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnonymousSignUp",
			Handler:    _Service_AnonymousSignUp_Handler,
		},
		{
			MethodName: "CreateShareGroup",
			Handler:    _Service_CreateShareGroup_Handler,
		},
		{
			MethodName: "JoinShareGroup",
			Handler:    _Service_JoinShareGroup_Handler,
		},
		{
			MethodName: "GetCurrentShareGroup",
			Handler:    _Service_GetCurrentShareGroup_Handler,
		},
		{
			MethodName: "UpdatePosition",
			Handler:    _Service_UpdatePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server.proto",
}
