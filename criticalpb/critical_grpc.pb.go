// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package criticalpb

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

// CriticalSectionGRPCClient is the client API for CriticalSectionGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CriticalSectionGRPCClient interface {
	GetIdFromServer(ctx context.Context, in *Message, opts ...grpc.CallOption) (*IdResponse, error)
	RequestAccessToCritical(ctx context.Context, in *Message, opts ...grpc.CallOption) (*AccessGranted, error)
	RetriveCriticalInformation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ReleaseAccessToCritical(ctx context.Context, in *Message, opts ...grpc.CallOption) (*AccessReleased, error)
}

type criticalSectionGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewCriticalSectionGRPCClient(cc grpc.ClientConnInterface) CriticalSectionGRPCClient {
	return &criticalSectionGRPCClient{cc}
}

func (c *criticalSectionGRPCClient) GetIdFromServer(ctx context.Context, in *Message, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/criticalpb.CriticalSectionGRPC/GetIdFromServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *criticalSectionGRPCClient) RequestAccessToCritical(ctx context.Context, in *Message, opts ...grpc.CallOption) (*AccessGranted, error) {
	out := new(AccessGranted)
	err := c.cc.Invoke(ctx, "/criticalpb.CriticalSectionGRPC/RequestAccessToCritical", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *criticalSectionGRPCClient) RetriveCriticalInformation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/criticalpb.CriticalSectionGRPC/RetriveCriticalInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *criticalSectionGRPCClient) ReleaseAccessToCritical(ctx context.Context, in *Message, opts ...grpc.CallOption) (*AccessReleased, error) {
	out := new(AccessReleased)
	err := c.cc.Invoke(ctx, "/criticalpb.CriticalSectionGRPC/ReleaseAccessToCritical", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CriticalSectionGRPCServer is the server API for CriticalSectionGRPC service.
// All implementations must embed UnimplementedCriticalSectionGRPCServer
// for forward compatibility
type CriticalSectionGRPCServer interface {
	GetIdFromServer(context.Context, *Message) (*IdResponse, error)
	RequestAccessToCritical(context.Context, *Message) (*AccessGranted, error)
	RetriveCriticalInformation(context.Context, *Message) (*Message, error)
	ReleaseAccessToCritical(context.Context, *Message) (*AccessReleased, error)
	mustEmbedUnimplementedCriticalSectionGRPCServer()
}

// UnimplementedCriticalSectionGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedCriticalSectionGRPCServer struct {
}

func (UnimplementedCriticalSectionGRPCServer) GetIdFromServer(context.Context, *Message) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdFromServer not implemented")
}
func (UnimplementedCriticalSectionGRPCServer) RequestAccessToCritical(context.Context, *Message) (*AccessGranted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccessToCritical not implemented")
}
func (UnimplementedCriticalSectionGRPCServer) RetriveCriticalInformation(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetriveCriticalInformation not implemented")
}
func (UnimplementedCriticalSectionGRPCServer) ReleaseAccessToCritical(context.Context, *Message) (*AccessReleased, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseAccessToCritical not implemented")
}
func (UnimplementedCriticalSectionGRPCServer) mustEmbedUnimplementedCriticalSectionGRPCServer() {}

// UnsafeCriticalSectionGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CriticalSectionGRPCServer will
// result in compilation errors.
type UnsafeCriticalSectionGRPCServer interface {
	mustEmbedUnimplementedCriticalSectionGRPCServer()
}

func RegisterCriticalSectionGRPCServer(s grpc.ServiceRegistrar, srv CriticalSectionGRPCServer) {
	s.RegisterService(&CriticalSectionGRPC_ServiceDesc, srv)
}

func _CriticalSectionGRPC_GetIdFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CriticalSectionGRPCServer).GetIdFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/criticalpb.CriticalSectionGRPC/GetIdFromServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CriticalSectionGRPCServer).GetIdFromServer(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _CriticalSectionGRPC_RequestAccessToCritical_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CriticalSectionGRPCServer).RequestAccessToCritical(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/criticalpb.CriticalSectionGRPC/RequestAccessToCritical",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CriticalSectionGRPCServer).RequestAccessToCritical(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _CriticalSectionGRPC_RetriveCriticalInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CriticalSectionGRPCServer).RetriveCriticalInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/criticalpb.CriticalSectionGRPC/RetriveCriticalInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CriticalSectionGRPCServer).RetriveCriticalInformation(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _CriticalSectionGRPC_ReleaseAccessToCritical_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CriticalSectionGRPCServer).ReleaseAccessToCritical(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/criticalpb.CriticalSectionGRPC/ReleaseAccessToCritical",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CriticalSectionGRPCServer).ReleaseAccessToCritical(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// CriticalSectionGRPC_ServiceDesc is the grpc.ServiceDesc for CriticalSectionGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CriticalSectionGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "criticalpb.CriticalSectionGRPC",
	HandlerType: (*CriticalSectionGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdFromServer",
			Handler:    _CriticalSectionGRPC_GetIdFromServer_Handler,
		},
		{
			MethodName: "RequestAccessToCritical",
			Handler:    _CriticalSectionGRPC_RequestAccessToCritical_Handler,
		},
		{
			MethodName: "RetriveCriticalInformation",
			Handler:    _CriticalSectionGRPC_RetriveCriticalInformation_Handler,
		},
		{
			MethodName: "ReleaseAccessToCritical",
			Handler:    _CriticalSectionGRPC_ReleaseAccessToCritical_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "criticalpb/critical.proto",
}