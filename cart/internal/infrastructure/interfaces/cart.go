package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/google/uuid"
)

type CartRepo interface {
	GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error)
	CreateCart(ctx context.Context, cart model.Cart) error
	DeleteCart(ctx context.Context, userID uuid.UUID) error
	DeleteCartCartlines(ctx context.Context, userID uuid.UUID) error

	GetCartline(ctx context.Context, userID, productID uuid.UUID) (*model.CartLine, error)
	CreateCartline(ctx context.Context, cartline *model.CartLine) error
	CreateCartlines(ctx context.Context, cartlines []*model.CartLine) error
	UpdateCartline(ctx context.Context, cartline model.CartLine) error
	DeleteCartline(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error
	DeleteProductCartlines(ctx context.Context, productID uuid.UUID) error
}
