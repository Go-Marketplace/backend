package usecase

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/google/uuid"
)

const (
	cartTTL = 5 * time.Minute
)

type CartUsecase struct {
	cartTaskRepo interfaces.CartTaskRepo
	cartRepo     interfaces.CartRepo
	logger       *logger.Logger
}

func NewCartUsecase(cartRepo interfaces.CartRepo, cartTaskRepo interfaces.CartTaskRepo, logger *logger.Logger) *CartUsecase {
	return &CartUsecase{
		cartTaskRepo: cartTaskRepo,
		cartRepo:     cartRepo,
		logger:       logger,
	}
}

func (usecase *CartUsecase) GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error) {
	return usecase.cartRepo.GetUserCart(ctx, userID)
}

func (usecase *CartUsecase) CreateCart(ctx context.Context, cart model.Cart) error {
	if err := usecase.cartRepo.CreateCart(ctx, cart); err != nil {
		return err
	}

	if err := usecase.cartTaskRepo.CreateCartTask(ctx, model.CartTask{
		UserID:    cart.UserID,
		Timestamp: time.Now().Add(cartTTL).Unix(),
	}); err != nil {
		return err
	}

	return nil
}

func (usecase *CartUsecase) GetCartline(ctx context.Context, userID, productID uuid.UUID) (*model.CartLine, error) {
	return usecase.cartRepo.GetCartline(ctx, userID, productID)
}

func (usecase *CartUsecase) CreateCartline(ctx context.Context, cartline *model.CartLine) error {
	return usecase.cartRepo.CreateCartline(ctx, cartline)
}

func (usecase *CartUsecase) CreateCartlines(ctx context.Context, cartlines []*model.CartLine) error {
	return usecase.cartRepo.CreateCartlines(ctx, cartlines)
}

func (usecase *CartUsecase) UpdateCartline(ctx context.Context, cartline model.CartLine) (*model.CartLine, error) {
	if err := usecase.cartRepo.UpdateCartline(ctx, cartline); err != nil {
		return nil, err
	}

	return usecase.GetCartline(ctx, cartline.UserID, cartline.ProductID)
}

func (usecase *CartUsecase) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	return usecase.cartRepo.DeleteCart(ctx, userID)
}

func (usecase *CartUsecase) DeleteCartline(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error {
	return usecase.cartRepo.DeleteCartline(ctx, userID, productID)
}

func (usecase *CartUsecase) DeleteCartCartlines(ctx context.Context, userID uuid.UUID) error {
	return usecase.cartRepo.DeleteCartCartlines(ctx, userID)
}
