package controller

import (
	"context"
	"time"

	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUser(ctx context.Context, userUsecase usecase.IUserUsecase, req *pbUser.GetUserRequest) (*model.User, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	user, err := userUsecase.GetUser(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, userUsecase usecase.IUserUsecase, req *pbUser.GetUserByEmailRequest) (*model.User, error) {
	user, err := userUsecase.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get user by email: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func GetUsers(ctx context.Context, userUsecase usecase.IUserUsecase, req *pbUser.GetUsersRequest) ([]*model.User, error) {
	users, err := userUsecase.GetUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get users: %s", err)
	}

	return users, nil
}

func CreateUser(
	ctx context.Context,
	userUsecase usecase.IUserUsecase,
	cartClient pbCart.CartClient,
	req *pbUser.CreateUserRequest,
) (*model.User, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	newUser := &model.User{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      model.RegisteredUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = newUser.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request user: %s", err)
	}

	user, err := userUsecase.CreateUser(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %s", err)
	}

	if _, err = cartClient.CreateCart(ctx, &pbCart.CreateCartRequest{
		UserId: req.UserId,
	}); err != nil {
		if errDelete := userUsecase.DeleteUser(ctx, userID); errDelete != nil {
			return nil, status.Errorf(codes.Internal, "Failed to delete user: %s", errDelete)
		}
		return nil, status.Errorf(codes.Internal, "Failed to create cart: %s", err)
	}

	return user, nil
}

func UpdateUser(ctx context.Context, userUsecase usecase.IUserUsecase, req *pbUser.UpdateUserRequest) (*model.User, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	newUser := &model.User{
		ID:        userID,
		FirstName: internal.Unwrap(req.FirstName),
		LastName:  internal.Unwrap(req.LastName),
		Address:   internal.Unwrap(req.Address),
		Phone:     internal.Unwrap(req.Phone),
		UpdatedAt: time.Now(),
	}

	if err = newUser.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request user: %s", err)
	}

	user, err := userUsecase.UpdateUser(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user: %s", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}

func DeleteUser(
	ctx context.Context,
	userUsecase usecase.IUserUsecase,
	orderClient pbOrder.OrderClient,
	productClient pbProduct.ProductClient,
	cartClient pbCart.CartClient,
	req *pbUser.DeleteUserRequest,
) error {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	user, err := userUsecase.GetUser(ctx, userID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get user: %s", err)
	}

	if user == nil {
		return status.Errorf(codes.NotFound, "User not found")
	}

	if err = userUsecase.DeleteUser(ctx, userID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete user: %s", err)
	}

	if _, err = orderClient.DeleteUserOrders(ctx, &pbOrder.DeleteUserOrdersRequest{
		UserId: req.UserId,
	}); err != nil {
		if _, errCreate := userUsecase.CreateUser(ctx, user); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to create user: %s", errCreate)
		}
		return status.Errorf(codes.Internal, "Failed to delete user orders: %s", err)
	}

	if _, err = cartClient.DeleteCart(ctx, &pbCart.DeleteCartRequest{
		UserId: req.UserId,
	}); err != nil {
		if _, errCreate := userUsecase.CreateUser(ctx, user); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to create user: %s", errCreate)
		}
		return status.Errorf(codes.Internal, "Failed to delete user cart: %s", err)
	}

	if _, err = productClient.DeleteUserProducts(ctx, &pbProduct.DeleteUserProductsRequest{
		UserId: req.UserId,
	}); err != nil {
		if _, errCreate := userUsecase.CreateUser(ctx, user); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to create user: %s", errCreate)
		}
		return status.Errorf(codes.Internal, "Failed to delete user products: %s", err)
	}

	return nil
}

func ChangeUserRole(ctx context.Context, userUsecase usecase.IUserUsecase, req *pbUser.ChangeUserRoleRequest) (*model.User, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	if req.Role >= pbUser.UserRole_SUPERADMIN || req.Role <= pbUser.UserRole_GUEST {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid role request")
	}

	user, err := userUsecase.ChangeUserRole(ctx, userID, model.UserRoles(req.Role))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to change user role: %s", err)
	}

	return user, nil
}
