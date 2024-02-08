package storageapi

import (
	"context"

	"google.golang.org/grpc/status"

	xo3ctx "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pkg/context"
	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
)

// UpdateProductState stub
func (i *Implementation) UpdateProductState(ctx context.Context, req *desc.UpdateProductStateRequest) (*desc.UpdateProductStateResponse, error) {
	if err := i.srv.UpdateProductState(ctx, req.ProductId, req.State, req.BlockReason, xo3ctx.GetAppName(ctx), xo3ctx.GetUserID(ctx)); err != nil {
		return nil, status.Errorf(21, "UpdateProductState: %v", err)
	}

	return &desc.UpdateProductStateResponse{Success: true}, nil
}
