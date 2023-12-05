package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Go-Marketplace/backend/cart/internal/api/grpc/controller"
	mocks "github.com/Go-Marketplace/backend/cart/internal/mocks/usecase"
	"github.com/Go-Marketplace/backend/cart/internal/model"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cartHelper(t *testing.T) *mocks.MockICartUsecase {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mocks.NewMockICartUsecase(mockCtrl)
}

func TestGetUserCart(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *pbCart.GetUserCartRequest
	}

	ctx := context.Background()
	userID := uuid.New()

	expectedCartFromUsecase := &model.Cart{
		UserID: userID,
	}
	expectedErrFromUsecase := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(usecase *mocks.MockICartUsecase)
		expectedCart *model.Cart
		expectedErr  error
	}{
		{
			name: "Successfully get cart",
			args: args{
				ctx: ctx,
				req: &pbCart.GetUserCartRequest{
					UserId: userID.String(),
				},
			},
			mock: func(usecase *mocks.MockICartUsecase) {
				usecase.EXPECT().GetUserCart(ctx, userID).Return(expectedCartFromUsecase, nil).Times(1)
			},
			expectedCart: expectedCartFromUsecase,
			expectedErr:  nil,
		},
		{
			name: "Got error when get cart",
			args: args{
				ctx: ctx,
				req: &pbCart.GetUserCartRequest{
					UserId: userID.String(),
				},
			},
			mock: func(usecase *mocks.MockICartUsecase) {
				usecase.EXPECT().GetUserCart(ctx, userID).Return(nil, expectedErrFromUsecase).Times(1)
			},
			expectedCart: nil,
			expectedErr:  status.Errorf(codes.Internal, "Failed to get user cart: %s", expectedErrFromUsecase),
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			cartUsecase := cartHelper(t)
			testcase.mock(cartUsecase)

			actualProduct, actualErr := controller.GetUserCart(
				testcase.args.ctx,
				cartUsecase,
				testcase.args.req,
			)

			assert.Equal(t, testcase.expectedCart, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
