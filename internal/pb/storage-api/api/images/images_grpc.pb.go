// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package images

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

// ImageStorageAPIClient is the client API for ImageStorageAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageStorageAPIClient interface {
	ProcessImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*ImageResponse, error)
}

type imageStorageAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewImageStorageAPIClient(cc grpc.ClientConnInterface) ImageStorageAPIClient {
	return &imageStorageAPIClient{cc}
}

func (c *imageStorageAPIClient) ProcessImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/dg.ocb.images.storage_api.pkg.images.ImageStorageAPI/ProcessImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageStorageAPIServer is the server API for ImageStorageAPI service.
// All implementations must embed UnimplementedImageStorageAPIServer
// for forward compatibility
type ImageStorageAPIServer interface {
	ProcessImage(context.Context, *Image) (*ImageResponse, error)
	mustEmbedUnimplementedImageStorageAPIServer()
}

// UnimplementedImageStorageAPIServer must be embedded to have forward compatible implementations.
type UnimplementedImageStorageAPIServer struct {
}

func (UnimplementedImageStorageAPIServer) ProcessImage(context.Context, *Image) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessImage not implemented")
}
func (UnimplementedImageStorageAPIServer) mustEmbedUnimplementedImageStorageAPIServer() {}

// UnsafeImageStorageAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageStorageAPIServer will
// result in compilation errors.
type UnsafeImageStorageAPIServer interface {
	mustEmbedUnimplementedImageStorageAPIServer()
}

func RegisterImageStorageAPIServer(s grpc.ServiceRegistrar, srv ImageStorageAPIServer) {
	s.RegisterService(&ImageStorageAPI_ServiceDesc, srv)
}

func _ImageStorageAPI_ProcessImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageStorageAPIServer).ProcessImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dg.ocb.images.storage_api.pkg.images.ImageStorageAPI/ProcessImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageStorageAPIServer).ProcessImage(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageStorageAPI_ServiceDesc is the grpc.ServiceDesc for ImageStorageAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageStorageAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dg.ocb.images.storage_api.pkg.images.ImageStorageAPI",
	HandlerType: (*ImageStorageAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessImage",
			Handler:    _ImageStorageAPI_ProcessImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.dg.ru/ocb/images/storage-api/api/images/images.proto",
}