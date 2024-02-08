// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductAPIClient is the client API for ProductAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductAPIClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	CreateProducts(ctx context.Context, in *CreateProductsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateProducts(ctx context.Context, in *UpdateProductsRequest, opts ...grpc.CallOption) (*UpdateProductsResponse, error)
	UpdateProductState(ctx context.Context, in *UpdateProductStateRequest, opts ...grpc.CallOption) (*UpdateProductStateResponse, error)
}

type productAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProductAPIClient(cc grpc.ClientConnInterface) ProductAPIClient {
	return &productAPIClient{cc}
}

func (c *productAPIClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productAPIClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productAPIClient) CreateProducts(ctx context.Context, in *CreateProductsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/CreateProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productAPIClient) UpdateProducts(ctx context.Context, in *UpdateProductsRequest, opts ...grpc.CallOption) (*UpdateProductsResponse, error) {
	out := new(UpdateProductsResponse)
	err := c.cc.Invoke(ctx, "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productAPIClient) UpdateProductState(ctx context.Context, in *UpdateProductStateRequest, opts ...grpc.CallOption) (*UpdateProductStateResponse, error) {
	out := new(UpdateProductStateResponse)
	err := c.cc.Invoke(ctx, "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProductState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductAPIServer is the server API for ProductAPI service.
// All implementations must embed UnimplementedProductAPIServer
// for forward compatibility
type ProductAPIServer interface {
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	CreateProducts(context.Context, *CreateProductsRequest) (*emptypb.Empty, error)
	UpdateProducts(context.Context, *UpdateProductsRequest) (*UpdateProductsResponse, error)
	UpdateProductState(context.Context, *UpdateProductStateRequest) (*UpdateProductStateResponse, error)
	mustEmbedUnimplementedProductAPIServer()
}

// UnimplementedProductAPIServer must be embedded to have forward compatible implementations.
type UnimplementedProductAPIServer struct {
}

func (UnimplementedProductAPIServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductAPIServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductAPIServer) CreateProducts(context.Context, *CreateProductsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProducts not implemented")
}
func (UnimplementedProductAPIServer) UpdateProducts(context.Context, *UpdateProductsRequest) (*UpdateProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProducts not implemented")
}
func (UnimplementedProductAPIServer) UpdateProductState(context.Context, *UpdateProductStateRequest) (*UpdateProductStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductState not implemented")
}
func (UnimplementedProductAPIServer) mustEmbedUnimplementedProductAPIServer() {}

// UnsafeProductAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductAPIServer will
// result in compilation errors.
type UnsafeProductAPIServer interface {
	mustEmbedUnimplementedProductAPIServer()
}

func RegisterProductAPIServer(s grpc.ServiceRegistrar, srv ProductAPIServer) {
	s.RegisterService(&ProductAPI_ServiceDesc, srv)
}

func _ProductAPI_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAPIServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAPIServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductAPI_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAPIServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAPIServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductAPI_CreateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAPIServer).CreateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/CreateProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAPIServer).CreateProducts(ctx, req.(*CreateProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductAPI_UpdateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAPIServer).UpdateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAPIServer).UpdateProducts(ctx, req.(*UpdateProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductAPI_UpdateProductState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductAPIServer).UpdateProductState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProductState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductAPIServer).UpdateProductState(ctx, req.(*UpdateProductStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductAPI_ServiceDesc is the grpc.ServiceDesc for ProductAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI",
	HandlerType: (*ProductAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductAPI_GetProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductAPI_GetProducts_Handler,
		},
		{
			MethodName: "CreateProducts",
			Handler:    _ProductAPI_CreateProducts_Handler,
		},
		{
			MethodName: "UpdateProducts",
			Handler:    _ProductAPI_UpdateProducts_Handler,
		},
		{
			MethodName: "UpdateProductState",
			Handler:    _ProductAPI_UpdateProductState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.dg.ru/ocb/product-creation/product-storage-api/api/product/product.proto",
}
