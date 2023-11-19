package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Go-Marketplace/backend/pkg/logger"
	redisWrap "github.com/Go-Marketplace/backend/pkg/redis"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type DiscountRepo struct {
	redis  *redisWrap.Redis
	logger *logger.Logger
}

func NewDiscountRepo(redis *redisWrap.Redis, logger *logger.Logger) *DiscountRepo {
	return &DiscountRepo{
		redis:  redis,
		logger: logger,
	}
}

func (repo *DiscountRepo) CreateDiscount(ctx context.Context, discount model.Discount) error {
	discountBinary, err := discount.MarshalBinary()
	if err != nil {
		return err
	}

	expiration := discount.EndedAt.Sub(discount.CreatedAt)
	if expiration.Seconds() <= 0 {
		return fmt.Errorf("expirations's ended_at cannot be less or equal than created_at")
	}

	err = repo.redis.Client.SetEX(
		ctx,
		discountKey(discount.ProductID.String()),
		discountBinary,
		expiration,
	).Err()
	if err != nil {
		return fmt.Errorf("failed to SetEX discount: %w", err)
	}

	return nil
}

func (repo *DiscountRepo) GetDiscount(ctx context.Context, productID uuid.UUID) (*model.Discount, error) {
	discountBinary, err := repo.redis.Client.Get(ctx, discountKey(productID.String())).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("failed to Get discount: %w", err)
	}

	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	discount := &model.Discount{}
	err = discount.UnmarshalBinary([]byte(discountBinary))
	if err != nil {
		return nil, err
	}

	return discount, nil
}

func (repo *DiscountRepo) DeleteDiscount(ctx context.Context, productID uuid.UUID) error {
	err := repo.redis.Client.Del(ctx, discountKey(productID.String())).Err()
	if err != nil {
		return fmt.Errorf("failed to Del discount: %w", err)
	}

	return nil
}

func discountKey(productID string) string {
	return "discount:" + productID
}
