package controller

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetCart(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.GetCartRequest) (*model.Cart, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	cart, err := cartUsecase.GetCart(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if cart == nil {
		return nil, status.Errorf(codes.NotFound, "Cart not found")
	}

	return cart, nil
}

func GetUserCart(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.GetUserCartRequest) (*model.Cart, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	cart, err := cartUsecase.GetUserCart(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if cart == nil {
		return nil, status.Errorf(codes.NotFound, "Cart not found")
	}

	return cart, nil
}

func CreateCart(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.CreateCartRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	cart := model.Cart{
		ID:        uuid.New(),
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = cartUsecase.CreateCart(ctx, cart); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func CreateCartline(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.CreateCartlineRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	cartID, err := uuid.Parse(req.CartId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid cart id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	cartline := model.CartLine{
		ID:        uuid.New(),
		CartID:    cartID,
		ProductID: productID,
		Name:      req.Name,
		Quantity:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = cartUsecase.CreateCartline(ctx, cartline); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func UpdateCartline(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.UpdateCartlineRequest) (*model.Cart, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	cartID, err := uuid.Parse(req.CartId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid cart id: %s", err)
	}

	cartline := model.CartLine{
		CartID:    cartID,
		Name:      req.Name,
		Quantity:  req.Quantity,
		UpdatedAt: time.Now(),
	}

	cart, err := cartUsecase.UpdateCartline(ctx, cartline)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return cart, nil
}

func DeleteCart(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.DeleteCartRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	if err := cartUsecase.DeleteCart(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func DeleteCartline(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.DeleteCartlineRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	if err := cartUsecase.DeleteCartline(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func DeleteCartCartlines(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.DeleteCartCartlinesRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	cartID, err := uuid.Parse(req.CartId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid cart id: %s", err)
	}

	if err := cartUsecase.DeleteCartCartlines(ctx, cartID); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}
