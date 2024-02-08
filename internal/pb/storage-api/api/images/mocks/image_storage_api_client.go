// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
	images "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images"
)

// ImageStorageAPIClient is an autogenerated mock type for the ImageStorageAPIClient type
type ImageStorageAPIClient struct {
	mock.Mock
}

// ProcessImage provides a mock function with given fields: ctx, in, opts
func (_m *ImageStorageAPIClient) ProcessImage(ctx context.Context, in *images.Image, opts ...grpc.CallOption) (*images.ImageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *images.ImageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *images.Image, ...grpc.CallOption) *images.ImageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*images.ImageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *images.Image, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}