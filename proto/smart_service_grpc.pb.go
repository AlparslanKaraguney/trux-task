// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: proto/smart_service.proto

package smartservice

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
	SmartService_CreateSmartModel_FullMethodName   = "/smartservice.SmartService/CreateSmartModel"
	SmartService_GetSmartModel_FullMethodName      = "/smartservice.SmartService/GetSmartModel"
	SmartService_CreateSmartFeature_FullMethodName = "/smartservice.SmartService/CreateSmartFeature"
	SmartService_GetSmartFeature_FullMethodName    = "/smartservice.SmartService/GetSmartFeature"
)

// SmartServiceClient is the client API for SmartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartServiceClient interface {
	CreateSmartModel(ctx context.Context, in *SmartModelRequest, opts ...grpc.CallOption) (*SmartModelResponse, error)
	GetSmartModel(ctx context.Context, in *SmartModelQuery, opts ...grpc.CallOption) (*SmartModelResponse, error)
	CreateSmartFeature(ctx context.Context, in *SmartFeatureRequest, opts ...grpc.CallOption) (*SmartFeatureResponse, error)
	GetSmartFeature(ctx context.Context, in *SmartFeatureQuery, opts ...grpc.CallOption) (*SmartFeatureResponse, error)
}

type smartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartServiceClient(cc grpc.ClientConnInterface) SmartServiceClient {
	return &smartServiceClient{cc}
}

func (c *smartServiceClient) CreateSmartModel(ctx context.Context, in *SmartModelRequest, opts ...grpc.CallOption) (*SmartModelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmartModelResponse)
	err := c.cc.Invoke(ctx, SmartService_CreateSmartModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartServiceClient) GetSmartModel(ctx context.Context, in *SmartModelQuery, opts ...grpc.CallOption) (*SmartModelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmartModelResponse)
	err := c.cc.Invoke(ctx, SmartService_GetSmartModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartServiceClient) CreateSmartFeature(ctx context.Context, in *SmartFeatureRequest, opts ...grpc.CallOption) (*SmartFeatureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmartFeatureResponse)
	err := c.cc.Invoke(ctx, SmartService_CreateSmartFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartServiceClient) GetSmartFeature(ctx context.Context, in *SmartFeatureQuery, opts ...grpc.CallOption) (*SmartFeatureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmartFeatureResponse)
	err := c.cc.Invoke(ctx, SmartService_GetSmartFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartServiceServer is the server API for SmartService service.
// All implementations must embed UnimplementedSmartServiceServer
// for forward compatibility.
type SmartServiceServer interface {
	CreateSmartModel(context.Context, *SmartModelRequest) (*SmartModelResponse, error)
	GetSmartModel(context.Context, *SmartModelQuery) (*SmartModelResponse, error)
	CreateSmartFeature(context.Context, *SmartFeatureRequest) (*SmartFeatureResponse, error)
	GetSmartFeature(context.Context, *SmartFeatureQuery) (*SmartFeatureResponse, error)
	mustEmbedUnimplementedSmartServiceServer()
}

// UnimplementedSmartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSmartServiceServer struct{}

func (UnimplementedSmartServiceServer) CreateSmartModel(context.Context, *SmartModelRequest) (*SmartModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSmartModel not implemented")
}
func (UnimplementedSmartServiceServer) GetSmartModel(context.Context, *SmartModelQuery) (*SmartModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmartModel not implemented")
}
func (UnimplementedSmartServiceServer) CreateSmartFeature(context.Context, *SmartFeatureRequest) (*SmartFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSmartFeature not implemented")
}
func (UnimplementedSmartServiceServer) GetSmartFeature(context.Context, *SmartFeatureQuery) (*SmartFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmartFeature not implemented")
}
func (UnimplementedSmartServiceServer) mustEmbedUnimplementedSmartServiceServer() {}
func (UnimplementedSmartServiceServer) testEmbeddedByValue()                      {}

// UnsafeSmartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartServiceServer will
// result in compilation errors.
type UnsafeSmartServiceServer interface {
	mustEmbedUnimplementedSmartServiceServer()
}

func RegisterSmartServiceServer(s grpc.ServiceRegistrar, srv SmartServiceServer) {
	// If the following call pancis, it indicates UnimplementedSmartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SmartService_ServiceDesc, srv)
}

func _SmartService_CreateSmartModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartServiceServer).CreateSmartModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmartService_CreateSmartModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartServiceServer).CreateSmartModel(ctx, req.(*SmartModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartService_GetSmartModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartModelQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartServiceServer).GetSmartModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmartService_GetSmartModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartServiceServer).GetSmartModel(ctx, req.(*SmartModelQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartService_CreateSmartFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartServiceServer).CreateSmartFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmartService_CreateSmartFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartServiceServer).CreateSmartFeature(ctx, req.(*SmartFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartService_GetSmartFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartFeatureQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartServiceServer).GetSmartFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmartService_GetSmartFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartServiceServer).GetSmartFeature(ctx, req.(*SmartFeatureQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartService_ServiceDesc is the grpc.ServiceDesc for SmartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartservice.SmartService",
	HandlerType: (*SmartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSmartModel",
			Handler:    _SmartService_CreateSmartModel_Handler,
		},
		{
			MethodName: "GetSmartModel",
			Handler:    _SmartService_GetSmartModel_Handler,
		},
		{
			MethodName: "CreateSmartFeature",
			Handler:    _SmartService_CreateSmartFeature_Handler,
		},
		{
			MethodName: "GetSmartFeature",
			Handler:    _SmartService_GetSmartFeature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/smart_service.proto",
}
