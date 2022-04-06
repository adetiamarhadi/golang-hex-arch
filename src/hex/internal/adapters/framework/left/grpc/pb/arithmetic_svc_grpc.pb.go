// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	GetAdd(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error)
	GetSubtract(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error)
	GetMultiply(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error)
	GetDivide(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) GetAdd(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetSubtract(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetSubtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetMultiply(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetMultiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetDivide(ctx context.Context, in *OperationParameter, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetDivide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations must embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	GetAdd(context.Context, *OperationParameter) (*Answer, error)
	GetSubtract(context.Context, *OperationParameter) (*Answer, error)
	GetMultiply(context.Context, *OperationParameter) (*Answer, error)
	GetDivide(context.Context, *OperationParameter) (*Answer, error)
	mustEmbedUnimplementedArithmeticServiceServer()
}

// UnimplementedArithmeticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) GetAdd(context.Context, *OperationParameter) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdd not implemented")
}
func (UnimplementedArithmeticServiceServer) GetSubtract(context.Context, *OperationParameter) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubtract not implemented")
}
func (UnimplementedArithmeticServiceServer) GetMultiply(context.Context, *OperationParameter) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiply not implemented")
}
func (UnimplementedArithmeticServiceServer) GetDivide(context.Context, *OperationParameter) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivide not implemented")
}
func (UnimplementedArithmeticServiceServer) mustEmbedUnimplementedArithmeticServiceServer() {}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_GetAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetAdd(ctx, req.(*OperationParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetSubtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetSubtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetSubtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetSubtract(ctx, req.(*OperationParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetMultiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetMultiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetMultiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetMultiply(ctx, req.(*OperationParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetDivide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetDivide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetDivide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetDivide(ctx, req.(*OperationParameter))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdd",
			Handler:    _ArithmeticService_GetAdd_Handler,
		},
		{
			MethodName: "GetSubtract",
			Handler:    _ArithmeticService_GetSubtract_Handler,
		},
		{
			MethodName: "GetMultiply",
			Handler:    _ArithmeticService_GetMultiply_Handler,
		},
		{
			MethodName: "GetDivide",
			Handler:    _ArithmeticService_GetDivide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_svc.proto",
}
