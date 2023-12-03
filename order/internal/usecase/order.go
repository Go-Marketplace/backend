package usecase

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/order/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/google/uuid"
)

type OrderUsecase struct {
	repo interfaces.OrderRepo
}

func NewOrderUsecase(repo interfaces.OrderRepo) *OrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}

func (usecase *OrderUsecase) GetOrder(ctx context.Context, orderID uuid.UUID) (*model.Order, error) {
	return usecase.repo.GetOrder(ctx, orderID)
}

func (usecase *OrderUsecase) GetOrders(ctx context.Context, searchParams dto.SearchOrderDTO) ([]*model.Order, error) {
	return usecase.repo.GetOrders(ctx, searchParams)
}

func (usecase *OrderUsecase) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	if err := usecase.repo.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	return usecase.GetOrder(ctx, order.ID)
}

func (usecase *OrderUsecase) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	return usecase.repo.DeleteOrder(ctx, orderID)
}

func (usecase *OrderUsecase) DeleteUserOrders(ctx context.Context, userID uuid.UUID) error {
	return usecase.repo.DeleteUserOrders(ctx, userID)
}

func (usecase *OrderUsecase) GetOrderline(ctx context.Context, orderID, productID uuid.UUID) (*model.Orderline, error) {
	return usecase.repo.GetOrderline(ctx, orderID, productID)
}

func (usecase *OrderUsecase) CreateOrderline(ctx context.Context, orderline *model.Orderline) error {
	return usecase.repo.CreateOrderline(ctx, orderline)
}

func (usecase *OrderUsecase) UpdateOrderline(ctx context.Context, orderline *model.Orderline) (*model.Orderline, error) {
	if err := usecase.repo.UpdateOrderline(ctx, orderline); err != nil {
		return nil, err
	}

	return usecase.GetOrderline(ctx, orderline.OrderID, orderline.ProductID)
}

func (usecase *OrderUsecase) DeleteOrderline(ctx context.Context, orderID, productID uuid.UUID) error {
	return usecase.repo.DeleteOrderline(ctx, orderID, productID)
}
