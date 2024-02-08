package loader

import (
	"context"

	"google.golang.org/grpc"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/config"
	catapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/api/api/categories"
	compapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/ml-data-consumer/api/ml-data-consumer"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/product-service-meta/api/sc"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images"
	valapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/validation-service/api/validation"
	"gitlab.dg.ru/platform/scratch/closer"
	mw "gitlab.dg.ru/platform/scratch/pkg/mw/grpc"
)

// InitValidationAPIClient ...
func InitValidationAPIClient(ctx context.Context, cl *closer.Closer) (valapi.ValidateAPIClient, error) {
	conn, err := mw.DialContext(
		ctx,
		config.GetValue(ctx, config.ValidationApiGateway).String(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(conn.Close)
	} else {
		cl.Add(conn.Close)
	}

	return valapi.NewValidateAPIClient(conn), nil
}

// InitCategoriesAPIClient ...
func InitCategoriesAPIClient(ctx context.Context, cl *closer.Closer) (catapi.CategoryAPIClient, error) {
	conn, err := mw.DialContext(
		ctx,
		config.GetValue(ctx, config.CategoriesApiGateway).String(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(conn.Close)
	} else {
		cl.Add(conn.Close)
	}

	return catapi.NewCategoryAPIClient(conn), nil
}

// InitCompetitorAPIClient ...
func InitCompetitorAPIClient(ctx context.Context, cl *closer.Closer) (compapi.CompetitorAPIClient, error) {
	conn, err := mw.DialContext(
		ctx,
		config.GetValue(ctx, config.CompetitorApiGateway).String(),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(conn.Close)
	} else {
		cl.Add(conn.Close)
	}

	return compapi.NewCompetitorAPIClient(conn), nil
}

// InitSellerCenterAPIClient ...
func InitSellerCenterAPIClient(ctx context.Context, cl *closer.Closer) (sc.SellerCenterAPIClient, error) {
	conn, err := mw.DialContext(
		ctx,
		config.GetValue(ctx, config.ProductServiceMetaGateway).String(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(conn.Close)
	} else {
		cl.Add(conn.Close)
	}

	return sc.NewSellerCenterAPIClient(conn), nil
}

// InitImageStorageAPIClient inits grpc client to ocb-image-storage-api service
func InitImageStorageAPIClient(ctx context.Context, cl *closer.Closer) (images.ImageStorageAPIClient, error) {
	conn, err := mw.DialContext(
		ctx,
		config.GetValue(ctx, config.ImageStorageApiGateway).String(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(conn.Close)
	} else {
		cl.Add(conn.Close)
	}

	return images.NewImageStorageAPIClient(conn), nil
}
