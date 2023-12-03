package usecase

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/user/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
)

type IUserUsecase interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	ChangeUserRole(ctx context.Context, userID uuid.UUID, role model.UserRoles) (*model.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type UserUsecase struct {
	repo   interfaces.UserRepo
	logger *logger.Logger
}

func NewUserUsecase(repo interfaces.UserRepo, logger *logger.Logger) *UserUsecase {
	return &UserUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (usecase *UserUsecase) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return usecase.repo.GetUser(ctx, userID)
}

func (usecase *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return usecase.repo.GetUserByEmail(ctx, email)
}

func (usecase *UserUsecase) GetUsers(ctx context.Context) ([]*model.User, error) {
	return usecase.repo.GetUsers(ctx)
}

func (usecase *UserUsecase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := usecase.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return usecase.GetUser(ctx, user.ID)
}

func (usecase *UserUsecase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return usecase.repo.DeleteUser(ctx, userID)
}

func (usecase *UserUsecase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := usecase.repo.UpdateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return usecase.repo.GetUser(ctx, user.ID)
}

func (usecase *UserUsecase) ChangeUserRole(ctx context.Context, userID uuid.UUID, role model.UserRoles) (*model.User, error) {
	if err := usecase.repo.ChangeUserRole(ctx, userID, role); err != nil {
		return nil, fmt.Errorf("failed to change user role: %w", err)
	}

	return usecase.repo.GetUser(ctx, userID)
}
