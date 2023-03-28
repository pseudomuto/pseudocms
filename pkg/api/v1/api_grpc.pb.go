// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1/api.proto

package v1

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
	HealthService_Ping_FullMethodName = "/api.v1.HealthService/Ping"
)

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	// Ping is used to test that the service is up and responding to requests.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, HealthService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations should embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	// Ping is used to test that the service is up and responding to requests.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedHealthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HealthService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}

const (
	AdminService_CreateDefinition_FullMethodName = "/api.v1.AdminService/CreateDefinition"
	AdminService_GetDefinition_FullMethodName    = "/api.v1.AdminService/GetDefinition"
	AdminService_CreateField_FullMethodName      = "/api.v1.AdminService/CreateField"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	// CreateDefinition creates a new definition object.
	CreateDefinition(ctx context.Context, in *CreateDefinitionRequest, opts ...grpc.CallOption) (*CreateDefinitionResponse, error)
	GetDefinition(ctx context.Context, in *GetDefinitionRequest, opts ...grpc.CallOption) (*GetDefinitionResponse, error)
	CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateDefinition(ctx context.Context, in *CreateDefinitionRequest, opts ...grpc.CallOption) (*CreateDefinitionResponse, error) {
	out := new(CreateDefinitionResponse)
	err := c.cc.Invoke(ctx, AdminService_CreateDefinition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetDefinition(ctx context.Context, in *GetDefinitionRequest, opts ...grpc.CallOption) (*GetDefinitionResponse, error) {
	out := new(GetDefinitionResponse)
	err := c.cc.Invoke(ctx, AdminService_GetDefinition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error) {
	out := new(CreateFieldResponse)
	err := c.cc.Invoke(ctx, AdminService_CreateField_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations should embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	// CreateDefinition creates a new definition object.
	CreateDefinition(context.Context, *CreateDefinitionRequest) (*CreateDefinitionResponse, error)
	GetDefinition(context.Context, *GetDefinitionRequest) (*GetDefinitionResponse, error)
	CreateField(context.Context, *CreateFieldRequest) (*CreateFieldResponse, error)
}

// UnimplementedAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) CreateDefinition(context.Context, *CreateDefinitionRequest) (*CreateDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefinition not implemented")
}
func (UnimplementedAdminServiceServer) GetDefinition(context.Context, *GetDefinitionRequest) (*GetDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefinition not implemented")
}
func (UnimplementedAdminServiceServer) CreateField(context.Context, *CreateFieldRequest) (*CreateFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateField not implemented")
}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_CreateDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateDefinition(ctx, req.(*CreateDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetDefinition(ctx, req.(*GetDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateField(ctx, req.(*CreateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDefinition",
			Handler:    _AdminService_CreateDefinition_Handler,
		},
		{
			MethodName: "GetDefinition",
			Handler:    _AdminService_GetDefinition_Handler,
		},
		{
			MethodName: "CreateField",
			Handler:    _AdminService_CreateField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}