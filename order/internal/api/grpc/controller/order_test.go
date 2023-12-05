package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/controller"
	mocks "github.com/Go-Marketplace/backend/order/internal/mocks/usecase"
	"github.com/Go-Marketplace/backend/order/internal/model"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func orderHelper(t *testing.T) *mocks.MockIOrderUsecase {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mocks.NewMockIOrderUsecase(mockCtrl)
}

func TestGetOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *pbOrder.GetOrderRequest
	}

	ctx := context.Background()
	orderID := uuid.New()

	expectedOrderFromUsecase := &model.Order{
		ID: orderID,
	}
	expectedErrFromUsecase := errors.New("test error")

	testcases := []struct {
		name          string
		args          args
		mock          func(usecase *mocks.MockIOrderUsecase)
		expectedOrder *model.Order
		expectedErr   error
	}{
		{
			name: "Successfully get order",
			args: args{
				ctx: ctx,
				req: &pbOrder.GetOrderRequest{
					OrderId: orderID.String(),
				},
			},
			mock: func(usecase *mocks.MockIOrderUsecase) {
				usecase.EXPECT().GetOrder(ctx, orderID).Return(expectedOrderFromUsecase, nil).Times(1)
			},
			expectedOrder: expectedOrderFromUsecase,
			expectedErr:   nil,
		},
		{
			name: "Got error when get order",
			args: args{
				ctx: ctx,
				req: &pbOrder.GetOrderRequest{
					OrderId: orderID.String(),
				},
			},
			mock: func(usecase *mocks.MockIOrderUsecase) {
				usecase.EXPECT().GetOrder(ctx, orderID).Return(nil, expectedErrFromUsecase).Times(1)
			},
			expectedOrder: nil,
			expectedErr:   status.Errorf(codes.Internal, "Failed to get order: %s", expectedErrFromUsecase),
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUsecase := orderHelper(t)
			testcase.mock(orderUsecase)

			actualOrder, actualErr := controller.GetOrder(
				testcase.args.ctx,
				orderUsecase,
				testcase.args.req,
			)

			assert.Equal(t, testcase.expectedOrder, actualOrder)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
