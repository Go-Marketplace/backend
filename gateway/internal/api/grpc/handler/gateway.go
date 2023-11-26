package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/gateway/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
)

type gatewayRoutes struct {
	pbGateway.GatewayServer

	orderClient pbOrder.OrderClient
	userClient  pbUser.UserClient
	jwtManager  *usecase.JWTManager
	logger      *logger.Logger
}

func NewGatewayRoutes(
	orderClient pbOrder.OrderClient,
	userClient pbUser.UserClient,
	jwtManager *usecase.JWTManager,
	logger *logger.Logger,
) *gatewayRoutes {
	return &gatewayRoutes{
		orderClient: orderClient,
		userClient:  userClient,
		jwtManager:  jwtManager,
		logger:      logger,
	}
}

func (router *gatewayRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.GetOrder(ctx, req)
}

func (router *gatewayRoutes) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.UserResponse, error) {
	return router.userClient.GetUser(ctx, req)
}

func (router *gatewayRoutes) RegisterUser(ctx context.Context, req *pbGateway.RegisterUserRequest) (*pbGateway.RegisterUserResponse, error) {
	resp, err := controller.RegisterUser(ctx, router.userClient, router.jwtManager, req)
	if err != nil {
		return nil, err
	}

	return &pbGateway.RegisterUserResponse{
		UserId:      resp.ID.String(),
		AccessToken: resp.Token,
	}, nil
}

func (router *gatewayRoutes) Login(ctx context.Context, req *pbGateway.LoginRequest) (*pbGateway.LoginResponse, error) {
	token, err := controller.Login(ctx, router.userClient, router.jwtManager, req)
	if err != nil {
		return nil, err
	}

	return &pbGateway.LoginResponse{
		AccessToken: token,
	}, nil
}
