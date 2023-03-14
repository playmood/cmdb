// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: apps/secret/pb/secret.proto

package secret

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
	Service_CreateSecret_FullMethodName   = "/playmood.cmdb.secret.Service/CreateSecret"
	Service_QuerySecret_FullMethodName    = "/playmood.cmdb.secret.Service/QuerySecret"
	Service_DescribeSecret_FullMethodName = "/playmood.cmdb.secret.Service/DescribeSecret"
	Service_DeleteSecret_FullMethodName   = "/playmood.cmdb.secret.Service/DeleteSecret"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	QuerySecret(ctx context.Context, in *QuerySecretRequest, opts ...grpc.CallOption) (*SecretSet, error)
	DescribeSecret(ctx context.Context, in *DescribeSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*Secret, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Service_CreateSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QuerySecret(ctx context.Context, in *QuerySecretRequest, opts ...grpc.CallOption) (*SecretSet, error) {
	out := new(SecretSet)
	err := c.cc.Invoke(ctx, Service_QuerySecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeSecret(ctx context.Context, in *DescribeSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Service_DescribeSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, Service_DeleteSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedServiceServer) QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySecret not implemented")
}
func (UnimplementedServiceServer) DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSecret not implemented")
}
func (UnimplementedServiceServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
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

func _Service_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QuerySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QuerySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_QuerySecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QuerySecret(ctx, req.(*QuerySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DescribeSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeSecret(ctx, req.(*DescribeSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "playmood.cmdb.secret.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _Service_CreateSecret_Handler,
		},
		{
			MethodName: "QuerySecret",
			Handler:    _Service_QuerySecret_Handler,
		},
		{
			MethodName: "DescribeSecret",
			Handler:    _Service_DescribeSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _Service_DeleteSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/secret/pb/secret.proto",
}
