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

func (usecase *CartUsecase) GetCart(ctx context.Context, id uuid.UUID) (*model.Cart, error) {
	return usecase.cartRepo.GetCart(ctx, id)
}

func (usecase *CartUsecase) GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error) {
	return usecase.cartRepo.GetUserCart(ctx, userID)
}

func (usecase *CartUsecase) CreateCart(ctx context.Context, cart model.Cart) error {
	if err := usecase.cartRepo.CreateCart(ctx, cart); err != nil {
		return err
	}

	if err := usecase.cartTaskRepo.CreateCartTask(ctx, model.CartTask{
		CartID:    cart.ID,
		Timestamp: time.Now().Add(cartTTL).Unix(),
	}); err != nil {
		return err
	}

	return nil
}

func (usecase *CartUsecase) CreateCartline(ctx context.Context, cartline model.CartLine) error {
	return usecase.cartRepo.CreateCartline(ctx, cartline)
}

func (usecase *CartUsecase) UpdateCartline(ctx context.Context, cartline model.CartLine) (*model.Cart, error) {
	if err := usecase.cartRepo.UpdateCartline(ctx, cartline); err != nil {
		return nil, err
	}

	return usecase.cartRepo.GetCart(ctx, cartline.ID)
}

func (usecase *CartUsecase) DeleteCart(ctx context.Context, id uuid.UUID) error {
	return usecase.cartRepo.DeleteCart(ctx, id)
}

func (usecase *CartUsecase) DeleteCartline(ctx context.Context, id uuid.UUID) error {
	return usecase.cartRepo.DeleteCartline(ctx, id)
}

func (usecase *CartUsecase) DeleteCartCartlines(ctx context.Context, cartID uuid.UUID) error {
	return usecase.cartRepo.DeleteCartCartlines(ctx, cartID)
}
