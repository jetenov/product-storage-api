package storageapi

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/service"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/product-service-meta/api/sc"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images"
	desc "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/scratch"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// Implementation ...
type Implementation struct {
	sc sc.SellerCenterAPIClient
	desc.UnimplementedProductAPIServer
	srv     *service.ProductService
	imgCli  images.ImageStorageAPIClient
	workCnt int
}

// NewProductAPI return new instance of Implementation.
func NewProductAPI(srv *service.ProductService, c sc.SellerCenterAPIClient, imgCli images.ImageStorageAPIClient, workCnt int) *Implementation {
	return &Implementation{srv: srv, sc: c, imgCli: imgCli, workCnt: workCnt}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Implementation) GetDescription() scratch.ServiceDesc {
	return desc.NewProductAPIServiceDesc(i)
}

func (i *Implementation) enrichAttrNames(ctx context.Context, products []*model.Product) ([]*model.Product, error) {
	if len(products) == 0 {
		return nil, nil
	}

	var req sc.GetAttributesNamesMapRequest
	uniq := make(map[int64]struct{})
	for _, p := range products {
		for _, a := range p.Attributes {
			if _, ok := uniq[a.ID]; ok {
				continue
			}

			uniq[a.ID] = struct{}{}
			req.Ids = append(req.Ids, a.ID)
		}
	}

	resp, err := i.sc.GetAttributesNamesMap(ctx, &req)
	if err != nil {
		return nil, err
	}

	names := make(map[int64]string, len(resp.Attributes))
	for _, a := range resp.Attributes {
		names[a.Id] = a.Name
	}

	for _, p := range products {
		for i, a := range p.Attributes {
			name, ok := names[a.ID]
			if !ok {
				continue
			}

			a.Name = name
			p.Attributes[i] = a
		}
	}

	return products, nil
}

func (i *Implementation) proceedImages(ctx context.Context, products []*model.Product) error {
	chunkSize := len(products) / i.workCnt
	if chunkSize == 0 {
		chunkSize = 1
	}

	gr, _ := errgroup.WithContext(ctx)
	for n := 0; n < len(products); n += chunkSize {
		n := n
		end := n + chunkSize

		if end > len(products) {
			end = len(products)
		}

		gr.Go(func() error {
			for _, p := range products[n:end] {
				for _, img := range p.Images {
					if img.URL == "" {
						logger.ErrorKV(ctx, "ceph or origin image does not exist", "product_id", p.ID)
						continue
					}

					resp, err := i.imgCli.ProcessImage(ctx, &images.Image{
						OrigUrl:   img.URL,
						CephUrl:   img.CephURL,
						ProductId: p.ID,
					})
					if err != nil {
						logger.ErrorKV(ctx, "ceph or origin image does not exist", "product_id", p.ID)
						continue
					}

					img.CephURL = resp.Url
				}
			}

			return nil
		})
	}

	if err := gr.Wait(); err != nil {
		return fmt.Errorf("failed to process images: %v", err)
	}

	return nil
}
