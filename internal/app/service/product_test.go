package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/producer"
	mocksProducer "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/producer/mocks"
	psapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/database-go/sql"
)

func TestProductService_sendMetricsMessage(t *testing.T) {
	type fields struct {
		pr           producer.KafkaProducer
		metricsTopic string
	}
	type args struct {
		ctx     context.Context
		product *model.Product
		userID  int64
		source  string
	}

	var (
		pr           = &mocksProducer.KafkaProducer{}
		metricsTopic = "ocb_product_metrics"
		ctx          = context.Background()
	)

	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeTest func(t *testing.T)
		want       error
		wantErr    bool
	}{
		{
			name: "Test with empty product",
			fields: fields{
				pr:           pr,
				metricsTopic: metricsTopic,
			},
			args: args{ctx: ctx, userID: 1, product: &model.Product{}, source: "testSource"},
			beforeTest: func(t *testing.T) {
				pr.On("SendMessage", ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test with nil product",
			fields: fields{
				pr:           pr,
				metricsTopic: metricsTopic,
			},
			args:    args{ctx: ctx, userID: 1, product: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with product",
			fields: fields{
				pr:           pr,
				metricsTopic: metricsTopic,
			},
			args: args{ctx: ctx, userID: 1, product: &model.Product{
				ID:                   "ID",
				Name:                 sql.NullString{Valid: true, String: "Name"},
				Fulfillment:          sql.NullInt64{Valid: true, Int64: 100},
				CommercialCategoryID: 111,
				State:                sql.NullString{Valid: true, String: "state"},
				Barcode:              sql.NullString{Valid: true, String: "barcode"},
			}},
			beforeTest: func(t *testing.T) {
				pr.On("SendMessage", ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ps := &ProductService{
				producer:     pr,
				metricsTopic: tt.fields.metricsTopic,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(t)
			}

			if err := ps.sendMetricsMessage(tt.args.ctx, tt.args.product, tt.args.userID, tt.args.source); (err != nil) != tt.wantErr {
				t.Errorf("sendMetricsMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getProductToMetrics(t *testing.T) {
	type args struct {
		ctx     context.Context
		product *model.Product
		userID  int64
		source  string
	}
	var (
		ctx       = context.Background()
		createdAt = time.Now()
		updatedAt = createdAt.Add(time.Hour)
		Brand     = &model.Brand{Title: "nike"}
		Sellers   = []*model.Seller{{Title: "seller", OGRN: "ogrn"}}
		Meta      = make(map[string]interface{})
		sellers   = []*psapi.Seller{{Title: "seller", Ogrn: "ogrn"}}
	)

	tests := []struct {
		name string
		args args
		want *psapi.MetricsProductUpdate
	}{
		{
			name: "Test nil product",
			args: args{ctx: ctx, product: nil, userID: 1, source: "testSource"},
			want: nil,
		},
		{
			name: "simple test",
			args: args{ctx: ctx, userID: 1, source: "testSource", product: &model.Product{
				ID:                   "ID",
				Name:                 sql.NullString{Valid: true, String: "Name"},
				Fulfillment:          sql.NullInt64{Valid: true, Int64: 100},
				CommercialCategoryID: 111,
				State:                sql.NullString{Valid: true, String: "state"},
				Barcode:              sql.NullString{Valid: true, String: "barcode"},
				CreatedAt:            createdAt,
				UpdatedAt:            updatedAt,
				Meta:                 Meta,
			}},
			want: &psapi.MetricsProductUpdate{
				Id:                   "ID",
				Name:                 "Name",
				Fulfillment:          100,
				CommercialCategoryId: 111,
				Barcode:              "barcode",
				State:                "state",
				Brand:                "nike",
				Sellers:              &psapi.Sellers{Sellers: sellers},
				CreatedAt:            timestamppb.New(createdAt),
				UpdatedAt:            timestamppb.New(updatedAt),
				UserIds:              []int64{1},
				LastUpdateSource:     "testSource",
			},
		},
	}
	for _, tt := range tests {
		Meta[model.BrandKey] = Brand
		Meta[model.SellersKey] = Sellers
		t.Run(tt.name, func(t *testing.T) {
			if got := getProductToMetrics(tt.args.ctx, tt.args.product, tt.args.userID, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProductToMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
