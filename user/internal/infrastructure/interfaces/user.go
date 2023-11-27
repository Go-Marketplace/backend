package interfaces

import (
	"context"

	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	UpdateUser(ctx context.Context, user model.User) error
	ChangeUserRole(ctx context.Context, userID uuid.UUID, role model.UserRoles) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
