// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: servermanager.proto

package pb

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
	Servermanager_CreateServer_FullMethodName  = "/pb.servermanager/createServer"
	Servermanager_SetServer_FullMethodName     = "/pb.servermanager/setServer"
	Servermanager_GetServer_FullMethodName     = "/pb.servermanager/getServer"
	Servermanager_LoginServer_FullMethodName   = "/pb.servermanager/loginServer"
	Servermanager_GetServerList_FullMethodName = "/pb.servermanager/getServerList"
)

// ServermanagerClient is the client API for Servermanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServermanagerClient interface {
	CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error)
	SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error)
	GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error)
	LoginServer(ctx context.Context, in *LoginServerReq, opts ...grpc.CallOption) (*LoginServerResp, error)
	GetServerList(ctx context.Context, in *GetServerListReq, opts ...grpc.CallOption) (*GetServerListResp, error)
}

type servermanagerClient struct {
	cc grpc.ClientConnInterface
}

func NewServermanagerClient(cc grpc.ClientConnInterface) ServermanagerClient {
	return &servermanagerClient{cc}
}

func (c *servermanagerClient) CreateServer(ctx context.Context, in *CreateServerReq, opts ...grpc.CallOption) (*CreateServerResp, error) {
	out := new(CreateServerResp)
	err := c.cc.Invoke(ctx, Servermanager_CreateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servermanagerClient) SetServer(ctx context.Context, in *SetServerReq, opts ...grpc.CallOption) (*SetServerResp, error) {
	out := new(SetServerResp)
	err := c.cc.Invoke(ctx, Servermanager_SetServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servermanagerClient) GetServer(ctx context.Context, in *GetServerReq, opts ...grpc.CallOption) (*GetServerResp, error) {
	out := new(GetServerResp)
	err := c.cc.Invoke(ctx, Servermanager_GetServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servermanagerClient) LoginServer(ctx context.Context, in *LoginServerReq, opts ...grpc.CallOption) (*LoginServerResp, error) {
	out := new(LoginServerResp)
	err := c.cc.Invoke(ctx, Servermanager_LoginServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servermanagerClient) GetServerList(ctx context.Context, in *GetServerListReq, opts ...grpc.CallOption) (*GetServerListResp, error) {
	out := new(GetServerListResp)
	err := c.cc.Invoke(ctx, Servermanager_GetServerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServermanagerServer is the server API for Servermanager service.
// All implementations must embed UnimplementedServermanagerServer
// for forward compatibility
type ServermanagerServer interface {
	CreateServer(context.Context, *CreateServerReq) (*CreateServerResp, error)
	SetServer(context.Context, *SetServerReq) (*SetServerResp, error)
	GetServer(context.Context, *GetServerReq) (*GetServerResp, error)
	LoginServer(context.Context, *LoginServerReq) (*LoginServerResp, error)
	GetServerList(context.Context, *GetServerListReq) (*GetServerListResp, error)
	mustEmbedUnimplementedServermanagerServer()
}

// UnimplementedServermanagerServer must be embedded to have forward compatible implementations.
type UnimplementedServermanagerServer struct {
}

func (UnimplementedServermanagerServer) CreateServer(context.Context, *CreateServerReq) (*CreateServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedServermanagerServer) SetServer(context.Context, *SetServerReq) (*SetServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetServer not implemented")
}
func (UnimplementedServermanagerServer) GetServer(context.Context, *GetServerReq) (*GetServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedServermanagerServer) LoginServer(context.Context, *LoginServerReq) (*LoginServerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginServer not implemented")
}
func (UnimplementedServermanagerServer) GetServerList(context.Context, *GetServerListReq) (*GetServerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (UnimplementedServermanagerServer) mustEmbedUnimplementedServermanagerServer() {}

// UnsafeServermanagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServermanagerServer will
// result in compilation errors.
type UnsafeServermanagerServer interface {
	mustEmbedUnimplementedServermanagerServer()
}

func RegisterServermanagerServer(s grpc.ServiceRegistrar, srv ServermanagerServer) {
	s.RegisterService(&Servermanager_ServiceDesc, srv)
}

func _Servermanager_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServermanagerServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Servermanager_CreateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServermanagerServer).CreateServer(ctx, req.(*CreateServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servermanager_SetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServermanagerServer).SetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Servermanager_SetServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServermanagerServer).SetServer(ctx, req.(*SetServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servermanager_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServermanagerServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Servermanager_GetServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServermanagerServer).GetServer(ctx, req.(*GetServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servermanager_LoginServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServermanagerServer).LoginServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Servermanager_LoginServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServermanagerServer).LoginServer(ctx, req.(*LoginServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servermanager_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServermanagerServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Servermanager_GetServerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServermanagerServer).GetServerList(ctx, req.(*GetServerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Servermanager_ServiceDesc is the grpc.ServiceDesc for Servermanager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Servermanager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.servermanager",
	HandlerType: (*ServermanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createServer",
			Handler:    _Servermanager_CreateServer_Handler,
		},
		{
			MethodName: "setServer",
			Handler:    _Servermanager_SetServer_Handler,
		},
		{
			MethodName: "getServer",
			Handler:    _Servermanager_GetServer_Handler,
		},
		{
			MethodName: "loginServer",
			Handler:    _Servermanager_LoginServer_Handler,
		},
		{
			MethodName: "getServerList",
			Handler:    _Servermanager_GetServerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servermanager.proto",
}
