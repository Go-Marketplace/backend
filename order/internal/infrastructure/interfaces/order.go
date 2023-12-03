package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/google/uuid"
)

type OrderRepo interface {
	GetOrder(ctx context.Context, orderID uuid.UUID) (*model.Order, error)
	GetOrders(ctx context.Context, searchParams dto.SearchOrderDTO) ([]*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	DeleteOrder(ctx context.Context, orderID uuid.UUID) error

	CreateOrderline(ctx context.Context, orderline *model.Orderline) error
	GetOrderline(ctx context.Context, orderID, productID uuid.UUID) (*model.Orderline, error)
	UpdateOrderline(ctx context.Context, order *model.Orderline) error
	DeleteOrderline(ctx context.Context, orderID, productID uuid.UUID) error
}
