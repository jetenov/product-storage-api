package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/lib/pq"
	"golang.org/x/sync/errgroup"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/db"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/model"
	"gitlab.dg.ru/platform/database-go/sql/balancer/role"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

const (
	queryGetProduct = `SELECT id, name, description, commercial_category_id, description_category_id, attributes,
			images, barcode, vat, price, depth, weight, height, width, created_at, updated_at, 
			state, block_reason, fulfillment, meta, canonical_ids
		FROM products WHERE id = $1`

	queryCreateProduct = `INSERT INTO products (id, name, description, commercial_category_id, description_category_id,
			attributes, images, barcode, vat, price, meta, depth, weight, height, width, state, fulfillment,
			canonical_ids, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6::jsonb, $7::jsonb, $8, $9, $10, $11::jsonb, $12, $13, $14, $15, $16, $17, $18, $19, $20)
		ON CONFLICT (id) do nothing`

	queryCreateHistory = `INSERT INTO products_history (product_id, fulfillment, status, source, user_id, diff)
		VALUES ($1, $2, $3, $4, $5, $6::jsonb)`

	queryUpdateProduct = `UPDATE products SET name = $1, description = $2, attributes = $3::jsonb, images = $4::jsonb,
			barcode = $5, vat = $6, price = $7, depth = $8, weight = $9, height = $10, width = $11, state = $12,
			fulfillment = $13, updated_at = $14, commercial_category_id = $15
		WHERE id = $16`

	queryUpdateProductState = `UPDATE products set state = $1, block_reason = $2, updated_at = $3
		WHERE id = $4`

	queryGetProducts = `SELECT id, name, description, commercial_category_id, description_category_id, attributes,
			images, barcode, vat, price, meta, depth, weight, height, width, created_at, updated_at, fulfillment, state
		FROM products WHERE id = any($1)`

	queryGetProductsLimit250 = `SELECT id, name, description, commercial_category_id, description_category_id,
			attributes, images, barcode, vat, price, meta, depth, weight, height, width, created_at, updated_at,
			fulfillment, state
		FROM products LIMIT 250`

	queryGetProductsImagesBatch = `SELECT id, images
		FROM products
		WHERE EXISTS(SELECT *
					 FROM jsonb_array_elements(images) images(e)
					 WHERE substr(images.e ->> 'ceph_url', 0, 28) = 'http://prod.s3.ceph.s.o3.ru')
		ORDER BY updated_at
		LIMIT %d`

	queryUpdateImages = `UPDATE products SET images = $1::jsonb, updated_at = now() WHERE id = $2`

	queryUpdateImagesHistory = `INSERT INTO products_history (product_id, fulfillment, status, user_id, diff, source)
		(SELECT product_id, fulfillment, status, 0, $1::jsonb, 'maintenance'
		 FROM products_history
		 WHERE product_id = $2
		 ORDER BY created_at DESC
		 LIMIT 1)`
)

// Product ...
type Product struct {
	sm *db.ShardManager
}

// NewProduct ...
func NewProduct(
	sm *db.ShardManager,
) *Product {
	return &Product{
		sm: sm,
	}
}

// GetProduct ...
func (pr *Product) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	sh, err := pr.sm.ShardByID(id)
	if err != nil {
		return nil, err
	}

	p := model.Product{Meta: make(map[string]interface{})}
	err = sh.Instance.Next(ctx, role.Read).
		QueryRow(ctx, queryGetProduct, id).
		Scan(&p.ID, &p.Name, &p.Description, &p.CommercialCategoryID, &p.DescriptionCategoryID, &p.Attributes,
			&p.Images, &p.Barcode, &p.Vat, &p.Price, &p.Depth, &p.Weight, &p.Height, &p.Width, &p.CreatedAt,
			&p.UpdatedAt, &p.State, &p.BlockReason, &p.Fulfillment, &p.Meta, pq.Array(&p.CanonicalIDs))
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// CreateProducts ...
func (pr *Product) CreateProducts(ctx context.Context, products []*model.Product, source string, userID int64) error {
	shards := make(map[string][]*model.Product)
	for _, p := range products {
		sh, err := pr.sm.ShardByID(p.ID)
		if err != nil {
			return err
		}

		shards[sh.Name] = append(shards[sh.Name], p)
	}

	gr, gCtx := errgroup.WithContext(ctx)
	for name, productsByShard := range shards {
		n := name
		prbs := make([]*model.Product, len(productsByShard))
		copy(prbs, productsByShard)

		gr.Go(func() error {
			return pr.createProducts(gCtx, n, prbs, source, userID)
		})
	}

	return gr.Wait()
}

func (pr *Product) createProducts(ctx context.Context, name string, products []*model.Product, source string, userID int64) error {
	sh, err := pr.sm.ShardByName(name)
	if err != nil {
		return err
	}

	tx, err := sh.Instance.Next(ctx, role.Write).Begin(ctx, nil)
	if err != nil {
		return err
	}

	for _, p := range products {
		res, err := tx.Exec(ctx, queryCreateProduct, p.ID, p.Name, p.Description, p.CommercialCategoryID,
			p.DescriptionCategoryID, p.Attributes, p.Images, p.Barcode, p.Vat, p.Price, p.Meta,
			p.Depth, p.Weight, p.Height, p.Width, p.State, p.Fulfillment, pq.Array(p.CanonicalIDs), p.CreatedAt, p.UpdatedAt)
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "create product rollback: %v", rErr)
			}

			return err
		}

		n, err := res.RowsAffected()
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "rows affected rollback: %v", rErr)
			}

			return err
		}

		if n != 1 {
			logger.WarnKV(ctx, "product already exists", "product_id", p.ID)
			continue
		}

		_, err = tx.Exec(ctx, queryCreateHistory, p.ID, p.Fulfillment, p.State, source, userID, p)
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "create products history rollback: %v", rErr)
			}

			return err
		}
	}

	return tx.Commit()
}

// GetProducts ...
func (pr *Product) GetProducts(ctx context.Context, ids []string) ([]*model.Product, error) {
	if len(ids) == 0 {
		return pr.getProducts250(ctx)
	}

	products := make([]*model.Product, 0, len(ids))
	shards := pr.sm.ShardPerIDs(ctx, ids)
	mu := sync.Mutex{}

	gr, gCtx := errgroup.WithContext(ctx)
	for name, idsByShard := range shards {
		n := name
		ibs := make([]string, len(idsByShard))
		copy(ibs, idsByShard)

		gr.Go(func() error {
			prod, err := pr.getProducts(gCtx, n, ibs)
			if err != nil {
				return err
			}

			mu.Lock()
			products = append(products, prod...)
			mu.Unlock()

			return nil
		})
	}

	if err := gr.Wait(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *Product) getProducts(ctx context.Context, name string, ids []string) ([]*model.Product, error) {
	sh, err := pr.sm.ShardByName(name)
	if err != nil {
		return nil, err
	}

	rows, err := sh.Instance.Next(ctx, role.Read).Query(ctx, queryGetProducts, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("query get products failed: %w", err)
	}

	products := make([]*model.Product, 0, len(ids))
	for rows.Next() {
		p := model.Product{Meta: make(map[string]interface{})}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CommercialCategoryID, &p.DescriptionCategoryID,
			&p.Attributes, &p.Images, &p.Barcode, &p.Vat, &p.Price, &p.Meta,
			&p.Depth, &p.Weight, &p.Height, &p.Width, &p.CreatedAt, &p.UpdatedAt, &p.Fulfillment, &p.State,
		); err != nil {
			return nil, fmt.Errorf("scan for products failed: %w", err)
		}

		products = append(products, &p)
	}

	return products, nil
}

// GetProductsImagesBatch ...
func (pr *Product) GetProductsImagesBatch(ctx context.Context, size int) ([]*model.ProductImages, error) {
	products := make([]*model.ProductImages, 0, size*len(pr.sm.Buckets))
	for _, bucket := range pr.sm.Buckets {
		rows, err := bucket.DB.Instance.Next(ctx, role.Read).Query(ctx, fmt.Sprintf(queryGetProductsImagesBatch, size))
		if err != nil {
			return nil, fmt.Errorf("query get products failed: %w", err)
		}

		for rows.Next() {
			var p model.ProductImages
			if err := rows.Scan(&p.ID, &p.Images); err != nil {
				return nil, fmt.Errorf("scan for products failed: %w", err)
			}

			products = append(products, &p)
		}
	}

	return products, nil
}

func (pr *Product) getProducts250(ctx context.Context) ([]*model.Product, error) {
	products := make([]*model.Product, 0, 1000)

	// TODO: сделать параллельные запросы к каждому шарду
	for _, sh := range pr.sm.Buckets {
		rows, err := sh.DB.Instance.Next(ctx, role.Read).Query(ctx, queryGetProductsLimit250)
		if err != nil {
			return nil, fmt.Errorf("query get products failed: %w", err)
		}

		for rows.Next() {
			p := model.Product{Meta: make(map[string]interface{})}
			err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.CommercialCategoryID, &p.DescriptionCategoryID,
				&p.Attributes, &p.Images, &p.Barcode, &p.Vat, &p.Price, &p.Meta, &p.Depth, &p.Weight, &p.Height,
				&p.Width, &p.CreatedAt, &p.UpdatedAt, &p.Fulfillment, &p.State)
			if err != nil {
				return nil, fmt.Errorf("scan for products failed: %w", err)
			}

			products = append(products, &p)
		}
	}

	return products, nil
}

// UpdateProducts ...
func (pr *Product) UpdateProducts(ctx context.Context, products []*model.Product, source string, userID int64, diffs map[string][]byte) error {
	shards := make(map[string][]*model.Product)
	for _, p := range products {
		sh, err := pr.sm.ShardByID(p.ID)
		if err != nil {
			return err
		}

		shards[sh.Name] = append(shards[sh.Name], p)
	}

	gr, gCtx := errgroup.WithContext(ctx)
	for name, productsByShard := range shards {
		n := name
		prbs := make([]*model.Product, len(productsByShard))
		copy(prbs, productsByShard)

		gr.Go(func() error {
			return pr.updateProducts(gCtx, n, prbs, source, userID, diffs)
		})
	}

	return gr.Wait()
}

func (pr *Product) updateProducts(ctx context.Context, name string, products []*model.Product, source string, userID int64, diffs map[string][]byte) error {
	sh, err := pr.sm.ShardByName(name)
	if err != nil {
		return err
	}

	tx, err := sh.Instance.Next(ctx, role.Write).Begin(ctx, nil)
	if err != nil {
		return err
	}

	for _, p := range products {
		res, err := tx.Exec(ctx, queryUpdateProduct, p.Name, p.Description, p.Attributes, p.Images, p.Barcode, p.Vat,
			p.Price, p.Depth, p.Weight, p.Height, p.Width, p.State, p.Fulfillment, p.UpdatedAt, p.CommercialCategoryID, p.ID)
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "update product rollback: %v", rErr)
			}

			return err
		}

		n, err := res.RowsAffected()
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "rows affected rollback: %v", rErr)
			}

			return err
		}

		if n != 1 {
			logger.ErrorKV(ctx, "product for update not found", "product_id", p.ID)
			continue
		}

		if _, err := tx.Exec(ctx, queryCreateHistory, p.ID, p.Fulfillment, p.State, source, userID, string(diffs[p.ID])); err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "create products history rollback: %v", rErr)
			}

			return err
		}
	}

	return tx.Commit()
}

// UpdateProductsImages ...
func (pr *Product) UpdateProductsImages(ctx context.Context, products []*model.ProductImages, diffs map[string][]byte) error {
	shards := make(map[string][]*model.ProductImages)
	for _, p := range products {
		sh, err := pr.sm.ShardByID(p.ID)
		if err != nil {
			return err
		}

		shards[sh.Name] = append(shards[sh.Name], p)
	}

	gr, gCtx := errgroup.WithContext(ctx)
	for name, productsByShard := range shards {
		n := name
		prbs := make([]*model.ProductImages, len(productsByShard))
		copy(prbs, productsByShard)

		gr.Go(func() error {
			return pr.updateProductsImages(gCtx, n, prbs, diffs)
		})
	}

	return gr.Wait()
}

func (pr *Product) updateProductsImages(ctx context.Context, name string, products []*model.ProductImages, diffs map[string][]byte) error {
	sh, err := pr.sm.ShardByName(name)
	if err != nil {
		return err
	}

	tx, err := sh.Instance.Next(ctx, role.Write).Begin(ctx, nil)
	if err != nil {
		return err
	}

	for _, p := range products {
		res, err := tx.Exec(ctx, queryUpdateImages, p.Images, p.ID)
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "update product rollback: %v", rErr)
			}

			return err
		}

		n, err := res.RowsAffected()
		if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "rows affected rollback: %v", rErr)
			}

			return err
		}

		if n != 1 {
			logger.ErrorKV(ctx, "product for update not found", "product_id", p.ID)
			continue
		}

		if _, err := tx.Exec(ctx, queryUpdateImagesHistory, string(diffs[p.ID]), p.ID); err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logger.Errorf(ctx, "create products history rollback: %v", rErr)
			}

			return err
		}
	}

	return tx.Commit()
}

// UpdateProductState ...
func (pr *Product) UpdateProductState(ctx context.Context, p *model.Product, source string, userID int64, diff []byte) error {
	sh, err := pr.sm.ShardByID(p.ID)
	if err != nil {
		return err
	}

	tx, err := sh.Instance.Next(ctx, role.Write).Begin(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.Exec(ctx, queryUpdateProductState, p.State, p.BlockReason, p.UpdatedAt, p.ID)
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			logger.Errorf(ctx, "update product state rollback: %v", rErr)
		}

		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			logger.Errorf(ctx, "rows affected rollback: %v", rErr)
		}

		return err
	}

	if n != 1 {
		if rErr := tx.Rollback(); rErr != nil {
			logger.Errorf(ctx, "affected rollback: %v", rErr)
		}

		return errors.New("product for update not found")
	}

	if _, err = tx.Exec(ctx, queryCreateHistory, p.ID, p.Fulfillment, p.State, source, userID, string(diff)); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			logger.Errorf(ctx, "create products history rollback: %v", rErr)
		}

		return err
	}

	return tx.Commit()
}
