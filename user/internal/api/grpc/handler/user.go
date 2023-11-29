package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
)

type userRoutes struct {
	pbUser.UnimplementedUserServer

	userUsecase *usecase.UserUsecase
	logger      *logger.Logger
}

func NewUserRoutes(userUsecase *usecase.UserUsecase, logger *logger.Logger) *userRoutes {
	return &userRoutes{
		userUsecase: userUsecase,
		logger:      logger,
	}
}

func (router *userRoutes) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.UserResponse, error) {
	user, err := controller.GetUser(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) GetUserByEmail(ctx context.Context, req *pbUser.GetUserByEmailRequest) (*pbUser.UserResponse, error) {
	user, err := controller.GetUserByEmail(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) GetAllUsers(ctx context.Context, req *pbUser.GetAllUsersRequest) (*pbUser.GetAllUsersResponse, error) {
	users, err := controller.GetAllUsers(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*pbUser.UserResponse, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, user.ToProto())
	}

	return &pbUser.GetAllUsersResponse{
		Users: protoUsers,
	}, nil
}

func (router *userRoutes) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error) {
	if err := controller.CreateUser(ctx, router.userUsecase, req); err != nil {
		return nil, err
	}

	return &pbUser.CreateUserResponse{
		UserId: req.UserId,
	}, nil
}

func (router *userRoutes) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UserResponse, error) {
	user, err := controller.UpdateUser(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error) {
	err := controller.DeleteUser(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return &pbUser.DeleteUserResponse{}, nil
}

func (router *userRoutes) ChangeUserRole(ctx context.Context, req *pbUser.ChangeUserRoleRequest) (*pbUser.UserResponse, error) {
	user, err := controller.ChangeUserRole(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}
