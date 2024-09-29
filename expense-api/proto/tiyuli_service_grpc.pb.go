// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: tiyuli_service.proto

package proto

import (
	context "context"
	expense "github.com/odedro987/tiyuli-server/expense-api/proto/expense"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TiyuliService_NewExpense_FullMethodName = "/tiyuli.TiyuliService/NewExpense"
	TiyuliService_GetExpense_FullMethodName = "/tiyuli.TiyuliService/GetExpense"
)

// TiyuliServiceClient is the client API for TiyuliService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TiyuliServiceClient interface {
	NewExpense(ctx context.Context, in *expense.NewExpenseRequest, opts ...grpc.CallOption) (*expense.NewExpenseResponse, error)
	GetExpense(ctx context.Context, in *expense.GetExpenseByIdRequest, opts ...grpc.CallOption) (*expense.GetExpenseByIdResponse, error)
}

type tiyuliServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTiyuliServiceClient(cc grpc.ClientConnInterface) TiyuliServiceClient {
	return &tiyuliServiceClient{cc}
}

func (c *tiyuliServiceClient) NewExpense(ctx context.Context, in *expense.NewExpenseRequest, opts ...grpc.CallOption) (*expense.NewExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(expense.NewExpenseResponse)
	err := c.cc.Invoke(ctx, TiyuliService_NewExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiyuliServiceClient) GetExpense(ctx context.Context, in *expense.GetExpenseByIdRequest, opts ...grpc.CallOption) (*expense.GetExpenseByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(expense.GetExpenseByIdResponse)
	err := c.cc.Invoke(ctx, TiyuliService_GetExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TiyuliServiceServer is the server API for TiyuliService service.
// All implementations must embed UnimplementedTiyuliServiceServer
// for forward compatibility.
type TiyuliServiceServer interface {
	NewExpense(context.Context, *expense.NewExpenseRequest) (*expense.NewExpenseResponse, error)
	GetExpense(context.Context, *expense.GetExpenseByIdRequest) (*expense.GetExpenseByIdResponse, error)
	mustEmbedUnimplementedTiyuliServiceServer()
}

// UnimplementedTiyuliServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTiyuliServiceServer struct{}

func (UnimplementedTiyuliServiceServer) NewExpense(context.Context, *expense.NewExpenseRequest) (*expense.NewExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewExpense not implemented")
}
func (UnimplementedTiyuliServiceServer) GetExpense(context.Context, *expense.GetExpenseByIdRequest) (*expense.GetExpenseByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpense not implemented")
}
func (UnimplementedTiyuliServiceServer) mustEmbedUnimplementedTiyuliServiceServer() {}
func (UnimplementedTiyuliServiceServer) testEmbeddedByValue()                       {}

// UnsafeTiyuliServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TiyuliServiceServer will
// result in compilation errors.
type UnsafeTiyuliServiceServer interface {
	mustEmbedUnimplementedTiyuliServiceServer()
}

func RegisterTiyuliServiceServer(s grpc.ServiceRegistrar, srv TiyuliServiceServer) {
	// If the following call pancis, it indicates UnimplementedTiyuliServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TiyuliService_ServiceDesc, srv)
}

func _TiyuliService_NewExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(expense.NewExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiyuliServiceServer).NewExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TiyuliService_NewExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiyuliServiceServer).NewExpense(ctx, req.(*expense.NewExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiyuliService_GetExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(expense.GetExpenseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiyuliServiceServer).GetExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TiyuliService_GetExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiyuliServiceServer).GetExpense(ctx, req.(*expense.GetExpenseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TiyuliService_ServiceDesc is the grpc.ServiceDesc for TiyuliService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TiyuliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tiyuli.TiyuliService",
	HandlerType: (*TiyuliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewExpense",
			Handler:    _TiyuliService_NewExpense_Handler,
		},
		{
			MethodName: "GetExpense",
			Handler:    _TiyuliService_GetExpense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tiyuli_service.proto",
}
