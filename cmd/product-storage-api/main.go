package main

import (
	"context"
	"strings"

	product_storage_api "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/dg/ocb/product"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/loader"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/repository"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/service"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/config"
	"gitlab.dg.ru/platform/scratch"
	_ "gitlab.dg.ru/platform/scratch/app/pflag"
	"gitlab.dg.ru/platform/scratch/closer"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

func main() {
	ctx := context.Background()

	a, err := scratch.New()
	if err != nil {
		logger.Fatalf(context.Background(), "can't create app: %s", err)
	}

	sm, err := loader.InitDBShardManager(ctx, nil, a.Healthcheck())
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init shard manager: %v", err)
	}

	loader.InitKafkaHealthcheck(ctx, a.Healthcheck())

	catCl, err := loader.InitCategoriesAPIClient(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init categories api client: %v", err)
	}

	valCl, err := loader.InitValidationAPIClient(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init validation api client: %v", err)
	}

	compCl, err := loader.InitCompetitorAPIClient(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init competitor api client: %v", err)
	}

	imgCli, err := loader.InitImageStorageAPIClient(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init image storage api client: %v", err)
	}

	pr, err := loader.InitKafkaProducer(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init producer: %v", err)
	}

	cl, err := loader.InitSellerCenterAPIClient(ctx, nil)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init seller center api client: %v", err)
	}

	createTopic := config.GetValue(ctx, config.KafkaCreateTopic).String()
	correctTopic := config.GetValue(ctx, config.KafkaCorrectTopic).String()
	fillTopic := config.GetValue(ctx, config.KafkaFillTopic).String()
	metricsTopic := config.GetValue(ctx, config.KafkaMetricsTopic).String()

	repo := repository.NewProduct(sm)

	l := strings.Split(config.GetValue(ctx, config.BrandList).String(), ";")
	brands := make(map[string]struct{}, len(l))
	for _, b := range l {
		brands[b] = struct{}{}
	}

	l = strings.Split(config.GetValue(ctx, config.OgrnList).String(), ";")
	ogrn := make(map[string]struct{}, len(l))
	for _, o := range l {
		ogrn[o] = struct{}{}
	}

	srv := service.NewProductService(repo, pr, catCl, valCl, compCl, createTopic, correctTopic, fillTopic,
		metricsTopic, ogrn, brands)

	workCnt := config.GetValue(ctx, config.ImageWorkersCount).Int()
	impl := product_storage_api.NewProductAPI(srv, cl, imgCli, workCnt)
	if err != nil {
		closer.CloseAll()
		logger.Fatalf(ctx, "init db: %v", err)
	}

	if err := a.Run(impl); err != nil {
		logger.Fatalf(ctx, "can't run app: %s", err)
	}
}
