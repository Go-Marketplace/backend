// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: cart.proto

package cart

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
	Cart_GetUserCart_FullMethodName         = "/cart.Cart/GetUserCart"
	Cart_CreateCart_FullMethodName          = "/cart.Cart/CreateCart"
	Cart_DeleteCart_FullMethodName          = "/cart.Cart/DeleteCart"
	Cart_DeleteCartCartlines_FullMethodName = "/cart.Cart/DeleteCartCartlines"
	Cart_CreateCartline_FullMethodName      = "/cart.Cart/CreateCartline"
	Cart_UpdateCartline_FullMethodName      = "/cart.Cart/UpdateCartline"
	Cart_DeleteCartline_FullMethodName      = "/cart.Cart/DeleteCartline"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	GetUserCart(ctx context.Context, in *GetUserCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResponse, error)
	DeleteCartCartlines(ctx context.Context, in *DeleteCartCartlinesRequest, opts ...grpc.CallOption) (*DeleteCartCartlinesResponse, error)
	CreateCartline(ctx context.Context, in *CreateCartlineRequest, opts ...grpc.CallOption) (*CartlineResponse, error)
	UpdateCartline(ctx context.Context, in *UpdateCartlineRequest, opts ...grpc.CallOption) (*CartlineResponse, error)
	DeleteCartline(ctx context.Context, in *DeleteCartlineRequest, opts ...grpc.CallOption) (*DeleteCartlineResponse, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) GetUserCart(ctx context.Context, in *GetUserCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, Cart_GetUserCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, Cart_CreateCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResponse, error) {
	out := new(DeleteCartResponse)
	err := c.cc.Invoke(ctx, Cart_DeleteCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) DeleteCartCartlines(ctx context.Context, in *DeleteCartCartlinesRequest, opts ...grpc.CallOption) (*DeleteCartCartlinesResponse, error) {
	out := new(DeleteCartCartlinesResponse)
	err := c.cc.Invoke(ctx, Cart_DeleteCartCartlines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) CreateCartline(ctx context.Context, in *CreateCartlineRequest, opts ...grpc.CallOption) (*CartlineResponse, error) {
	out := new(CartlineResponse)
	err := c.cc.Invoke(ctx, Cart_CreateCartline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) UpdateCartline(ctx context.Context, in *UpdateCartlineRequest, opts ...grpc.CallOption) (*CartlineResponse, error) {
	out := new(CartlineResponse)
	err := c.cc.Invoke(ctx, Cart_UpdateCartline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) DeleteCartline(ctx context.Context, in *DeleteCartlineRequest, opts ...grpc.CallOption) (*DeleteCartlineResponse, error) {
	out := new(DeleteCartlineResponse)
	err := c.cc.Invoke(ctx, Cart_DeleteCartline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	GetUserCart(context.Context, *GetUserCartRequest) (*CartResponse, error)
	CreateCart(context.Context, *CreateCartRequest) (*CartResponse, error)
	DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResponse, error)
	DeleteCartCartlines(context.Context, *DeleteCartCartlinesRequest) (*DeleteCartCartlinesResponse, error)
	CreateCartline(context.Context, *CreateCartlineRequest) (*CartlineResponse, error)
	UpdateCartline(context.Context, *UpdateCartlineRequest) (*CartlineResponse, error)
	DeleteCartline(context.Context, *DeleteCartlineRequest) (*DeleteCartlineResponse, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) GetUserCart(context.Context, *GetUserCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCart not implemented")
}
func (UnimplementedCartServer) CreateCart(context.Context, *CreateCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServer) DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServer) DeleteCartCartlines(context.Context, *DeleteCartCartlinesRequest) (*DeleteCartCartlinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartCartlines not implemented")
}
func (UnimplementedCartServer) CreateCartline(context.Context, *CreateCartlineRequest) (*CartlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartline not implemented")
}
func (UnimplementedCartServer) UpdateCartline(context.Context, *UpdateCartlineRequest) (*CartlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartline not implemented")
}
func (UnimplementedCartServer) DeleteCartline(context.Context, *DeleteCartlineRequest) (*DeleteCartlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartline not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_GetUserCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetUserCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_GetUserCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetUserCart(ctx, req.(*GetUserCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_CreateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_DeleteCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCart(ctx, req.(*DeleteCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_DeleteCartCartlines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartCartlinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).DeleteCartCartlines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_DeleteCartCartlines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCartCartlines(ctx, req.(*DeleteCartCartlinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_CreateCartline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CreateCartline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_CreateCartline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CreateCartline(ctx, req.(*CreateCartlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_UpdateCartline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).UpdateCartline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_UpdateCartline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).UpdateCartline(ctx, req.(*UpdateCartlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_DeleteCartline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).DeleteCartline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_DeleteCartline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCartline(ctx, req.(*DeleteCartlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCart",
			Handler:    _Cart_GetUserCart_Handler,
		},
		{
			MethodName: "CreateCart",
			Handler:    _Cart_CreateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _Cart_DeleteCart_Handler,
		},
		{
			MethodName: "DeleteCartCartlines",
			Handler:    _Cart_DeleteCartCartlines_Handler,
		},
		{
			MethodName: "CreateCartline",
			Handler:    _Cart_CreateCartline_Handler,
		},
		{
			MethodName: "UpdateCartline",
			Handler:    _Cart_UpdateCartline_Handler,
		},
		{
			MethodName: "DeleteCartline",
			Handler:    _Cart_DeleteCartline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
