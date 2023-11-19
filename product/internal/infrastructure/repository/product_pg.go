package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
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
		&product.Weight,
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

func (repo *ProductRepo) GetProduct(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	rows, err := repo.pg.Pool.Query(ctx, getProductByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getProductByID: %w", err)
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

func (repo *ProductRepo) getAllProductsByFilters(ctx context.Context, query string, filters ...interface{}) ([]*model.Product, error) {
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

func (repo *ProductRepo) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := repo.getAllProductsByFilters(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) GetAllUserProducts(ctx context.Context, userID uuid.UUID) ([]*model.Product, error) {
	products, err := repo.getAllProductsByFilters(ctx, getAllUserProducts, userID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) GetAllCategoryProducts(ctx context.Context, categoryID int32) ([]*model.Product, error) {
	products, err := repo.getAllProductsByFilters(ctx, getAllCategoryProducts, categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) CreateProduct(ctx context.Context, product model.Product) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		createProduct,
		product.ID,
		product.UserID,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Weight,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createProduct: %w", err)
	}

	return nil
}

func (repo *ProductRepo) UpdateProduct(ctx context.Context, product model.Product) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		updateProduct,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Weight,
		product.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec updateProduct: %w", err)
	}

	return nil
}

func (repo *ProductRepo) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteProduct, id)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteProduct: %w", err)
	}

	return nil
}

func (repo *ProductRepo) GetCategory(ctx context.Context, id int32) (*model.Category, error) {
	rows, err := repo.pg.Pool.Query(ctx, getCategoryByID, id)
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
	rows, err := repo.pg.Pool.Query(ctx, getAllCategories)
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
