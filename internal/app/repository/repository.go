package repository

import (
	"context"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
)

// ProductRepository ...
type ProductRepository interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	GetProducts(ctx context.Context, ids []string) ([]*model.Product, error)
	CreateProducts(ctx context.Context, products []*model.Product, source string, userID int64) error
	UpdateProducts(ctx context.Context, products []*model.Product, source string, userID int64, diffs map[string][]byte) error
	UpdateProductState(ctx context.Context, product *model.Product, source string, userID int64, diff []byte) error
}
