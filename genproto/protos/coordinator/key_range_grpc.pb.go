// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spqr

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

// KeyRangeServiceClient is the client API for KeyRangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyRangeServiceClient interface {
	ListKeyRange(ctx context.Context, in *ListKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error)
	LockKeyRange(ctx context.Context, in *LockKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error)
	UnlockKeyRange(ctx context.Context, in *UnlockKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error)
	SplitKeyRange(ctx context.Context, in *SplitKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error)
}

type keyRangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyRangeServiceClient(cc grpc.ClientConnInterface) KeyRangeServiceClient {
	return &keyRangeServiceClient{cc}
}

func (c *keyRangeServiceClient) ListKeyRange(ctx context.Context, in *ListKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error) {
	out := new(KeyRangeReply)
	err := c.cc.Invoke(ctx, "/yandex.router.KeyRangeService/ListKeyRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) LockKeyRange(ctx context.Context, in *LockKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error) {
	out := new(KeyRangeReply)
	err := c.cc.Invoke(ctx, "/yandex.router.KeyRangeService/LockKeyRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) UnlockKeyRange(ctx context.Context, in *UnlockKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error) {
	out := new(KeyRangeReply)
	err := c.cc.Invoke(ctx, "/yandex.router.KeyRangeService/UnlockKeyRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) SplitKeyRange(ctx context.Context, in *SplitKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error) {
	out := new(KeyRangeReply)
	err := c.cc.Invoke(ctx, "/yandex.router.KeyRangeService/SplitKeyRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyRangeServiceServer is the server API for KeyRangeService service.
// All implementations must embed UnimplementedKeyRangeServiceServer
// for forward compatibility
type KeyRangeServiceServer interface {
	ListKeyRange(context.Context, *ListKeyRangeRequest) (*KeyRangeReply, error)
	LockKeyRange(context.Context, *LockKeyRangeRequest) (*KeyRangeReply, error)
	UnlockKeyRange(context.Context, *UnlockKeyRangeRequest) (*KeyRangeReply, error)
	SplitKeyRange(context.Context, *SplitKeyRangeRequest) (*KeyRangeReply, error)
	mustEmbedUnimplementedKeyRangeServiceServer()
}

// UnimplementedKeyRangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyRangeServiceServer struct {
}

func (UnimplementedKeyRangeServiceServer) ListKeyRange(context.Context, *ListKeyRangeRequest) (*KeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) LockKeyRange(context.Context, *LockKeyRangeRequest) (*KeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) UnlockKeyRange(context.Context, *UnlockKeyRangeRequest) (*KeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) SplitKeyRange(context.Context, *SplitKeyRangeRequest) (*KeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) mustEmbedUnimplementedKeyRangeServiceServer() {}

// UnsafeKeyRangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyRangeServiceServer will
// result in compilation errors.
type UnsafeKeyRangeServiceServer interface {
	mustEmbedUnimplementedKeyRangeServiceServer()
}

func RegisterKeyRangeServiceServer(s grpc.ServiceRegistrar, srv KeyRangeServiceServer) {
	s.RegisterService(&KeyRangeService_ServiceDesc, srv)
}

func _KeyRangeService_ListKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).ListKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.router.KeyRangeService/ListKeyRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).ListKeyRange(ctx, req.(*ListKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_LockKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).LockKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.router.KeyRangeService/LockKeyRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).LockKeyRange(ctx, req.(*LockKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_UnlockKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).UnlockKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.router.KeyRangeService/UnlockKeyRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).UnlockKeyRange(ctx, req.(*UnlockKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_SplitKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).SplitKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.router.KeyRangeService/SplitKeyRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).SplitKeyRange(ctx, req.(*SplitKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyRangeService_ServiceDesc is the grpc.ServiceDesc for KeyRangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyRangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.router.KeyRangeService",
	HandlerType: (*KeyRangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKeyRange",
			Handler:    _KeyRangeService_ListKeyRange_Handler,
		},
		{
			MethodName: "LockKeyRange",
			Handler:    _KeyRangeService_LockKeyRange_Handler,
		},
		{
			MethodName: "UnlockKeyRange",
			Handler:    _KeyRangeService_UnlockKeyRange_Handler,
		},
		{
			MethodName: "SplitKeyRange",
			Handler:    _KeyRangeService_SplitKeyRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/coordinator/key_range.proto",
}