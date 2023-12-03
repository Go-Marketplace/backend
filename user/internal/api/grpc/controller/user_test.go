package controller_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/api/grpc/controller"
	mocks "github.com/Go-Marketplace/backend/user/internal/mocks/usecase"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func userHelper(t *testing.T) *mocks.MockIUserUsecase {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mocks.NewMockIUserUsecase(mockCtrl)
}

func TestChangeUserRole(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *pbUser.ChangeUserRoleRequest
	}

	ctx := context.Background()
	userID := uuid.New()
	var role model.UserRoles = 1

	expectedUserFromRepo := &model.User{
		ID:        userID,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.ru",
	}
	expectedErrFromUsecase := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(usecase *mocks.MockIUserUsecase)
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name: "Successfully change user role",
			args: args{
				ctx: ctx,
				req: &pbUser.ChangeUserRoleRequest{
					UserId: userID.String(),
					Role:   pbUser.UserRole(role),
				},
			},
			mock: func(usecase *mocks.MockIUserUsecase) {
				usecase.EXPECT().ChangeUserRole(ctx, userID, role).Return(expectedUserFromRepo, nil).Times(1)
			},
			expectedUser: expectedUserFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Invalid user id in request",
			args: args{
				ctx: ctx,
				req: &pbUser.ChangeUserRoleRequest{
					UserId: "123",
					Role:   pbUser.UserRole(role),
				},
			},
			mock:         func(usecase *mocks.MockIUserUsecase) {},
			expectedUser: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "Invalid user id: %s", fmt.Errorf("invalid UUID length: 3")),
		},
		{
			name: "Failed user validation",
			args: args{
				ctx: ctx,
				req: &pbUser.ChangeUserRoleRequest{
					UserId: userID.String(),
					Role:   pbUser.UserRole(9),
				},
			},
			mock:         func(usecase *mocks.MockIUserUsecase) {},
			expectedUser: nil,
			expectedErr:  status.Errorf(codes.InvalidArgument, "Invalid request: %s", fmt.Errorf("Key: 'User.Role' Error:Field validation for 'Role' failed on the 'oneof' tag")),
		},
		{
			name: "Got error when change user role in usecase",
			args: args{
				ctx: ctx,
				req: &pbUser.ChangeUserRoleRequest{
					UserId: userID.String(),
					Role:   pbUser.UserRole(role),
				},
			},
			mock: func(usecase *mocks.MockIUserUsecase) {
				usecase.EXPECT().ChangeUserRole(ctx, userID, role).Return(nil, expectedErrFromUsecase).Times(1)
			},
			expectedUser: nil,
			expectedErr:  status.Errorf(codes.Internal, "Failed to change user role: %s", expectedErrFromUsecase),
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase := userHelper(t)
			testcase.mock(userUsecase)

			actualUser, actualErr := controller.ChangeUserRole(
				testcase.args.ctx,
				userUsecase,
				testcase.args.req,
			)

			assert.Equal(t, testcase.expectedUser, actualUser)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
