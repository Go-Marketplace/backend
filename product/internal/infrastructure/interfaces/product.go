package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
)

type ProductRepo interface {
	GetProduct(ctx context.Context, id uuid.UUID) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]*model.Product, error)
	GetAllUserProducts(ctx context.Context, userID uuid.UUID) ([]*model.Product, error)
	GetAllCategoryProducts(ctx context.Context, categoryID int32) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ModerateProduct(ctx context.Context, product model.Product) error

	GetAllCategories(ctx context.Context) ([]*model.Category, error)
	GetCategory(ctx context.Context, id int32) (*model.Category, error)
}
