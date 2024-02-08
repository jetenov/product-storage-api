package storageapi

import (
	"context"
	"time"

	"google.golang.org/grpc/status"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	xo3ctx "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pkg/context"
	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/database-go/sql"
)

// UpdateProducts stub.
func (i *Implementation) UpdateProducts(ctx context.Context, req *desc.UpdateProductsRequest) (*desc.UpdateProductsResponse, error) {
	if req.Products == nil {
		return nil, status.Errorf(3, "Invalid argument")
	}

	errorList, err := i.srv.UpdateProducts(ctx, convertUpdProductsFromPB(req.Products), xo3ctx.GetAppName(ctx), xo3ctx.GetUserID(ctx))
	if err != nil {
		return nil, status.Errorf(21, "UpdateProduct: %v", err)
	}

	return &desc.UpdateProductsResponse{Success: true, Errors: i.ConvertErrorListToResp(errorList)}, nil
}

func convertUpdProductsFromPB(products []*desc.ProductForUpdate) []*model.Product {
	responseProduct := make([]*model.Product, 0, len(products))

	for _, p := range products {
		var attributes []*model.Attribute
		for _, attribute := range p.Attributes {
			attr := model.Attribute{ID: attribute.Id, Name: attribute.Name, ComplexID: attribute.ComplexId}
			for _, v := range attribute.Values {
				attr.AttributeValue = append(attr.AttributeValue, model.Value{ID: v.Id, Value: v.Value})
			}
			attributes = append(attributes, &attr)
		}

		var images []*model.Image
		for _, image := range p.Images {
			images = append(images, &model.Image{URL: image.Url, CephURL: image.CephUrl, IsMain: image.IsMain})
		}

		responseProduct = append(responseProduct, &model.Product{
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
			UpdatedAt:            time.Now(),
			CommercialCategoryID: p.CommercialCategoryId,
		})
	}

	return responseProduct
}

// ConvertErrorListToResp ...
func (i *Implementation) ConvertErrorListToResp(list *model.ErrorList) []*desc.Error {
	if list == nil {
		return nil
	}

	errorList := make([]*desc.Error, 0, len(list.List))
	for _, r := range list.List {
		errorList = append(errorList, &desc.Error{Attribute: r.Attribute, Description: r.Description})
	}

	return errorList
}
