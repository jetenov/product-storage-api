package storageapi

import (
	"context"
	"fmt"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	competitor "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/ml-data-consumer/api/ml-data-consumer"
	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// GetProduct stub.
func (i *Implementation) GetProduct(ctx context.Context, req *desc.GetProductRequest) (*desc.GetProductResponse, error) {
	product, err := i.srv.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, status.Errorf(21, "GetProduct: %v", err)
	}

	products, err := i.enrichAttrNames(ctx, []*model.Product{product})
	if err != nil {
		logger.Errorf(ctx, "seller center client: enrich attribute names: %w", err)
	} else if len(products) == 1 {
		product = products[0]
	}

	if err = i.proceedImages(ctx, []*model.Product{product}); err != nil {
		return nil, fmt.Errorf("proceed images: %v", err)
	}

	competitors, err := i.srv.GetCompetitors(ctx, product.CanonicalIDs)
	if err != nil {
		return nil, status.Errorf(21, "GetCompetitors: %v", err)
	}

	return &desc.GetProductResponse{Product: convertProductToResponse(ctx, product, competitors)}, nil
}

func convertProductToResponse(ctx context.Context, p *model.Product, competitors []*competitor.CompetitorProduct) *desc.Product {
	meta, err := p.Meta.ToProto()
	if err != nil {
		logger.ErrorKV(ctx, fmt.Errorf("failed marshal meta to proto: %w", err).Error(), "product_id", p.ID)
	}

	return &desc.Product{
		Id:                    p.ID,
		Attributes:            p.AttributesToProto(),
		Images:                p.ImagesToProto(),
		Name:                  p.Name.String,
		Barcode:               p.Barcode.String,
		Vat:                   p.Vat.Int64,
		Price:                 p.Price.Int64,
		Description:           p.Description.String,
		CommercialCategoryId:  p.CommercialCategoryID,
		DescriptionCategoryId: p.DescriptionCategoryID,
		Meta:                  meta,
		Depth:                 p.Depth.Int64,
		Weight:                p.Weight.Int64,
		Width:                 p.Width.Int64,
		Height:                p.Height.Int64,
		State:                 p.State.String,
		Fulfillment:           p.Fulfillment.Int64,
		CreatedAt:             timestamppb.New(p.CreatedAt),
		UpdatedAt:             timestamppb.New(p.UpdatedAt),
		CompetitorProducts:    competitors,
	}
}
