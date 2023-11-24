package controller

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetOrder(ctx context.Context, orderUseCase usecase.OrderUseCase, id string) (*model.Order, error) {
	orderID, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	order, err := orderUseCase.GetOrder(ctx, orderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if order == nil {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return order, nil
}

func GetAllUserOrders(ctx context.Context, orderUseCase usecase.OrderUseCase, id string) ([]*model.Order, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	userOrders, err := orderUseCase.GetAllUserOrders(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return userOrders, nil
}

func CancelOrder(ctx context.Context, orderUseCase usecase.OrderUseCase, id string) error {
	orderID, err := uuid.Parse(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	err = orderUseCase.CancelOrder(ctx, orderID)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}