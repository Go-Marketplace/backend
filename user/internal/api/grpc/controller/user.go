package controller

import (
	"context"

	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUser(ctx context.Context, userUsecase usecase.UserUsecase, id string) (*model.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	user, err := userUsecase.GetUser(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return user, nil
}

func GetAllUsers(ctx context.Context, userUsecase usecase.UserUsecase) ([]*model.User, error) {
	users, err := userUsecase.GetAllUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return users, nil
}

func CreateUser(ctx context.Context, userUsecase usecase.UserUsecase, user model.User) error {
	err := userUsecase.CreateUser(ctx, user)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func DeleteUser(ctx context.Context, userUsecase usecase.UserUsecase, id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	err = userUsecase.DeleteUser(ctx, userID)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}
