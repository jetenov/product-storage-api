package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/status"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/loader"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/repository"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/config"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/storage-api/api/images"
	_ "gitlab.dg.ru/platform/scratch/app/pflag"
	"gitlab.dg.ru/platform/scratch/closer"
	"gitlab.dg.ru/platform/scratch/pkg/opts"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

func init() {
	_ = opts.Set(opts.Options{
		AppName:    config.AppName,
		AppVersion: "1.0",
	})
}

const (
	shardsCnt = 4

	msgImageNotFound = "process image: original image is unavailable: unexpected http code: 404"
)

func main() {
	appCloser := closer.New(syscall.SIGTERM, syscall.SIGINT)
	defer appCloser.CloseAll()

	ctx := context.Background()

	sm, err := loader.InitDBShardManager(ctx, appCloser, nil)
	if err != nil {
		appCloser.CloseAll()
		logger.Fatalf(ctx, "init shard manager: %v", err)
	}

	imcl, err := loader.InitImageStorageAPIClient(ctx, appCloser)
	if err != nil {
		appCloser.CloseAll()
		logger.Fatalf(ctx, "init image storage api client: %v", err)
	}

	workCnt := config.GetValue(ctx, config.ImageWorkersCount).Int()
	batchSize := config.GetValue(ctx, config.BatchSize).Int()
	saveBatchSize := config.GetValue(ctx, config.SaveBatchSize).Int()
	pref := config.GetValue(ctx, config.CephImagePrefix).String()

	repo := repository.NewProduct(sm)
	products, err := repo.GetProductsImagesBatch(ctx, batchSize)
	if err != nil {
		appCloser.CloseAll()
		logger.Fatalf(ctx, "failed to get products: %v", err)
	}

	reduce := make(chan *model.ProductImages, batchSize*shardsCnt)
	var rwg sync.WaitGroup
	rwg.Add(1)
	go func() {
		defer rwg.Done()

		batch := make([]*model.ProductImages, 0, saveBatchSize)
		diffs := make(map[string][]byte, batchSize*shardsCnt)
		for p := range reduce {
			batch = append(batch, p)
			js, _ := json.Marshal(struct {
				Images model.Images `json:"images"`
			}{Images: p.Images})
			diffs[p.ID] = js

			if len(batch) == saveBatchSize {
				if err := repo.UpdateProductsImages(ctx, batch, diffs); err != nil {
					logger.Errorf(ctx, "failed to save products images: %v", err)
				}

				batch = batch[:0]
			}
		}

		if len(batch) > 0 {
			if err := repo.UpdateProductsImages(ctx, batch, diffs); err != nil {
				logger.Errorf(ctx, "failed to save products images: %v", err)
			}
		}
	}()

	chunkSize := batchSize * shardsCnt / workCnt
	ch := make(chan struct{}, len(products)/chunkSize+1)
	go func() {
		appCloser.Wait()
		for i := 0; i < len(products)/chunkSize+1; i++ {
			ch <- struct{}{}
		}
	}()
	defer func() { close(ch) }()

	gr, _ := errgroup.WithContext(ctx)
	for n := 0; n < len(products); n += chunkSize {
		n := n
		end := n + chunkSize

		if end > len(products) {
			end = len(products)
		}

		gr.Go(func() error {
			for _, p := range products[n:end] {
				select {
				case <-ch:
					return nil
				default:
				}

				var isMainRemoved, isImageChanged bool
				var removed int
				for i, img := range p.Images {
					if img.URL == "" {
						if img.IsMain {
							isMainRemoved = true
						}
						removed++

						logger.ErrorKV(ctx, "original image does not exist", "product_id", p.ID)
						continue
					}

					if strings.HasPrefix(img.CephURL, pref) {
						p.Images[i-removed] = img
						continue
					}

					resp, err := imcl.ProcessImage(ctx, &images.Image{
						OrigUrl:   img.URL,
						CephUrl:   img.CephURL,
						ProductId: p.ID,
					})
					if err != nil {
						if st, ok := status.FromError(err); ok && st.Message() == msgImageNotFound {
							if img.IsMain {
								isMainRemoved = true
							}
							removed++
						} else {
							p.Images[i-removed] = img
						}

						logger.ErrorKV(ctx, fmt.Sprintf("failed to process image: %v", err), "product_id", p.ID)
						continue
					}

					isImageChanged = true
					img.CephURL = resp.Url
					p.Images[i-removed] = img
				}
				p.Images = p.Images[:len(p.Images)-removed]
				if isMainRemoved && len(p.Images) > 1 {
					p.Images[0].IsMain = true
				}

				if removed > 0 || isImageChanged {
					reduce <- p
				}
			}

			return nil
		})
	}

	if err := gr.Wait(); err != nil {
		logger.Errorf(ctx, "failed to process images: %v", err)
	}

	close(reduce)
	rwg.Wait()
}
