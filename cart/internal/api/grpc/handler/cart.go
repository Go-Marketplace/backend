package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/cart/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
)

type cartRoutes struct {
	pbCart.UnimplementedCartServer

	cartUsecase *usecase.CartUsecase
	logger      *logger.Logger
}

func NewCartRoutes(cartUsecase *usecase.CartUsecase, logger *logger.Logger) *cartRoutes {
	return &cartRoutes{
		cartUsecase: cartUsecase,
		logger:      logger,
	}
}

func (router *cartRoutes) GetUserCart(ctx context.Context, req *pbCart.GetUserCartRequest) (*pbCart.CartResponse, error) {
	cart, err := controller.GetUserCart(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cart.ToProto(), nil
}

func (router *cartRoutes) CreateCart(ctx context.Context, req *pbCart.CreateCartRequest) (*pbCart.CartResponse, error) {
	cart, err := controller.CreateCart(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cart.ToProto(), nil
}

func (router *cartRoutes) CreateCartline(ctx context.Context, req *pbCart.CreateCartlineRequest) (*pbCart.CartlineResponse, error) {
	cartline, err := controller.CreateCartline(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cartline.ToProto(), nil
}

func (router *cartRoutes) UpdateCartline(ctx context.Context, req *pbCart.UpdateCartlineRequest) (*pbCart.CartResponse, error) {
	cart, err := controller.UpdateCartline(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cart.ToProto(), nil
}

func (router *cartRoutes) DeleteCart(ctx context.Context, req *pbCart.DeleteCartRequest) (*pbCart.DeleteCartResponse, error) {
	if err := controller.DeleteCart(ctx, router.cartUsecase, req); err != nil {
		return nil, err
	}

	return &pbCart.DeleteCartResponse{}, nil
}

func (router *cartRoutes) DeleteCartline(ctx context.Context, req *pbCart.DeleteCartlineRequest) (*pbCart.CartResponse, error) {
	cart, err := controller.DeleteCartline(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cart.ToProto(), nil
}

func (router *cartRoutes) DeleteCartCartlines(ctx context.Context, req *pbCart.DeleteCartCartlinesRequest) (*pbCart.CartResponse, error) {
	cart, err := controller.DeleteCartCartlines(ctx, router.cartUsecase, req)
	if err != nil {
		return nil, err
	}

	return cart.ToProto(), nil
}
