package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
)

type ProductRepo interface {
	GetProducts(ctx context.Context, searchParams dto.SearchProductsDTO) ([]*model.Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error

	GetAllCategories(ctx context.Context) ([]*model.Category, error)
	GetCategory(ctx context.Context, id int32) (*model.Category, error)
}
