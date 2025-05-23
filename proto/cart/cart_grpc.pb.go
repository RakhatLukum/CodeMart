// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: cart.proto

package cartpb

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
	CartService_AddToCart_FullMethodName              = "/cart.CartService/AddToCart"
	CartService_RemoveFromCart_FullMethodName         = "/cart.CartService/RemoveFromCart"
	CartService_GetCart_FullMethodName                = "/cart.CartService/GetCart"
	CartService_ClearCart_FullMethodName              = "/cart.CartService/ClearCart"
	CartService_CartItemCount_FullMethodName          = "/cart.CartService/CartItemCount"
	CartService_HasProduct_FullMethodName             = "/cart.CartService/HasProduct"
	CartService_ReplaceCart_FullMethodName            = "/cart.CartService/ReplaceCart"
	CartService_GetCartTotal_FullMethodName           = "/cart.CartService/GetCartTotal"
	CartService_GetCartProducts_FullMethodName        = "/cart.CartService/GetCartProducts"
	CartService_AddMultipleToCart_FullMethodName      = "/cart.CartService/AddMultipleToCart"
	CartService_RemoveMultipleFromCart_FullMethodName = "/cart.CartService/RemoveMultipleFromCart"
	CartService_GetCartSummary_FullMethodName         = "/cart.CartService/GetCartSummary"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	AddToCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartResponse, error)
	RemoveFromCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartResponse, error)
	GetCart(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartList, error)
	ClearCart(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartResponse, error)
	CartItemCount(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartCountResponse, error)
	HasProduct(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartHasResponse, error)
	ReplaceCart(ctx context.Context, in *ReplaceCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	GetCartTotal(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartTotalResponse, error)
	GetCartProducts(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartProductList, error)
	AddMultipleToCart(ctx context.Context, in *AddMultipleRequest, opts ...grpc.CallOption) (*CartResponse, error)
	RemoveMultipleFromCart(ctx context.Context, in *RemoveMultipleRequest, opts ...grpc.CallOption) (*CartResponse, error)
	GetCartSummary(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartSummaryResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddToCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_AddToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveFromCart(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_RemoveFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartList)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ClearCart(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_ClearCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CartItemCount(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartCountResponse)
	err := c.cc.Invoke(ctx, CartService_CartItemCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) HasProduct(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*CartHasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartHasResponse)
	err := c.cc.Invoke(ctx, CartService_HasProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ReplaceCart(ctx context.Context, in *ReplaceCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_ReplaceCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartTotal(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartTotalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartTotalResponse)
	err := c.cc.Invoke(ctx, CartService_GetCartTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartProducts(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartProductList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartProductList)
	err := c.cc.Invoke(ctx, CartService_GetCartProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddMultipleToCart(ctx context.Context, in *AddMultipleRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_AddMultipleToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveMultipleFromCart(ctx context.Context, in *RemoveMultipleRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_RemoveMultipleFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartSummary(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*CartSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartSummaryResponse)
	err := c.cc.Invoke(ctx, CartService_GetCartSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility.
type CartServiceServer interface {
	AddToCart(context.Context, *CartItemRequest) (*CartResponse, error)
	RemoveFromCart(context.Context, *CartItemRequest) (*CartResponse, error)
	GetCart(context.Context, *UserIdRequest) (*CartList, error)
	ClearCart(context.Context, *UserIdRequest) (*CartResponse, error)
	CartItemCount(context.Context, *UserIdRequest) (*CartCountResponse, error)
	HasProduct(context.Context, *CartItemRequest) (*CartHasResponse, error)
	ReplaceCart(context.Context, *ReplaceCartRequest) (*CartResponse, error)
	GetCartTotal(context.Context, *UserIdRequest) (*CartTotalResponse, error)
	GetCartProducts(context.Context, *UserIdRequest) (*CartProductList, error)
	AddMultipleToCart(context.Context, *AddMultipleRequest) (*CartResponse, error)
	RemoveMultipleFromCart(context.Context, *RemoveMultipleRequest) (*CartResponse, error)
	GetCartSummary(context.Context, *UserIdRequest) (*CartSummaryResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartServiceServer struct{}

func (UnimplementedCartServiceServer) AddToCart(context.Context, *CartItemRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveFromCart(context.Context, *CartItemRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *UserIdRequest) (*CartList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) ClearCart(context.Context, *UserIdRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServiceServer) CartItemCount(context.Context, *UserIdRequest) (*CartCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemCount not implemented")
}
func (UnimplementedCartServiceServer) HasProduct(context.Context, *CartItemRequest) (*CartHasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasProduct not implemented")
}
func (UnimplementedCartServiceServer) ReplaceCart(context.Context, *ReplaceCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceCart not implemented")
}
func (UnimplementedCartServiceServer) GetCartTotal(context.Context, *UserIdRequest) (*CartTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartTotal not implemented")
}
func (UnimplementedCartServiceServer) GetCartProducts(context.Context, *UserIdRequest) (*CartProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartProducts not implemented")
}
func (UnimplementedCartServiceServer) AddMultipleToCart(context.Context, *AddMultipleRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMultipleToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveMultipleFromCart(context.Context, *RemoveMultipleRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMultipleFromCart not implemented")
}
func (UnimplementedCartServiceServer) GetCartSummary(context.Context, *UserIdRequest) (*CartSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartSummary not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}
func (UnimplementedCartServiceServer) testEmbeddedByValue()                     {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	// If the following call pancis, it indicates UnimplementedCartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddToCart(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_RemoveFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveFromCart(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ClearCart(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CartItemCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CartItemCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_CartItemCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CartItemCount(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_HasProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).HasProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_HasProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).HasProduct(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ReplaceCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ReplaceCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ReplaceCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ReplaceCart(ctx, req.(*ReplaceCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCartTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartTotal(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCartProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartProducts(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddMultipleToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddMultipleToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddMultipleToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddMultipleToCart(ctx, req.(*AddMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveMultipleFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveMultipleFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_RemoveMultipleFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveMultipleFromCart(ctx, req.(*RemoveMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCartSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartSummary(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToCart",
			Handler:    _CartService_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _CartService_RemoveFromCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _CartService_ClearCart_Handler,
		},
		{
			MethodName: "CartItemCount",
			Handler:    _CartService_CartItemCount_Handler,
		},
		{
			MethodName: "HasProduct",
			Handler:    _CartService_HasProduct_Handler,
		},
		{
			MethodName: "ReplaceCart",
			Handler:    _CartService_ReplaceCart_Handler,
		},
		{
			MethodName: "GetCartTotal",
			Handler:    _CartService_GetCartTotal_Handler,
		},
		{
			MethodName: "GetCartProducts",
			Handler:    _CartService_GetCartProducts_Handler,
		},
		{
			MethodName: "AddMultipleToCart",
			Handler:    _CartService_AddMultipleToCart_Handler,
		},
		{
			MethodName: "RemoveMultipleFromCart",
			Handler:    _CartService_RemoveMultipleFromCart_Handler,
		},
		{
			MethodName: "GetCartSummary",
			Handler:    _CartService_GetCartSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
