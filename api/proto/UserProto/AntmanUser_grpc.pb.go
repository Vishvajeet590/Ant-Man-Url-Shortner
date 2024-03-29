// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Antman

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

// AntmanUserRoutesClient is the client API for AntmanUserRoutes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AntmanUserRoutesClient interface {
	CreateNewUser(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetUrlStat(ctx context.Context, in *GetStatRequest, opts ...grpc.CallOption) (*GetStatResponse, error)
	GetUrlStatList(ctx context.Context, in *GetStatListRequest, opts ...grpc.CallOption) (*GetStatListResponse, error)
}

type antmanUserRoutesClient struct {
	cc grpc.ClientConnInterface
}

func NewAntmanUserRoutesClient(cc grpc.ClientConnInterface) AntmanUserRoutesClient {
	return &antmanUserRoutesClient{cc}
}

func (c *antmanUserRoutesClient) CreateNewUser(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/AntmanServer.AntmanUserRoutes/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antmanUserRoutesClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/AntmanServer.AntmanUserRoutes/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antmanUserRoutesClient) GetUrlStat(ctx context.Context, in *GetStatRequest, opts ...grpc.CallOption) (*GetStatResponse, error) {
	out := new(GetStatResponse)
	err := c.cc.Invoke(ctx, "/AntmanServer.AntmanUserRoutes/GetUrlStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antmanUserRoutesClient) GetUrlStatList(ctx context.Context, in *GetStatListRequest, opts ...grpc.CallOption) (*GetStatListResponse, error) {
	out := new(GetStatListResponse)
	err := c.cc.Invoke(ctx, "/AntmanServer.AntmanUserRoutes/GetUrlStatList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntmanUserRoutesServer is the server API for AntmanUserRoutes service.
// All implementations must embed UnimplementedAntmanUserRoutesServer
// for forward compatibility
type AntmanUserRoutesServer interface {
	CreateNewUser(context.Context, *SignUpRequest) (*SignUpResponse, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUrlStat(context.Context, *GetStatRequest) (*GetStatResponse, error)
	GetUrlStatList(context.Context, *GetStatListRequest) (*GetStatListResponse, error)
	mustEmbedUnimplementedAntmanUserRoutesServer()
}

// UnimplementedAntmanUserRoutesServer must be embedded to have forward compatible implementations.
type UnimplementedAntmanUserRoutesServer struct {
}

func (UnimplementedAntmanUserRoutesServer) CreateNewUser(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedAntmanUserRoutesServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAntmanUserRoutesServer) GetUrlStat(context.Context, *GetStatRequest) (*GetStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrlStat not implemented")
}
func (UnimplementedAntmanUserRoutesServer) GetUrlStatList(context.Context, *GetStatListRequest) (*GetStatListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrlStatList not implemented")
}
func (UnimplementedAntmanUserRoutesServer) mustEmbedUnimplementedAntmanUserRoutesServer() {}

// UnsafeAntmanUserRoutesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AntmanUserRoutesServer will
// result in compilation errors.
type UnsafeAntmanUserRoutesServer interface {
	mustEmbedUnimplementedAntmanUserRoutesServer()
}

func RegisterAntmanUserRoutesServer(s grpc.ServiceRegistrar, srv AntmanUserRoutesServer) {
	s.RegisterService(&AntmanUserRoutes_ServiceDesc, srv)
}

func _AntmanUserRoutes_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntmanUserRoutesServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AntmanServer.AntmanUserRoutes/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntmanUserRoutesServer).CreateNewUser(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntmanUserRoutes_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntmanUserRoutesServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AntmanServer.AntmanUserRoutes/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntmanUserRoutesServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntmanUserRoutes_GetUrlStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntmanUserRoutesServer).GetUrlStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AntmanServer.AntmanUserRoutes/GetUrlStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntmanUserRoutesServer).GetUrlStat(ctx, req.(*GetStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntmanUserRoutes_GetUrlStatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntmanUserRoutesServer).GetUrlStatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AntmanServer.AntmanUserRoutes/GetUrlStatList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntmanUserRoutesServer).GetUrlStatList(ctx, req.(*GetStatListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AntmanUserRoutes_ServiceDesc is the grpc.ServiceDesc for AntmanUserRoutes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AntmanUserRoutes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AntmanServer.AntmanUserRoutes",
	HandlerType: (*AntmanUserRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewUser",
			Handler:    _AntmanUserRoutes_CreateNewUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AntmanUserRoutes_LoginUser_Handler,
		},
		{
			MethodName: "GetUrlStat",
			Handler:    _AntmanUserRoutes_GetUrlStat_Handler,
		},
		{
			MethodName: "GetUrlStatList",
			Handler:    _AntmanUserRoutes_GetUrlStatList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/UserProto/AntmanUser.proto",
}
