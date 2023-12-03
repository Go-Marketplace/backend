package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
)

type orderRoutes struct {
	pbOrder.UnimplementedOrderServer

	orderUsecase  *usecase.OrderUsecase
	cartClient    pbCart.CartClient
	productClient pbProduct.ProductClient
	logger        *logger.Logger
}

func NewOrderRoutes(
	orderUsecase *usecase.OrderUsecase,
	cartClient pbCart.CartClient,
	productClient pbProduct.ProductClient,
	logger *logger.Logger,
) *orderRoutes {
	return &orderRoutes{
		orderUsecase:  orderUsecase,
		cartClient:    cartClient,
		productClient: productClient,
		logger:        logger,
	}
}

func (router *orderRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.OrderResponse, error) {
	order, err := controller.GetOrder(ctx, router.orderUsecase, req)
	if err != nil {
		return nil, err
	}

	return order.ToProto(), nil
}

func (router *orderRoutes) GetOrders(ctx context.Context, req *pbOrder.GetOrdersRequest) (*pbOrder.OrdersResponse, error) {
	orders, err := controller.GetOrders(ctx, router.orderUsecase, req)
	if err != nil {
		return nil, err
	}

	protoOrders := make([]*pbOrder.OrderResponse, 0, len(orders))
	for _, order := range orders {
		if order != nil {
			protoOrders = append(protoOrders, order.ToProto())
		}
	}

	return &pbOrder.OrdersResponse{
		Orders: protoOrders,
	}, nil
}

func (router *orderRoutes) CreateOrder(ctx context.Context, req *pbOrder.CreateOrderRequest) (*pbOrder.OrderResponse, error) {
	order, err := controller.CreateOrder(
		ctx,
		router.orderUsecase,
		router.cartClient,
		router.productClient,
		req,
	)
	if err != nil {
		return nil, err
	}

	return order.ToProto(), nil
}

func (router *orderRoutes) DeleteOrder(ctx context.Context, req *pbOrder.DeleteOrderRequest) (*pbOrder.DeleteOrderResponse, error) {
	if err := controller.DeleteOrder(ctx, router.orderUsecase, router.productClient, req); err != nil {
		return nil, err
	}

	return &pbOrder.DeleteOrderResponse{}, nil
}

func (router *orderRoutes) DeleteUserOrders(ctx context.Context, req *pbOrder.DeleteUserOrdersRequest) (*pbOrder.DeleteUserOrdersResponse, error) {
	if err := controller.DeleteUserOrders(ctx, router.orderUsecase, router.productClient, req); err != nil {
		return nil, err
	}

	return &pbOrder.DeleteUserOrdersResponse{}, nil
}

func (router *orderRoutes) GetOrderline(ctx context.Context, req *pbOrder.GetOrderlineRequest) (*pbOrder.OrderlineResponse, error) {
	orderline, err := controller.GetOrderline(ctx, router.orderUsecase, req)
	if err != nil {
		return nil, err
	}

	return orderline.ToProto(), nil
}

func (router *orderRoutes) UpdateOrderline(ctx context.Context, req *pbOrder.UpdateOrderlineRequest) (*pbOrder.OrderlineResponse, error) {
	orderline, err := controller.UpdateOrderline(ctx, router.orderUsecase, req)
	if err != nil {
		return nil, err
	}

	return orderline.ToProto(), nil
}

func (router *orderRoutes) DeleteOrderline(ctx context.Context, req *pbOrder.DeleteOrderlineRequest) (*pbOrder.DeleteOrderlineResponse, error) {
	if err := controller.DeleteOrderline(ctx, router.orderUsecase, router.productClient, req); err != nil {
		return nil, err
	}

	return &pbOrder.DeleteOrderlineResponse{}, nil
}
