package controller

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/gateway/internal/usecase"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserClaim struct {
	ID   string          `json:"id"`
	Role pbUser.UserRole `json:"role"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

type RegisterUserResponse struct {
	ID    uuid.UUID `json:"id" mapstructure:"id"`
	Token string    `json:"token" mapstructure:"token"`
}

func RegisterUser(
	ctx context.Context,
	userClient pbUser.UserClient,
	jwtManager *usecase.JWTManager,
	req *pbGateway.RegisterUserRequest,
) (*RegisterUserResponse, error) {
	user, err := userClient.GetUserByEmail(ctx, &pbUser.GetUserByEmailRequest{
		Email: req.Email,
	})
	userErrStatus, ok := status.FromError(err)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to get status from err: %s", err)
	}

	if err != nil && userErrStatus.Code() != codes.NotFound {
		return nil, status.Errorf(codes.Internal, "failed to get user by email: %s", err)
	}

	if user != nil {
		return nil, status.Errorf(codes.AlreadyExists, "user is already registered")
	}

	if req.Password, err = HashPassword(req.Password); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	id := uuid.New()
	createUserReq := &pbUser.CreateUserRequest{
		UserId:    id.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	if _, err := userClient.CreateUser(ctx, createUserReq); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	token, err := jwtManager.CreateToken(UserClaim{
		ID:   id.String(),
		Role: pbUser.UserRole_USER,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create token: %s", err)
	}

	return &RegisterUserResponse{
		ID:    id,
		Token: token,
	}, nil
}

func Login(
	ctx context.Context,
	userClient pbUser.UserClient,
	jwtManager *usecase.JWTManager,
	req *pbGateway.LoginRequest,
) (string, error) {
	user, err := userClient.GetUserByEmail(ctx, &pbUser.GetUserByEmailRequest{
		Email: req.Email,
	})
	userErrStatus, ok := status.FromError(err)
	if !ok {
		return "", status.Errorf(codes.Internal, "failed to get status from err: %s", err)
	}

	if err != nil && userErrStatus.Code() != codes.NotFound {
		return "", status.Errorf(codes.Internal, "failed to get user by email: %s", err)
	}

	if userErrStatus.Code() == codes.NotFound {
		return "", status.Errorf(codes.Canceled, "unregistered user")
	}

	if err := VerifyPassword(user.Password, req.Password); err != nil {
		return "", status.Errorf(codes.Canceled, "passwords don't match")
	}

	token, err := jwtManager.CreateToken(UserClaim{
		ID:   user.UserId,
		Role: user.Role,
	})
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to create token: %s", err)
	}

	return token, nil
}
