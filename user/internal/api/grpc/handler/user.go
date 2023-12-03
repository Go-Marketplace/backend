package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
)

type userRoutes struct {
	pbUser.UnimplementedUserServer

	userUsecase   *usecase.UserUsecase
	productClient pbProduct.ProductClient
	orderClient   pbOrder.OrderClient
	cartClient    pbCart.CartClient
	logger        *logger.Logger
}

func NewUserRoutes(
	userUsecase *usecase.UserUsecase,
	productClient pbProduct.ProductClient,
	orderClient pbOrder.OrderClient,
	cartClient pbCart.CartClient,
	logger *logger.Logger,
) *userRoutes {
	return &userRoutes{
		userUsecase:   userUsecase,
		productClient: productClient,
		orderClient:   orderClient,
		cartClient:    cartClient,
		logger:        logger,
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

func (router *userRoutes) GetUsers(ctx context.Context, req *pbUser.GetUsersRequest) (*pbUser.UsersResponse, error) {
	users, err := controller.GetUsers(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*pbUser.UserResponse, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, user.ToProto())
	}

	return &pbUser.UsersResponse{
		Users: protoUsers,
	}, nil
}

func (router *userRoutes) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.UserResponse, error) {
	user, err := controller.CreateUser(ctx, router.userUsecase, router.cartClient, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UserResponse, error) {
	user, err := controller.UpdateUser(ctx, router.userUsecase, req)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error) {
	if err := controller.DeleteUser(
		ctx,
		router.userUsecase,
		router.orderClient,
		router.productClient,
		router.cartClient,
		req,
	); err != nil {
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
