package controller

import (
	"context"
	"fmt"
	"time"

	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUser(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.GetUserRequest) (*model.User, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	user, err := userUsecase.GetUser(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.GetUserByEmailRequest) (*model.User, error) {
	user, err := userUsecase.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func GetAllUsers(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.GetAllUsersRequest) ([]*model.User, error) {
	users, err := userUsecase.GetAllUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return users, nil
}

func CreateUser(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.CreateUserRequest) error {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}

	user := model.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      model.RegisteredUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = userUsecase.CreateUser(ctx, user); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func UpdateUser(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.UpdateUserRequest) (*model.User, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	newUser := model.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		Phone:     req.Phone,
		UpdatedAt: time.Now(),
	}

	user, err := userUsecase.UpdateUser(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func DeleteUser(ctx context.Context, userUsecase usecase.UserUsecase, req *pbUser.DeleteUserRequest) error {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	err = userUsecase.DeleteUser(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}
