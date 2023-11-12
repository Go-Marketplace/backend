package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/google/uuid"
)

type OrderRepo interface {
	GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error)
	GetAllUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error)
	CreateOrder(ctx context.Context, order model.Order) error
	CancelOrder(ctx context.Context, id uuid.UUID) error
}
