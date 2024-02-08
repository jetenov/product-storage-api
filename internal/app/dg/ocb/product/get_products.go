package storageapi

import (
	"context"
	"fmt"

	"google.golang.org/grpc/status"

	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// GetProducts stub. Please implement it.
func (i *Implementation) GetProducts(ctx context.Context, req *desc.GetProductsRequest) (*desc.GetProductsResponse, error) {
	products, err := i.srv.GetProducts(ctx, req.ProductIds)
	if err != nil {
		return nil, status.Errorf(21, "GetProductsList: %v", err)
	}

	products, err = i.enrichAttrNames(ctx, products)
	if err != nil {
		logger.Errorf(ctx, "seller center client: enrich attribute names: %w", err)
	}

	if err = i.proceedImages(ctx, products); err != nil {
		return nil, fmt.Errorf("proceed images: %v", err)
	}

	var productList []*desc.Product
	for _, p := range products {
		productList = append(productList, convertProductToResponse(ctx, p, nil))
	}

	return &desc.GetProductsResponse{Products: productList}, nil
}
