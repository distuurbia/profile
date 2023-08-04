// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services.proto

package profile

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error)
	GetPasswordAndIDByUsername(ctx context.Context, in *GetPasswordAndIDByUsernameRequest, opts ...grpc.CallOption) (*GetPasswordAndIDByUsernameResponse, error)
	GetRefreshTokenByID(ctx context.Context, in *GetRefreshTokenByIDRequest, opts ...grpc.CallOption) (*GetRefreshTokenByIDResponse, error)
	AddRefreshToken(ctx context.Context, in *AddRefreshTokenRequest, opts ...grpc.CallOption) (*AddRefreshTokenResponse, error)
	DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileResponse, error) {
	out := new(CreateProfileResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetPasswordAndIDByUsername(ctx context.Context, in *GetPasswordAndIDByUsernameRequest, opts ...grpc.CallOption) (*GetPasswordAndIDByUsernameResponse, error) {
	out := new(GetPasswordAndIDByUsernameResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/GetPasswordAndIDByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetRefreshTokenByID(ctx context.Context, in *GetRefreshTokenByIDRequest, opts ...grpc.CallOption) (*GetRefreshTokenByIDResponse, error) {
	out := new(GetRefreshTokenByIDResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/GetRefreshTokenByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddRefreshToken(ctx context.Context, in *AddRefreshTokenRequest, opts ...grpc.CallOption) (*AddRefreshTokenResponse, error) {
	out := new(AddRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/AddRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error) {
	out := new(DeleteProfileResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/DeleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
	GetPasswordAndIDByUsername(context.Context, *GetPasswordAndIDByUsernameRequest) (*GetPasswordAndIDByUsernameResponse, error)
	GetRefreshTokenByID(context.Context, *GetRefreshTokenByIDRequest) (*GetRefreshTokenByIDResponse, error)
	AddRefreshToken(context.Context, *AddRefreshTokenRequest) (*AddRefreshTokenResponse, error)
	DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetPasswordAndIDByUsername(context.Context, *GetPasswordAndIDByUsernameRequest) (*GetPasswordAndIDByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordAndIDByUsername not implemented")
}
func (UnimplementedProfileServiceServer) GetRefreshTokenByID(context.Context, *GetRefreshTokenByIDRequest) (*GetRefreshTokenByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshTokenByID not implemented")
}
func (UnimplementedProfileServiceServer) AddRefreshToken(context.Context, *AddRefreshTokenRequest) (*AddRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRefreshToken not implemented")
}
func (UnimplementedProfileServiceServer) DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetPasswordAndIDByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordAndIDByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetPasswordAndIDByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/GetPasswordAndIDByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetPasswordAndIDByUsername(ctx, req.(*GetPasswordAndIDByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetRefreshTokenByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefreshTokenByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetRefreshTokenByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/GetRefreshTokenByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetRefreshTokenByID(ctx, req.(*GetRefreshTokenByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/AddRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddRefreshToken(ctx, req.(*AddRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).DeleteProfile(ctx, req.(*DeleteProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _ProfileService_CreateProfile_Handler,
		},
		{
			MethodName: "GetPasswordAndIDByUsername",
			Handler:    _ProfileService_GetPasswordAndIDByUsername_Handler,
		},
		{
			MethodName: "GetRefreshTokenByID",
			Handler:    _ProfileService_GetRefreshTokenByID_Handler,
		},
		{
			MethodName: "AddRefreshToken",
			Handler:    _ProfileService_AddRefreshToken_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _ProfileService_DeleteProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}