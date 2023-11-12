package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
)

type gatewayRoutes struct {
	pbGateway.GatewayServer

	orderClient pbOrder.OrderClient
	logger      *logger.Logger
}

func NewGatewayRoutes(orderClient pbOrder.OrderClient, logger *logger.Logger) *gatewayRoutes {
	return &gatewayRoutes{
		orderClient: orderClient,
		logger:      logger,
	}
}

func (router *gatewayRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.GetOrderResponse, error) {
	return router.orderClient.GetOrder(ctx, req)
}
