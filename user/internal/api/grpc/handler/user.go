package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userRoutes struct {
	pbUser.UnimplementedUserServer

	userUsecase usecase.UserUsecase
	logger      *logger.Logger
}

func NewUserRoutes(userUsecase usecase.UserUsecase, logger *logger.Logger) *userRoutes {
	return &userRoutes{
		userUsecase: userUsecase,
		logger:      logger,
	}
}

func (router *userRoutes) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.UserResponse, error) {
	user, err := controller.GetUser(ctx, router.userUsecase, req.Id)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (router *userRoutes) GetAllUsers(ctx context.Context, req *pbUser.GetAllUsersRequest) (*pbUser.GetAllUsersResponse, error) {
	users, err := controller.GetAllUsers(ctx, router.userUsecase)
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

func (router *userRoutes) CreateUser(ctx context.Context, req *pbUser.UserRequest) (*pbUser.UserResponse, error) {
	user, err := model.FromProtoToUser(req)
	if err != nil || user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user: %s", err)
	}

	err = controller.CreateUser(ctx, router.userUsecase, *user)
	if err != nil {
		return nil, err
	}

	return &pbUser.UserResponse{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (router *userRoutes) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error) {
	err := controller.DeleteUser(ctx, router.userUsecase, req.Id)
	if err != nil {
		return nil, err
	}

	return &pbUser.DeleteUserResponse{}, nil
}
