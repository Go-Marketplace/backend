package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
