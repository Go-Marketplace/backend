package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func CreateCart(ctx context.Context, cartUsecase *usecase.CartUsecase, req *pbCart.CreateCartRequest) (*model.Cart, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	cart := model.Cart{
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = cartUsecase.CreateCart(ctx, cart); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return &cart, nil
}

func CreateCartline(
	ctx context.Context,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	req *pbCart.CreateCartlineRequest,
) (*model.CartLine, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	productResp, err := productClient.GetProduct(ctx, &pbProduct.GetProductRequest{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get product: %s", err)
	}

	cartline := &model.CartLine{
		UserID:    userID,
		ProductID: productID,
		Name:      productResp.Name,
		Quantity:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = cartUsecase.CreateCartline(ctx, cartline); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create cartline: %s", err)
	}

	if _, err = productClient.UpdateProduct(ctx, &pbProduct.UpdateProductRequest{
		ProductId: req.ProductId,
		Quantity:  productResp.Quantity - 1,
	}); err != nil {
		if err = cartUsecase.DeleteCartline(ctx, cartline.UserID, cartline.ProductID); err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to delete cartline: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to update product: %s", err)
	}

	return cartline, nil
}

func returnProducts(ctx context.Context, productClient pbProduct.ProductClient, cartlines ...*model.CartLine) error {
	productIDs := make([]string, 0, len(cartlines))
	for _, cartline := range cartlines {
		productIDs = append(productIDs, cartline.ProductID.String())
	}

	productsResp, err := productClient.GetProducts(ctx, &pbProduct.GetProductsRequest{
		ProductIds: productIDs,
	})
	if err != nil {
		return fmt.Errorf("failed to get products: %w", err)
	}

	if len(productIDs) != len(productsResp.Products) {
		return fmt.Errorf("cartline doesn't match with product")
	}

	updateProductRequests := make([]*pbProduct.UpdateProductRequest, 0, len(cartlines))
	for i, cartline := range cartlines {
		updateProductRequests = append(updateProductRequests, &pbProduct.UpdateProductRequest{
			ProductId: cartline.ProductID.String(),
			Quantity:  productsResp.Products[i].Quantity + cartline.Quantity,
		})
	}

	if _, err = productClient.UpdateProducts(ctx, &pbProduct.UpdateProductsRequest{
		Products: updateProductRequests,
	}); err != nil {
		return fmt.Errorf("failed to update products: %w", err)
	}

	return nil
}

func UpdateCartline(
	ctx context.Context,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	req *pbCart.UpdateCartlineRequest,
) (*model.CartLine, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	oldCartline, err := cartUsecase.GetCartline(ctx, userID, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if oldCartline == nil {
		return nil, status.Errorf(codes.NotFound, "Cartline not found")
	}

	newCartline := model.CartLine{
		UserID:    userID,
		ProductID: productID,
		Name:      req.Name,
		Quantity:  req.Quantity,
	}

	cartline, err := cartUsecase.UpdateCartline(ctx, newCartline)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update cartline: %s", err)
	}

	if req.Quantity != 0 {
		diff := oldCartline.Quantity - req.Quantity

		productResp, err := productClient.GetProduct(ctx, &pbProduct.GetProductRequest{
			ProductId: req.ProductId,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to get product: %s", err)
		}

		if _, err = productClient.UpdateProduct(ctx, &pbProduct.UpdateProductRequest{
			ProductId: req.ProductId,
			Quantity:  productResp.Quantity + diff,
		}); err != nil {
			if _, err = cartUsecase.UpdateCartline(ctx, *oldCartline); err != nil {
				return nil, status.Errorf(codes.Internal, "Failed to update cartline: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "Failed to update product: %s", err)
		}
	}

	return cartline, nil
}

func DeleteCart(
	ctx context.Context,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	req *pbCart.DeleteCartRequest,
) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	cart, err := cartUsecase.GetUserCart(ctx, userID)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if cart == nil {
		return status.Errorf(codes.NotFound, "Cart not found")
	}

	if err = cartUsecase.DeleteCart(ctx, userID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete cart: %s", err)
	}

	if err = returnProducts(ctx, productClient, cart.Cartlines...); err != nil {
		if err = cartUsecase.CreateCart(ctx, *cart); err != nil {
			return status.Errorf(codes.Internal, "Failed to create cart: %s", err)
		}
		if err = cartUsecase.CreateCartlines(ctx, cart.Cartlines); err != nil {
			return status.Errorf(codes.Internal, "Failed to create cartlines: %s", err)
		}
		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}

func DeleteCartline(
	ctx context.Context,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	req *pbCart.DeleteCartlineRequest,
) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	cartline, err := cartUsecase.GetCartline(ctx, userID, productID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get cartline: %s", err)
	}

	if err = cartUsecase.DeleteCartline(ctx, userID, productID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete cartline: %s", err)
	}

	if err = returnProducts(ctx, productClient, cartline); err != nil {
		if err = cartUsecase.CreateCartline(ctx, cartline); err != nil {
			return status.Errorf(codes.Internal, "Failed to create cartline: %s", err)
		}
		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}

func DeleteCartCartlines(
	ctx context.Context,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	req *pbCart.DeleteCartCartlinesRequest,
) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	cart, err := cartUsecase.GetUserCart(ctx, userID)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if cart == nil {
		return status.Errorf(codes.NotFound, "Cart not found")
	}

	if err = cartUsecase.DeleteCartCartlines(ctx, userID); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if err = returnProducts(ctx, productClient, cart.Cartlines...); err != nil {
		if err = cartUsecase.CreateCartlines(ctx, cart.Cartlines); err != nil {
			return status.Errorf(codes.Internal, "Failed to create cartlines: %s", err)
		}
		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}
