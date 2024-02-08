package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/producer"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/repository"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/api/api/categories"
	competitor "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/ml-data-consumer/api/ml-data-consumer"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/validation-service/api/validation"
	psapi "gitlab.dg.ru/ocb/product-creation/product-storage-api/pkg/product"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// App sources
const (
	TaskAPISource   = "ocb-tasks-api"
	CorrectorSource = "ocb-autocorrection-service"

	priceRequiredValidationMsg = "Цена должна быть заполнена"

	blockReasonLimitAttempts           = 2
	blockReasonCategoryMismatch        = 5
	blockReasonDifferentDictionaryData = 6
)

// ProductRepository ...
type ProductRepository interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	GetProducts(ctx context.Context, ids []string) ([]*model.Product, error)
	CreateProducts(ctx context.Context, products []*model.Product, source string, userID int64) error
	UpdateProducts(ctx context.Context, products []*model.Product, source string, userID int64, diffs map[string][]byte) error
	UpdateProductState(ctx context.Context, product *model.Product, source string, userID int64, diff []byte) error
}

// ProductService ...
type ProductService struct {
	repo         repository.ProductRepository
	producer     producer.KafkaProducer
	catCli       categories.CategoryAPIClient
	valCli       validation.ValidateAPIClient
	mlCLi        competitor.CompetitorAPIClient
	createTopic  string
	correctTopic string
	fillTopic    string
	metricsTopic string
	ogrn         map[string]struct{}
	brands       map[string]struct{}
}

// NewProductService ...
func NewProductService(
	pr repository.ProductRepository,
	producer producer.KafkaProducer,
	catCli categories.CategoryAPIClient,
	valCli validation.ValidateAPIClient,
	mlCli competitor.CompetitorAPIClient,
	createTopic string,
	correctTopic string,
	fillTopic string,
	metricsTopic string,
	ogrn map[string]struct{},
	brands map[string]struct{},
) *ProductService {
	return &ProductService{
		repo:         pr,
		producer:     producer,
		catCli:       catCli,
		valCli:       valCli,
		mlCLi:        mlCli,
		createTopic:  createTopic,
		correctTopic: correctTopic,
		fillTopic:    fillTopic,
		ogrn:         ogrn,
		brands:       brands,
		metricsTopic: metricsTopic,
	}
}

// GetProduct ...
func (ps *ProductService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	return ps.repo.GetProduct(ctx, id)
}

// GetProducts ...
func (ps *ProductService) GetProducts(ctx context.Context, ids []string) ([]*model.Product, error) {
	return ps.repo.GetProducts(ctx, ids)
}

// CreateProducts ...
func (ps *ProductService) CreateProducts(ctx context.Context, products []*model.Product, source string, userID int64) error {
	validationResults, err := ps.getValidationResults(ctx, products)
	if err != nil {
		return err
	}

	productsForCreate := make([]*model.Product, 0, len(products))
	for _, p := range products {
		result, ok := validationResults[p.ID]
		if !ok {
			logger.ErrorKV(ctx, "validation result not found", "product_id", p.ID)
			continue
		}

		p.Fulfillment = sql.NullInt64{Int64: result.Fulfillment, Valid: result.Fulfillment != 0}
		p.State = sql.NullString{String: model.CorrectionStatus, Valid: true}
		if result.Fulfillment == 100 {
			p.State = sql.NullString{String: model.CreatedStatus, Valid: true}
		}

		productsForCreate = append(productsForCreate, p)
	}

	if err := ps.repo.CreateProducts(ctx, productsForCreate, source, userID); err != nil {
		return err
	}

	for _, p := range productsForCreate {
		var errorList []*psapi.Error
		for _, e := range validationResults[p.ID].Errors {
			errorList = append(errorList, &psapi.Error{Attribute: e.Attribute, Description: e.Description})
		}

		if err := ps.sendMessage(ctx, p, errorList); err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("produce product: %v", err), "product_id", p.ID)
		}

		if err := ps.sendMetricsMessage(ctx, p, userID, source); err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("produce metric product: %v", err), "product_id", p.ID)
		}
	}

	return nil
}

// UpdateProducts ...
func (ps *ProductService) UpdateProducts(ctx context.Context, products []*model.Product, source string, userID int64) (*model.ErrorList, error) {
	currentProducts, err := ps.getCurrentProducts(ctx, products)
	if err != nil {
		return nil, err
	}

	if len(currentProducts) == 0 {
		logger.Warnf(ctx, "update products: products for update not found")
		return nil, nil
	}

	for _, p := range products {
		cur, ok := currentProducts[p.ID]
		if !ok {
			continue
		}

		// TODO: обогатить всеми необходимыми полями
		if p.CommercialCategoryID == 0 {
			p.CommercialCategoryID = cur.CommercialCategoryID
		}
		p.Meta = cur.Meta
	}

	gr, gCtx := errgroup.WithContext(ctx)

	validationResults := make(map[string]*validation.Result, len(products))
	gr.Go(func() error {
		validationResults, err = ps.getValidationResults(gCtx, products)

		return err
	})

	catParameters := make(map[int64]*categories.CategoryParameters)
	gr.Go(func() error {
		catParameters = ps.getCategoriesParameters(ctx, getUniqCategories(products))

		return nil
	})

	if err = gr.Wait(); err != nil {
		return nil, err
	}

	productsForUpdate := make([]*model.Product, 0, len(products))
	diffs := make(map[string][]byte, len(products))
	for _, p := range products {
		val, ok := validationResults[p.ID]
		if !ok {
			logger.ErrorKV(ctx, "validation result not found", "product_id", p.ID)
			continue
		}

		// TODO: выпилить это в отдельный метод
		if len(products) == 1 && len(val.Errors) > 0 {
			errorList := model.NewErrorList()
			for _, e := range val.Errors {
				if e.Description == priceRequiredValidationMsg {
					continue
				}

				errorList.Add(e.Attribute, e.Description)
			}

			if len(errorList.List) > 0 {
				return errorList, nil
			}
		}

		cur, ok := currentProducts[p.ID]
		if !ok {
			logger.ErrorKV(ctx, "old product not found", "product_id", p.ID)
			continue
		}

		cat, ok := catParameters[p.CommercialCategoryID]
		if !ok {
			logger.ErrorKV(ctx, "commercial category not found", "product_id", p.ID, "category_id",
				p.CommercialCategoryID)
			continue
		}

		p.Fulfillment = sql.NullInt64{Int64: val.Fulfillment, Valid: val.Fulfillment != 0}
		p.MinFulfillment = sql.NullInt64{Int64: cat.MinFulfillment.Value, Valid: true}
		p.State = sql.NullString{String: ps.getProductState(ctx, source, p), Valid: true}
		p.RootCategoryID = cat.RootCategoryId

		diffs[p.ID], err = p.MergePatch(cur)
		if err != nil {
			logger.ErrorKV(ctx, "failed to create diff", "product_id", p.ID)
			continue
		}

		productsForUpdate = append(productsForUpdate, p)
	}

	if err := ps.repo.UpdateProducts(ctx, productsForUpdate, source, userID, diffs); err != nil {
		return nil, err
	}

	for _, p := range productsForUpdate {
		if err := ps.sendMessage(ctx, p, nil); err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("produce product: %v", err), "product_id", p.ID)
		}

		if err := ps.sendMetricsMessage(ctx, p, userID, source); err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("produce metric product: %v", err), "product_id", p.ID)
		}
	}

	return nil, nil
}

// UpdateProductState ...
func (ps *ProductService) UpdateProductState(ctx context.Context, id, state string, reason int64, source string, userID int64) error {
	cur, err := ps.repo.GetProduct(ctx, id)
	if err != nil {
		return err
	}

	if source != TaskAPISource {
		return fmt.Errorf("incorrect source: %s", source)
	}

	if !(cur.State.String == model.CorrectionStatus || cur.State.String == model.FillingStatus) {
		return fmt.Errorf("incorrect product status: %s", cur.State.String)
	}

	switch reason {
	case blockReasonLimitAttempts, blockReasonCategoryMismatch, blockReasonDifferentDictionaryData:
	default:
		return fmt.Errorf("incorrect block reason: %d", reason)
	}

	tmp := *cur
	upd := &tmp
	upd.State = sql.NullString{String: state, Valid: true}
	upd.BlockReason = sql.NullInt64{Int64: reason, Valid: true}
	upd.UpdatedAt = time.Now()

	diff, err := upd.MergePatch(cur)
	if err != nil {
		return err
	}

	if err := ps.sendMetricsMessage(ctx, upd, userID, source); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("produce metrics product: %v", err), "product_id", upd.ID)
	}

	return ps.repo.UpdateProductState(ctx, upd, source, userID, diff)
}

func (ps *ProductService) validateProducts(ctx context.Context, products []*model.Product) ([]*validation.Result, error) {
	var list []*validation.Product

	for _, product := range products {
		list = append(list, &validation.Product{
			Id:                   product.ID,
			Name:                 product.Name.String,
			Description:          product.Description.String,
			Barcode:              product.Barcode.String,
			Vat:                  product.Vat.Int64,
			Price:                product.Price.Int64,
			Depth:                product.Depth.Int64,
			Weight:               product.Weight.Int64,
			Height:               product.Height.Int64,
			Width:                product.Width.Int64,
			CommercialCategoryId: product.CommercialCategoryID,
			Images:               product.ImagesToValidation(),
			Attributes:           product.AttributesToValidation(),
		})
	}

	resp, err := ps.valCli.ValidateProducts(ctx, &validation.ValidateProductsRequest{Products: list})
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// GetCompetitors ...
func (ps *ProductService) GetCompetitors(ctx context.Context, ids []string) ([]*competitor.CompetitorProduct, error) {
	resp, err := ps.mlCLi.GetCompetitorProducts(ctx, &competitor.GetCompetitorProductsRequest{CanonicalIds: ids})
	if err != nil {
		return nil, err
	}

	return resp.CompetitorProducts, nil
}

func (ps *ProductService) getCurrentProducts(ctx context.Context, products []*model.Product) (map[string]*model.Product, error) {
	ids := make([]string, 0, len(products))
	for _, p := range products {
		ids = append(ids, p.ID)
	}

	pr, err := ps.GetProducts(ctx, ids)
	if err != nil {
		return nil, err
	}

	cur := make(map[string]*model.Product, len(pr))
	for _, p := range pr {
		cur[p.ID] = p
	}

	return cur, nil
}

func (ps *ProductService) getValidationResults(ctx context.Context, products []*model.Product) (map[string]*validation.Result, error) {
	resp, err := ps.validateProducts(ctx, products)
	if err != nil {
		return nil, err
	}

	vr := make(map[string]*validation.Result, len(resp))
	for _, result := range resp {
		vr[result.ProductId] = result
	}

	return vr, nil
}

func (ps *ProductService) getCategoriesParameters(ctx context.Context, ids []int64) map[int64]*categories.CategoryParameters {
	result := make(map[int64]*categories.CategoryParameters, len(ids))
	for _, id := range ids {
		resp, err := ps.catCli.GetCategoriesParameters(ctx, &categories.GetCategoriesParametersRequest{CategoryId: id})
		if err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("failed to get category params: %v", err), "category_id", id)
			continue
		}

		result[id] = resp.Items
	}

	return result
}

func (ps *ProductService) sendMessage(ctx context.Context, product *model.Product, errs []*psapi.Error) error {
	switch product.State.String {
	case model.CreatedStatus:
		pr := getProductToCreate(product)

		msg, err := proto.Marshal(pr)
		if err != nil {
			return err
		}

		return ps.producer.SendMessage(ctx, ps.createTopic, pr.Id, string(msg))
	case model.CorrectionStatus:
		pr := getProductToCorrect(product, errs)

		msg, err := proto.Marshal(pr)
		if err != nil {
			return err
		}

		return ps.producer.SendMessage(ctx, ps.correctTopic, pr.Id, string(msg))
	case model.FillingStatus:
		pr := getProductToFill(product)

		msg, err := proto.Marshal(pr)
		if err != nil {
			return err
		}

		return ps.producer.SendMessage(ctx, ps.fillTopic, pr.Id, string(msg))
	default:
		return fmt.Errorf("wrong product status: %s", product.State.String)
	}
}

func (ps *ProductService) sendMetricsMessage(ctx context.Context, product *model.Product, userID int64, source string) error {
	if product == nil {
		return fmt.Errorf("nil product")
	}

	pr := getProductToMetrics(ctx, product, userID, source)
	metricsMessage := &psapi.MetricsMessage{MessageType: &psapi.MetricsMessage_ProductUpdate{ProductUpdate: pr}}

	msg, err := proto.Marshal(metricsMessage)
	if err != nil {
		return err
	}

	return ps.producer.SendMessage(ctx, ps.metricsTopic, pr.Id, string(msg))
}

func getUniqCategories(products []*model.Product) []int64 {
	var res []int64
	uniqCats := make(map[int64]struct{})
	for _, p := range products {
		if _, ok := uniqCats[p.CommercialCategoryID]; ok {
			continue
		}

		uniqCats[p.CommercialCategoryID] = struct{}{}
		res = append(res, p.CommercialCategoryID)
	}

	return res
}

func (ps *ProductService) getProductState(ctx context.Context, source string, p *model.Product) string {
	if source == CorrectorSource {
		if p.Fulfillment.Int64 == 100 {
			return model.CreatedStatus
		}

		if p.Fulfillment.Int64 > p.MinFulfillment.Int64 {
			if ps.checkOGRN(ctx, p) || ps.checkBrand(ctx, p) {
				return model.FillingStatus
			}
		}

		return model.BlockedStatus
	}

	return model.CreatedStatus
}

func (ps *ProductService) checkOGRN(ctx context.Context, p *model.Product) bool {
	sellers, err := p.Meta.Sellers()
	if err != nil {
		logger.ErrorKV(ctx, err.Error(), "product_id", p.ID)
		return false
	}

	for _, s := range sellers {
		if _, ok := ps.ogrn[s.OGRN]; ok {
			return true
		}
	}

	return false
}

func (ps *ProductService) checkBrand(ctx context.Context, p *model.Product) bool {
	brand, err := p.Meta.Brand()
	if err != nil {
		logger.ErrorKV(ctx, err.Error(), "product_id", p.ID)
		return false
	}

	if _, ok := ps.brands[brand.Title]; !ok {
		return false
	}

	return true
}

func getProductToCreate(product *model.Product) *psapi.ProductToCreate {
	return &psapi.ProductToCreate{
		Id:                    product.ID,
		Attributes:            product.AttributesToProto(),
		Images:                product.ImagesToProto(),
		Name:                  product.Name.String,
		Barcode:               product.Barcode.String,
		Vat:                   product.Vat.Int64,
		Price:                 product.Price.Int64,
		Description:           product.Description.String,
		CommercialCategoryId:  product.CommercialCategoryID,
		DescriptionCategoryId: product.DescriptionCategoryID,
		Depth:                 product.Depth.Int64,
		Weight:                product.Weight.Int64,
		Width:                 product.Width.Int64,
		Height:                product.Height.Int64,
	}
}

func getProductToCorrect(product *model.Product, errorList []*psapi.Error) *psapi.ProductToCorrect {
	return &psapi.ProductToCorrect{
		Id:                    product.ID,
		Attributes:            product.AttributesToProto(),
		Images:                product.ImagesToProto(),
		Name:                  product.Name.String,
		Barcode:               product.Barcode.String,
		Vat:                   product.Vat.Int64,
		Price:                 product.Price.Int64,
		Description:           product.Description.String,
		CommercialCategoryId:  product.CommercialCategoryID,
		DescriptionCategoryId: product.DescriptionCategoryID,
		Depth:                 product.Depth.Int64,
		Weight:                product.Weight.Int64,
		Width:                 product.Width.Int64,
		Height:                product.Height.Int64,
		Errors:                errorList,
	}
}

func getProductToFill(product *model.Product) *psapi.ProductToFill {
	return &psapi.ProductToFill{
		Id:                 product.ID,
		Fulfillment:        product.Fulfillment.Int64,
		PriorityCategoryId: product.CommercialCategoryID,
		RootCategoryId:     product.RootCategoryID,
	}
}

func getProductToMetrics(ctx context.Context, product *model.Product, userID int64, source string) *psapi.MetricsProductUpdate {
	if product == nil {
		return nil
	}

	brand, err := product.Meta.Brand()
	if err != nil {
		logger.ErrorKV(ctx, err.Error(), "product_id", product.ID)
	}

	var title string
	if brand != nil {
		title = brand.Title
	}

	sellersMeta, err := product.Meta.Sellers()
	if err != nil {
		logger.ErrorKV(ctx, err.Error(), "product_id", product.ID)
	}

	sellers := make([]*psapi.Seller, 0, len(sellersMeta))
	for _, s := range sellersMeta {
		sellers = append(sellers, &psapi.Seller{Title: s.Title, Ogrn: s.OGRN})
	}

	return &psapi.MetricsProductUpdate{
		Id:                   product.ID,
		Name:                 product.Name.String,
		Fulfillment:          product.Fulfillment.Int64,
		CommercialCategoryId: product.CommercialCategoryID,
		Barcode:              product.Barcode.String,
		State:                product.State.String,
		Brand:                title,
		BlockReason:          product.BlockReason.Int64,
		Sellers:              &psapi.Sellers{Sellers: sellers},
		CreatedAt:            timestamppb.New(product.CreatedAt),
		UpdatedAt:            timestamppb.New(product.UpdatedAt),
		UserIds:              []int64{userID},
		LastUpdateSource:     source,
	}
}
