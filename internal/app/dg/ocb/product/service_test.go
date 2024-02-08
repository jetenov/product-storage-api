package storageapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/service"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/product-service-meta/api/sc"
	scMocks "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/product-service-meta/api/sc/mocks"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images"
	imgMocks "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images/mocks"
)

func TestImplementation_proceedImages(t *testing.T) {
	var (
		mockSCAPIClient = &scMocks.SellerCenterAPIClient{}
		imgClient       = &imgMocks.ImageStorageAPIClient{}

		ctx = context.Background()
	)

	type fields struct {
		sc     sc.SellerCenterAPIClient
		srv    *service.ProductService
		imgCli images.ImageStorageAPIClient
	}
	type args struct {
		ctx      context.Context
		products []*model.Product
	}
	tests := []struct {
		name       string
		fields     fields
		beforeTest func(t *testing.T)
		args       args
		wantErr    bool
		afterTest  func(t *testing.T, products []*model.Product)
	}{
		{
			name: "Try to process single product",
			fields: fields{
				sc:     mockSCAPIClient,
				srv:    nil,
				imgCli: imgClient,
			},
			beforeTest: func(t *testing.T) {
				imgClient.
					On("ProcessImage", mock.Anything, mock.Anything).
					Return(&images.ImageResponse{
						Url:    "http://ceph2",
						Exists: true,
					}, nil).Once()
			},
			args: args{
				ctx: ctx,
				products: []*model.Product{
					{
						ID: "123",
						Images: model.Images{
							{
								CephURL: "https://ceph",
								URL:     "https://yarr.ru",
								IsMain:  false,
							},
						},
					},
				},
			},
			wantErr: false,
			afterTest: func(t *testing.T, products []*model.Product) {
				assert.NotEmpty(t, products)
				assert.Equal(t, products[0].Images[0].CephURL, "http://ceph2")
			},
		},
		{
			name: "Try to process products",
			fields: fields{
				sc:     mockSCAPIClient,
				srv:    nil,
				imgCli: imgClient,
			},
			beforeTest: func(t *testing.T) {
				imgClient.
					On("ProcessImage", mock.Anything, mock.Anything).
					Return(&images.ImageResponse{
						Url:    "http://ceph2",
						Exists: true,
					}, nil).Twice()
			},
			args: args{
				ctx: ctx,
				products: []*model.Product{
					{
						ID: "123",
						Images: model.Images{
							{
								CephURL: "https://ceph",
								URL:     "https://yarr.ru",
								IsMain:  false,
							},
						},
					},
					{
						ID: "456",
						Images: model.Images{
							{
								CephURL: "https://ceph22",
								URL:     "https://yarr.ru",
								IsMain:  false,
							},
						},
					},
				},
			},
			wantErr: false,
			afterTest: func(t *testing.T, products []*model.Product) {
				assert.NotEmpty(t, products)
				assert.Equal(t, products[0].Images[0].CephURL, "http://ceph2")
				assert.Equal(t, products[1].Images[0].CephURL, "http://ceph2")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Implementation{
				sc:      tt.fields.sc,
				srv:     tt.fields.srv,
				imgCli:  tt.fields.imgCli,
				workCnt: 10,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(t)
			}
			if err := i.proceedImages(tt.args.ctx, tt.args.products); (err != nil) != tt.wantErr {
				t.Errorf("proceedImages() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.afterTest != nil {
				tt.afterTest(t, tt.args.products)
			}
		})
	}
}
