package usecase

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/google/uuid"
)

type OrderUseCase struct {
	repo interfaces.OrderRepo
}

func NewOrderUseCase(repo interfaces.OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
	}
}

func (usecase *OrderUseCase) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	return usecase.repo.GetOrder(ctx, id)
}

func (usecase *OrderUseCase) GetAllUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	return usecase.repo.GetAllUserOrders(ctx, userID)
}

func (usecase *OrderUseCase) CreateOrder(ctx context.Context, order model.Order) error {
	return usecase.repo.CreateOrder(ctx, order)
}

func (usecase *OrderUseCase) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	return usecase.repo.DeleteOrder(ctx, id)
}
