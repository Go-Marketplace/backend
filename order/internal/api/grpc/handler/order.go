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

func (router *orderRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.OrderResponse, error) {
	order, err := controller.GetOrder(ctx, router.orderUseCase, req)
	if err != nil {
		return nil, err
	}

	return order.ToProto(), nil
}

func (router *orderRoutes) GetAllUserOrders(ctx context.Context, req *pbOrder.GetAllUserOrdersRequest) (*pbOrder.GetAllUserOrdersResponse, error) {
	userOrders, err := controller.GetAllUserOrders(ctx, router.orderUseCase, req)
	if err != nil {
		return nil, err
	}

	protoUserOrders := make([]*pbOrder.OrderResponse, 0, len(userOrders))
	for _, userOrder := range userOrders {
		if userOrder != nil {
			protoUserOrders = append(protoUserOrders, userOrder.ToProto())
		}
	}

	return &pbOrder.GetAllUserOrdersResponse{
		Orders: protoUserOrders,
	}, nil
}

func (router *orderRoutes) CancelOrder(ctx context.Context, req *pbOrder.DeleteOrderRequest) (*pbOrder.DeleteOrderResponse, error) {
	if err := controller.DeleteOrder(ctx, router.orderUseCase, req); err != nil {
		return nil, err
	}

	return &pbOrder.DeleteOrderResponse{}, nil
}
