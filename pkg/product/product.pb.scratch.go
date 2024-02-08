// Code generated by protoc-gen-scratch. DO NOT EDIT.
// versions:
// 	protoc-gen-scratch: v0.3.5
// 	protoc:             v3.17.1
// source: gitlab.dg.ru/ocb/product-creation/product-storage-api/api/product/product.proto

//go:generate esc -o swagger.go -pkg product -modtime 0 product.swagger.json
//go:generate scratch ast implement --name="Implementation" --service-name="ProductAPI" --source="gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product" --service-import-path="gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/dg/ocb/product_creation/product_storage_api/pkg/product"
//go:generate scratch ast update-main --service-name="ProductAPI" --source="gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product" --service-import-path="gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/dg/ocb/product_creation/product_storage_api/pkg/product"

package product

import (
	context "context"
	go_grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// ProductAPIServiceDesc is description for the ProductAPIServer.
type ProductAPIServiceDesc struct {
	svc ProductAPIServer
	i   grpc.UnaryServerInterceptor
}

func NewProductAPIServiceDesc(i ProductAPIServer) *ProductAPIServiceDesc {
	return &ProductAPIServiceDesc{svc: i}
}

func (d *ProductAPIServiceDesc) SwaggerDef() []byte {
	return FSMustByte(false, "/product.swagger.json")
}

func (d *ProductAPIServiceDesc) RegisterGRPC(s *grpc.Server) {
	RegisterProductAPIServer(s, d.svc)
}

func (d *ProductAPIServiceDesc) RegisterGateway(ctx context.Context, mux *runtime.ServeMux) error {
	if d.i == nil {
		return RegisterProductAPIHandlerServer(ctx, mux, d.svc)
	}
	return RegisterProductAPIHandlerServer(ctx, mux, &proxyProductAPIServer{
		ProductAPIServer: d.svc,
		interceptor:      d.i,
	})
}

// WithHTTPUnaryInterceptor adds GRPC Server Interceptor for HTTP gateway requests. Call again for multiple Interceptors.
func (d *ProductAPIServiceDesc) WithHTTPUnaryInterceptor(u grpc.UnaryServerInterceptor) {
	if d.i == nil {
		d.i = u
	} else {
		d.i = go_grpc_middleware.ChainUnaryServer(d.i, u)
	}
}

type proxyProductAPIServer struct {
	ProductAPIServer
	interceptor grpc.UnaryServerInterceptor
}

func (p *proxyProductAPIServer) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	info := &grpc.UnaryServerInfo{
		Server:     p.ProductAPIServer,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.ProductAPIServer.GetProduct(ctx, req.(*GetProductRequest))
	}
	resp, err := p.interceptor(ctx, req, info, handler)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.(*GetProductResponse), err
}

func (p *proxyProductAPIServer) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	info := &grpc.UnaryServerInfo{
		Server:     p.ProductAPIServer,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.ProductAPIServer.GetProducts(ctx, req.(*GetProductsRequest))
	}
	resp, err := p.interceptor(ctx, req, info, handler)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.(*GetProductsResponse), err
}

func (p *proxyProductAPIServer) CreateProducts(ctx context.Context, req *CreateProductsRequest) (*emptypb.Empty, error) {
	info := &grpc.UnaryServerInfo{
		Server:     p.ProductAPIServer,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/CreateProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.ProductAPIServer.CreateProducts(ctx, req.(*CreateProductsRequest))
	}
	resp, err := p.interceptor(ctx, req, info, handler)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.(*emptypb.Empty), err
}

func (p *proxyProductAPIServer) UpdateProducts(ctx context.Context, req *UpdateProductsRequest) (*UpdateProductsResponse, error) {
	info := &grpc.UnaryServerInfo{
		Server:     p.ProductAPIServer,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.ProductAPIServer.UpdateProducts(ctx, req.(*UpdateProductsRequest))
	}
	resp, err := p.interceptor(ctx, req, info, handler)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.(*UpdateProductsResponse), err
}

func (p *proxyProductAPIServer) UpdateProductState(ctx context.Context, req *UpdateProductStateRequest) (*UpdateProductStateResponse, error) {
	info := &grpc.UnaryServerInfo{
		Server:     p.ProductAPIServer,
		FullMethod: "/dg.ocb.product_creation.product_storage_api.pkg.product.ProductAPI/UpdateProductState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.ProductAPIServer.UpdateProductState(ctx, req.(*UpdateProductStateRequest))
	}
	resp, err := p.interceptor(ctx, req, info, handler)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.(*UpdateProductStateResponse), err
}
