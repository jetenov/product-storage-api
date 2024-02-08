package storageapi

import (
	"context"
	"time"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	xo3ctx "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pkg/context"
	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/database-go/sql"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// CreateProducts stub
func (i *Implementation) CreateProducts(ctx context.Context, req *desc.CreateProductsRequest) (*emptypb.Empty, error) {
	if len(req.Products) == 0 {
		return nil, status.Errorf(3, "Invalid argument")
	}

	products := convertProductsFromPB(ctx, req.Products)
	if err := i.proceedImages(ctx, products); err != nil {
		return nil, status.Errorf(21, "CreateProducts: %v", err)
	}

	err := i.srv.CreateProducts(ctx, products, xo3ctx.GetAppName(ctx), xo3ctx.GetUserID(ctx))
	if err != nil {
		return nil, status.Errorf(21, "CreateProducts: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func convertProductsFromPB(ctx context.Context, products []*desc.ProductForCreate) []*model.Product {
	responseProducts := make([]*model.Product, 0, len(products))

	for _, p := range products {
		attributes := make([]*model.Attribute, 0, len(p.Attributes))
		for _, attribute := range p.Attributes {
			attr := model.Attribute{
				ID:             attribute.Id,
				ComplexID:      attribute.ComplexId,
				Name:           attribute.Name,
				AttributeValue: make([]model.Value, 0, len(attribute.Values)),
			}
			for _, v := range attribute.Values {
				attr.AttributeValue = append(attr.AttributeValue, model.Value{ID: v.Id, Value: v.Value})
			}

			attributes = append(attributes, &attr)
		}

		images := make([]*model.Image, 0, len(p.Images))
		for _, image := range p.Images {
			images = append(images, &model.Image{URL: image.Url, CephURL: image.CephUrl, IsMain: image.IsMain})
		}

		lp := &model.Product{
			ID:         p.Id,
			Attributes: attributes,
			Images:     images,
			Name: sql.NullString{
				String: p.Name,
				Valid:  p.Name != "",
			},
			Barcode: sql.NullString{
				String: p.Barcode,
				Valid:  p.Barcode != "",
			},
			Vat: sql.NullInt64{
				Int64: p.Vat,
				Valid: p.Vat != 0,
			},
			Price: sql.NullInt64{
				Int64: p.Price,
				Valid: p.Price != 0,
			},
			Description: sql.NullString{
				String: p.Description,
				Valid:  p.Description != "",
			},
			CommercialCategoryID:  p.CommercialCategoryId,
			DescriptionCategoryID: p.DescriptionCategoryId,
			Depth: sql.NullInt64{
				Int64: p.Depth,
				Valid: p.Depth != 0,
			},
			Weight: sql.NullInt64{
				Int64: p.Weight,
				Valid: p.Weight != 0,
			},
			Width: sql.NullInt64{
				Int64: p.Width,
				Valid: p.Width != 0,
			},
			Height: sql.NullInt64{
				Int64: p.Height,
				Valid: p.Height != 0,
			},
			CanonicalIDs: p.CanonicalIds,
			Meta:         make(map[string]interface{}),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := lp.Meta.Fill(p.Meta); err != nil {
			logger.ErrorKV(ctx, err.Error(), "product_id", lp.ID)
		}

		responseProducts = append(responseProducts, lp)
	}

	return responseProducts
}
