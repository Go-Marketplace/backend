package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/google/uuid"
)

type CartRepo interface {
	GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error)
	CreateCart(ctx context.Context, cart model.Cart) error
	DeleteCart(ctx context.Context, id uuid.UUID) error
	DeleteCartCartlines(ctx context.Context, cartID uuid.UUID) error

	CreateCartline(ctx context.Context, cartline model.CartLine) error
	UpdateCartline(ctx context.Context, cartline model.CartLine) error
	DeleteCartline(ctx context.Context, cartline model.CartLine) error
}
