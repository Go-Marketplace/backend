package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
)

type DiscountRepo interface {
	CreateDiscount(ctx context.Context, discount model.Discount) error
	GetDiscount(ctx context.Context, productID uuid.UUID) (*model.Discount, error)
	DeleteDiscount(ctx context.Context, productID uuid.UUID) error
}
