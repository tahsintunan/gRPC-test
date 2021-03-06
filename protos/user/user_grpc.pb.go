// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserAuthClient is the client API for UserAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthClient interface {
	Register(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*ApiRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*ApiRes, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*ApiRes, error)
}

type userAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthClient(cc grpc.ClientConnInterface) UserAuthClient {
	return &userAuthClient{cc}
}

func (c *userAuthClient) Register(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*ApiRes, error) {
	out := new(ApiRes)
	err := c.cc.Invoke(ctx, "/user.userAuth/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*ApiRes, error) {
	out := new(ApiRes)
	err := c.cc.Invoke(ctx, "/user.userAuth/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*ApiRes, error) {
	out := new(ApiRes)
	err := c.cc.Invoke(ctx, "/user.userAuth/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthServer is the server API for UserAuth service.
// All implementations must embed UnimplementedUserAuthServer
// for forward compatibility
type UserAuthServer interface {
	Register(context.Context, *RegReq) (*ApiRes, error)
	Login(context.Context, *LoginReq) (*ApiRes, error)
	Logout(context.Context, *LogoutReq) (*ApiRes, error)
	mustEmbedUnimplementedUserAuthServer()
}

// UnimplementedUserAuthServer must be embedded to have forward compatible implementations.
type UnimplementedUserAuthServer struct {
}

func (UnimplementedUserAuthServer) Register(context.Context, *RegReq) (*ApiRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserAuthServer) Login(context.Context, *LoginReq) (*ApiRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserAuthServer) Logout(context.Context, *LogoutReq) (*ApiRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserAuthServer) mustEmbedUnimplementedUserAuthServer() {}

// UnsafeUserAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthServer will
// result in compilation errors.
type UnsafeUserAuthServer interface {
	mustEmbedUnimplementedUserAuthServer()
}

func RegisterUserAuthServer(s grpc.ServiceRegistrar, srv UserAuthServer) {
	s.RegisterService(&UserAuth_ServiceDesc, srv)
}

func _UserAuth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userAuth/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).Register(ctx, req.(*RegReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userAuth/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userAuth/logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuth_ServiceDesc is the grpc.ServiceDesc for UserAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.userAuth",
	HandlerType: (*UserAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _UserAuth_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _UserAuth_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _UserAuth_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user/user.proto",
}
