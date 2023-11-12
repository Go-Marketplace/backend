package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
)

type orderRoutes struct {
	pbOrder.UnimplementedOrderServer

	orderUseCase usecase.OrderUseCase
	logger       *logger.Logger
}

func NewOrderRoutes(orderUseCase usecase.OrderUseCase, logger *logger.Logger) *orderRoutes {
	return &orderRoutes{
		orderUseCase: orderUseCase,
		logger:       logger,
	}
}

func (router *orderRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.GetOrderResponse, error) {
	order, err := controller.GetOrder(ctx, router.orderUseCase, req.Id)
	if err != nil {
		return nil, err
	}

	return order.ToProto(), nil
}

func (router *orderRoutes) GetAllUserOrders(ctx context.Context, req *pbOrder.GetAllUserOrdersRequest) (*pbOrder.GetAllUserOrdersResponse, error) {
	userOrders, err := controller.GetAllUserOrders(ctx, router.orderUseCase, req.UserId)
	if err != nil {
		return nil, err
	}

	userOrdersPb := make([]*pbOrder.GetOrderResponse, 0, len(userOrders))
	for _, userOrder := range userOrders {
		if userOrder != nil {
			userOrdersPb = append(userOrdersPb, userOrder.ToProto())
		}
	}

	return &pbOrder.GetAllUserOrdersResponse{
		Orders: userOrdersPb,
	}, nil
}

func (router *orderRoutes) CancelOrder(ctx context.Context, req *pbOrder.CancelOrderRequest) (*pbOrder.CancelOrderResponse, error) {
	if err := controller.CancelOrder(ctx, router.orderUseCase, req.Id); err != nil {
		return nil, err
	}

	return &pbOrder.CancelOrderResponse{}, nil
}
