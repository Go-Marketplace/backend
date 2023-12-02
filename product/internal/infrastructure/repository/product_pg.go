package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepo struct {
	pg     *postgres.Postgres
	logger *logger.Logger
}

func NewProductRepo(pg *postgres.Postgres, logger *logger.Logger) *ProductRepo {
	return &ProductRepo{
		pg:     pg,
		logger: logger,
	}
}

func scanProduct(rows pgx.Rows, product *model.Product) error {
	return rows.Scan(
		&product.ID,
		&product.UserID,
		&product.CategoryID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.Moderated,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
}

func scanCategory(rows pgx.Rows, category *model.Category) error {
	return rows.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
}

func (repo *ProductRepo) GetProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	query := getProductQuery(productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query from: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query %v get: %w", sqlQuery, err)
	}

	product := &model.Product{}
	found := false
	for rows.Next() {
		err := scanProduct(rows, product)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		found = true
	}

	if !found {
		return nil, nil
	}

	return product, nil
}

func (repo *ProductRepo) getProductsByFilters(ctx context.Context, query string, filters ...interface{}) ([]*model.Product, error) {
	rows, err := repo.pg.Pool.Query(ctx, query, filters...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query %s: %w", query, err)
	}

	products := make([]*model.Product, 0)
	for rows.Next() {
		product := &model.Product{}
		err := scanProduct(rows, product)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductRepo) GetProducts(ctx context.Context, searchParams dto.SearchProductsDTO) ([]*model.Product, error) {
	query := searchProductsQuery(searchParams)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query from getAllProducts: %w", err)
	}

	products, err := repo.getProductsByFilters(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) CreateProduct(ctx context.Context, product model.Product) error {
	query := createProductQuery(product)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec sqlQuery %v: %w", sqlQuery, err)
	}

	return nil
}

func (repo *ProductRepo) UpdateProduct(ctx context.Context, product model.Product) error {
	query := updateProductQuery(product)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateProduct: %w", err)
	}

	return nil
}

func (repo *ProductRepo) UpdateProducts(ctx context.Context, products []model.Product) error {
	batch := &pgx.Batch{}
	for _, product := range products {
		query := updateProductQuery(product)

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("failed to get sql query: %w", err)
		}

		batch.Queue(sqlQuery, args...)
	}

	batchResults := repo.pg.Pool.SendBatch(ctx, batch)
	return batchResults.Close()
}

func (repo *ProductRepo) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	query := deleteProductQuery(productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteProduct: %w", err)
	}

	return nil
}

func (repo *ProductRepo) GetCategory(ctx context.Context, categoryID int32) (*model.Category, error) {
	query := getCategory(categoryID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getCategoryByID: %w", err)
	}

	category := &model.Category{}
	found := false
	for rows.Next() {
		err := scanCategory(rows, category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		found = true
	}

	if !found {
		return nil, nil
	}

	return category, nil
}

func (repo *ProductRepo) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	query := getAllCategoriesQuery()

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getAllCategories: %w", err)
	}

	categories := make([]*model.Category, 0)
	for rows.Next() {
		category := &model.Category{}
		err := scanCategory(rows, category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
