// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: auth_service.proto

package authservicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_PermissionCreate_FullMethodName = "/authservice.AuthService/PermissionCreate"
	AuthService_PermissionRead_FullMethodName   = "/authservice.AuthService/PermissionRead"
	AuthService_PermissionUpdate_FullMethodName = "/authservice.AuthService/PermissionUpdate"
	AuthService_PermissionDelete_FullMethodName = "/authservice.AuthService/PermissionDelete"
	AuthService_RoleCreate_FullMethodName       = "/authservice.AuthService/RoleCreate"
	AuthService_RoleRead_FullMethodName         = "/authservice.AuthService/RoleRead"
	AuthService_RoleUpdate_FullMethodName       = "/authservice.AuthService/RoleUpdate"
	AuthService_RoleDelete_FullMethodName       = "/authservice.AuthService/RoleDelete"
	AuthService_UserRegister_FullMethodName     = "/authservice.AuthService/UserRegister"
	AuthService_UserRead_FullMethodName         = "/authservice.AuthService/UserRead"
	AuthService_UserUpdate_FullMethodName       = "/authservice.AuthService/UserUpdate"
	AuthService_UserDelete_FullMethodName       = "/authservice.AuthService/UserDelete"
	AuthService_Login_FullMethodName            = "/authservice.AuthService/Login"
	AuthService_Logout_FullMethodName           = "/authservice.AuthService/Logout"
	AuthService_ValidateToken_FullMethodName    = "/authservice.AuthService/ValidateToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Auth is service for managing permissions and roles.
type AuthServiceClient interface {
	// Permissions messages
	PermissionCreate(ctx context.Context, in *PermissionCreateRequest, opts ...grpc.CallOption) (*PermissionCreateResponse, error)
	PermissionRead(ctx context.Context, in *PermissionReadRequest, opts ...grpc.CallOption) (*PermissionReadResponse, error)
	PermissionUpdate(ctx context.Context, in *PermissionUpdateRequest, opts ...grpc.CallOption) (*PermissionUpdateResponse, error)
	PermissionDelete(ctx context.Context, in *PermissionDeleteRequest, opts ...grpc.CallOption) (*PermissionDeleteResponse, error)
	// Role messages
	RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	RoleRead(ctx context.Context, in *RoleReadRequest, opts ...grpc.CallOption) (*RoleReadResponse, error)
	RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error)
	RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	// User messages
	UserRegister(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserRead(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	// Login logs in a user and returns an auth token.
	Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// Logout logs out a user and revokes their auth token.
	Logout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error)
	// ValidateToken validates a user's auth token.
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) PermissionCreate(ctx context.Context, in *PermissionCreateRequest, opts ...grpc.CallOption) (*PermissionCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionCreateResponse)
	err := c.cc.Invoke(ctx, AuthService_PermissionCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PermissionRead(ctx context.Context, in *PermissionReadRequest, opts ...grpc.CallOption) (*PermissionReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionReadResponse)
	err := c.cc.Invoke(ctx, AuthService_PermissionRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PermissionUpdate(ctx context.Context, in *PermissionUpdateRequest, opts ...grpc.CallOption) (*PermissionUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionUpdateResponse)
	err := c.cc.Invoke(ctx, AuthService_PermissionUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PermissionDelete(ctx context.Context, in *PermissionDeleteRequest, opts ...grpc.CallOption) (*PermissionDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionDeleteResponse)
	err := c.cc.Invoke(ctx, AuthService_PermissionDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, AuthService_RoleCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RoleRead(ctx context.Context, in *RoleReadRequest, opts ...grpc.CallOption) (*RoleReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleReadResponse)
	err := c.cc.Invoke(ctx, AuthService_RoleRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleUpdateResponse)
	err := c.cc.Invoke(ctx, AuthService_RoleUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, AuthService_RoleDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserRegister(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, AuthService_UserRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserRead(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReadResponse)
	err := c.cc.Invoke(ctx, AuthService_UserRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, AuthService_UserUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, AuthService_UserDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLogoutResponse)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
//
// Auth is service for managing permissions and roles.
type AuthServiceServer interface {
	// Permissions messages
	PermissionCreate(context.Context, *PermissionCreateRequest) (*PermissionCreateResponse, error)
	PermissionRead(context.Context, *PermissionReadRequest) (*PermissionReadResponse, error)
	PermissionUpdate(context.Context, *PermissionUpdateRequest) (*PermissionUpdateResponse, error)
	PermissionDelete(context.Context, *PermissionDeleteRequest) (*PermissionDeleteResponse, error)
	// Role messages
	RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	RoleRead(context.Context, *RoleReadRequest) (*RoleReadResponse, error)
	RoleUpdate(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error)
	RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	// User messages
	UserRegister(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserRead(context.Context, *UserReadRequest) (*UserReadResponse, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	// Login logs in a user and returns an auth token.
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// Logout logs out a user and revokes their auth token.
	Logout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error)
	// ValidateToken validates a user's auth token.
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) PermissionCreate(context.Context, *PermissionCreateRequest) (*PermissionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionCreate not implemented")
}
func (UnimplementedAuthServiceServer) PermissionRead(context.Context, *PermissionReadRequest) (*PermissionReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionRead not implemented")
}
func (UnimplementedAuthServiceServer) PermissionUpdate(context.Context, *PermissionUpdateRequest) (*PermissionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionUpdate not implemented")
}
func (UnimplementedAuthServiceServer) PermissionDelete(context.Context, *PermissionDeleteRequest) (*PermissionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionDelete not implemented")
}
func (UnimplementedAuthServiceServer) RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCreate not implemented")
}
func (UnimplementedAuthServiceServer) RoleRead(context.Context, *RoleReadRequest) (*RoleReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleRead not implemented")
}
func (UnimplementedAuthServiceServer) RoleUpdate(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleUpdate not implemented")
}
func (UnimplementedAuthServiceServer) RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}
func (UnimplementedAuthServiceServer) UserRegister(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedAuthServiceServer) UserRead(context.Context, *UserReadRequest) (*UserReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRead not implemented")
}
func (UnimplementedAuthServiceServer) UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedAuthServiceServer) UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_PermissionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PermissionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PermissionCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PermissionCreate(ctx, req.(*PermissionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PermissionRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PermissionRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PermissionRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PermissionRead(ctx, req.(*PermissionReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PermissionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PermissionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PermissionUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PermissionUpdate(ctx, req.(*PermissionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PermissionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PermissionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PermissionDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PermissionDelete(ctx, req.(*PermissionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RoleCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RoleCreate(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RoleRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RoleRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RoleRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RoleRead(ctx, req.(*RoleReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RoleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RoleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RoleUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RoleUpdate(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RoleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RoleDelete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserRegister(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserRead(ctx, req.(*UserReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*UserLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authservice.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermissionCreate",
			Handler:    _AuthService_PermissionCreate_Handler,
		},
		{
			MethodName: "PermissionRead",
			Handler:    _AuthService_PermissionRead_Handler,
		},
		{
			MethodName: "PermissionUpdate",
			Handler:    _AuthService_PermissionUpdate_Handler,
		},
		{
			MethodName: "PermissionDelete",
			Handler:    _AuthService_PermissionDelete_Handler,
		},
		{
			MethodName: "RoleCreate",
			Handler:    _AuthService_RoleCreate_Handler,
		},
		{
			MethodName: "RoleRead",
			Handler:    _AuthService_RoleRead_Handler,
		},
		{
			MethodName: "RoleUpdate",
			Handler:    _AuthService_RoleUpdate_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _AuthService_RoleDelete_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _AuthService_UserRegister_Handler,
		},
		{
			MethodName: "UserRead",
			Handler:    _AuthService_UserRead_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _AuthService_UserUpdate_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _AuthService_UserDelete_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
