// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: product.proto

package product

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
	Product_GetProduct_FullMethodName             = "/product.Product/GetProduct"
	Product_GetAllProducts_FullMethodName         = "/product.Product/GetAllProducts"
	Product_GetAllUserProducts_FullMethodName     = "/product.Product/GetAllUserProducts"
	Product_GetAllCategoryProducts_FullMethodName = "/product.Product/GetAllCategoryProducts"
	Product_CreateProduct_FullMethodName          = "/product.Product/CreateProduct"
	Product_UpdateProduct_FullMethodName          = "/product.Product/UpdateProduct"
	Product_ModerateProduct_FullMethodName        = "/product.Product/ModerateProduct"
	Product_DeleteProduct_FullMethodName          = "/product.Product/DeleteProduct"
	Product_GetCategory_FullMethodName            = "/product.Product/GetCategory"
	Product_GetAllCategories_FullMethodName       = "/product.Product/GetAllCategories"
	Product_CreateDiscount_FullMethodName         = "/product.Product/CreateDiscount"
	Product_DeleteDiscount_FullMethodName         = "/product.Product/DeleteDiscount"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	GetAllUserProducts(ctx context.Context, in *GetAllUserProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	GetAllCategoryProducts(ctx context.Context, in *GetAllCategoryProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	ModerateProduct(ctx context.Context, in *ModerateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
	CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteDiscount(ctx context.Context, in *DeleteDiscountRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, Product_GetAllProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetAllUserProducts(ctx context.Context, in *GetAllUserProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, Product_GetAllUserProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetAllCategoryProducts(ctx context.Context, in *GetAllCategoryProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, Product_GetAllCategoryProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ModerateProduct(ctx context.Context, in *ModerateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_ModerateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, Product_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, Product_GetCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, Product_GetAllCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_CreateDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteDiscount(ctx context.Context, in *DeleteDiscountRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, Product_DeleteDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	GetProduct(context.Context, *GetProductRequest) (*ProductResponse, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*ProductsResponse, error)
	GetAllUserProducts(context.Context, *GetAllUserProductsRequest) (*ProductsResponse, error)
	GetAllCategoryProducts(context.Context, *GetAllCategoryProductsRequest) (*ProductsResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*ProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error)
	ModerateProduct(context.Context, *ModerateProductRequest) (*ProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	GetCategory(context.Context, *GetCategoryRequest) (*CategoryResponse, error)
	GetAllCategories(context.Context, *GetAllCategoriesRequest) (*CategoriesResponse, error)
	CreateDiscount(context.Context, *CreateDiscountRequest) (*ProductResponse, error)
	DeleteDiscount(context.Context, *DeleteDiscountRequest) (*ProductResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) GetProduct(context.Context, *GetProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedProductServer) GetAllUserProducts(context.Context, *GetAllUserProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserProducts not implemented")
}
func (UnimplementedProductServer) GetAllCategoryProducts(context.Context, *GetAllCategoryProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategoryProducts not implemented")
}
func (UnimplementedProductServer) CreateProduct(context.Context, *CreateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServer) UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServer) ModerateProduct(context.Context, *ModerateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModerateProduct not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) GetCategory(context.Context, *GetCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedProductServer) GetAllCategories(context.Context, *GetAllCategoriesRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedProductServer) CreateDiscount(context.Context, *CreateDiscountRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscount not implemented")
}
func (UnimplementedProductServer) DeleteDiscount(context.Context, *DeleteDiscountRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscount not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetAllUserProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetAllUserProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetAllUserProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetAllUserProducts(ctx, req.(*GetAllUserProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetAllCategoryProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoryProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetAllCategoryProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetAllCategoryProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetAllCategoryProducts(ctx, req.(*GetAllCategoryProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ModerateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModerateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ModerateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ModerateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ModerateProduct(ctx, req.(*ModerateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetAllCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetAllCategories(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateDiscount(ctx, req.(*CreateDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteDiscount(ctx, req.(*DeleteDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _Product_GetAllProducts_Handler,
		},
		{
			MethodName: "GetAllUserProducts",
			Handler:    _Product_GetAllUserProducts_Handler,
		},
		{
			MethodName: "GetAllCategoryProducts",
			Handler:    _Product_GetAllCategoryProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Product_UpdateProduct_Handler,
		},
		{
			MethodName: "ModerateProduct",
			Handler:    _Product_ModerateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _Product_GetCategory_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _Product_GetAllCategories_Handler,
		},
		{
			MethodName: "CreateDiscount",
			Handler:    _Product_CreateDiscount_Handler,
		},
		{
			MethodName: "DeleteDiscount",
			Handler:    _Product_DeleteDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
