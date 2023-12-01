package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/pkg/logger"
	redisWrap "github.com/Go-Marketplace/backend/pkg/redis"
	"github.com/go-redis/redis/v8"
)

type CartTaskRepo struct {
	redis  *redisWrap.Redis
	logger *logger.Logger
}

func NewCartTaskRepo(redis *redisWrap.Redis, logger *logger.Logger) *CartTaskRepo {
	return &CartTaskRepo{
		redis:  redis,
		logger: logger,
	}
}

func (repo *CartTaskRepo) CreateCartTask(ctx context.Context, task model.CartTask) error {
	taskBinary, err := task.MarshalBinary()
	if err != nil {
		return err
	}

	z := &redis.Z{
		Score:  float64(task.Timestamp),
		Member: taskBinary,
	}
	err = repo.redis.Client.ZAdd(ctx, cartTasksKey, z).Err()
	if err != nil {
		return fmt.Errorf("failed to ZAdd new task: %w", err)
	}

	repo.logger.Info("Create cart task for %s cart", task.UserID)

	return nil
}

func (repo *CartTaskRepo) GetCartTasks(ctx context.Context, to int64) ([]*model.CartTask, error) {
	pipe := repo.redis.Client.TxPipeline()
	pipe.ZRangeByScore(ctx, cartTasksKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: strconv.FormatInt(to, 10),
	})
	pipe.ZRemRangeByScore(ctx, cartTasksKey, "-inf", strconv.FormatInt(to, 10))

	response, err := pipe.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Exec getCartTasks: %w", err)
	}

	data, err := response[0].(*redis.StringSliceCmd).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to read CartTasks: %w", err)
	}

	tasks := make([]*model.CartTask, 0, len(data))
	for _, val := range data {
		task := &model.CartTask{}
		if err := task.UnmarshalBinary([]byte(val)); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

const cartTasksKey = "cart-tasks"
