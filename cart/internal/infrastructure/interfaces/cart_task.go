package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/cart/internal/model"
)

type CartTaskRepo interface {
	CreateCartTask(ctx context.Context, task model.CartTask) error
	GetCartTasks(ctx context.Context, to int64) ([]*model.CartTask, error)
}
