// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// CategoryAPIClient is the client API for CategoryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryAPIClient interface {
	GetCategoryTree(ctx context.Context, in *GetCategoryTreeRequest, opts ...grpc.CallOption) (*GetCategoryTreeResponse, error)
}

type categoryAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryAPIClient(cc grpc.ClientConnInterface) CategoryAPIClient {
	return &categoryAPIClient{cc}
}

func (c *categoryAPIClient) GetCategoryTree(ctx context.Context, in *GetCategoryTreeRequest, opts ...grpc.CallOption) (*GetCategoryTreeResponse, error) {
	out := new(GetCategoryTreeResponse)
	err := c.cc.Invoke(ctx, "/category.v1.CategoryAPI/GetCategoryTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryAPIServer is the server API for CategoryAPI service.
// All implementations must embed UnimplementedCategoryAPIServer
// for forward compatibility
type CategoryAPIServer interface {
	GetCategoryTree(context.Context, *GetCategoryTreeRequest) (*GetCategoryTreeResponse, error)
	mustEmbedUnimplementedCategoryAPIServer()
}

// UnimplementedCategoryAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryAPIServer struct {
}

func (UnimplementedCategoryAPIServer) GetCategoryTree(context.Context, *GetCategoryTreeRequest) (*GetCategoryTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryTree not implemented")
}
func (UnimplementedCategoryAPIServer) mustEmbedUnimplementedCategoryAPIServer() {}

// UnsafeCategoryAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryAPIServer will
// result in compilation errors.
type UnsafeCategoryAPIServer interface {
	mustEmbedUnimplementedCategoryAPIServer()
}

func RegisterCategoryAPIServer(s grpc.ServiceRegistrar, srv CategoryAPIServer) {
	s.RegisterService(&CategoryAPI_ServiceDesc, srv)
}

func _CategoryAPI_GetCategoryTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryAPIServer).GetCategoryTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/category.v1.CategoryAPI/GetCategoryTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryAPIServer).GetCategoryTree(ctx, req.(*GetCategoryTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryAPI_ServiceDesc is the grpc.ServiceDesc for CategoryAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "category.v1.CategoryAPI",
	HandlerType: (*CategoryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategoryTree",
			Handler:    _CategoryAPI_GetCategoryTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.dg.ru/marketplace/go/seller-api/api/category/v1/category.proto",
}
