package usecase

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/user/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repo   interfaces.UserRepo
	logger *logger.Logger
}

func NewUserUsecase(repo interfaces.UserRepo, logger *logger.Logger) UserUsecase {
	return UserUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (usecase *UserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return usecase.repo.GetUser(ctx, id)
}

func (usecase *UserUsecase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return usecase.repo.GetAllUsers(ctx)
}

func (usecase *UserUsecase) CreateUser(ctx context.Context, user model.User) error {
	return usecase.repo.CreateUser(ctx, user)
}

func (usecase *UserUsecase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return usecase.repo.DeleteUser(ctx, id)
}
