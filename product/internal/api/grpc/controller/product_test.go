package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Go-Marketplace/backend/product/internal/api/grpc/controller"
	mocks "github.com/Go-Marketplace/backend/product/internal/mocks/usecase"
	"github.com/Go-Marketplace/backend/product/internal/model"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func productHelper(t *testing.T) *mocks.MockIProductUsecase {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mocks.NewMockIProductUsecase(mockCtrl)
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *pbProduct.GetProductRequest
	}

	ctx := context.Background()
	productID := uuid.New()

	expectedProductFromUsecase := &model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}
	expectedErrFromUsecase := errors.New("test error")

	testcases := []struct {
		name            string
		args            args
		mock            func(usecase *mocks.MockIProductUsecase)
		expectedProduct *model.Product
		expectedErr     error
	}{
		{
			name: "Successfully get product",
			args: args{
				ctx: ctx,
				req: &pbProduct.GetProductRequest{
					ProductId: productID.String(),
				},
			},
			mock: func(usecase *mocks.MockIProductUsecase) {
				usecase.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromUsecase, nil).Times(1)
			},
			expectedProduct: expectedProductFromUsecase,
			expectedErr:     nil,
		},
		{
			name: "Got error when get product",
			args: args{
				ctx: ctx,
				req: &pbProduct.GetProductRequest{
					ProductId: productID.String(),
				},
			},
			mock: func(usecase *mocks.MockIProductUsecase) {
				usecase.EXPECT().GetProduct(ctx, productID).Return(nil, expectedErrFromUsecase).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     status.Errorf(codes.Internal, "Internal error: %s", expectedErrFromUsecase),
		},
		{
			name: "Got error when product is nil",
			args: args{
				ctx: ctx,
				req: &pbProduct.GetProductRequest{
					ProductId: productID.String(),
				},
			},
			mock: func(usecase *mocks.MockIProductUsecase) {
				usecase.EXPECT().GetProduct(ctx, productID).Return(nil, nil).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     status.Errorf(codes.NotFound, "Product not found"),
		},
		{
			name: "Got error when product id is invalid",
			args: args{
				ctx: ctx,
				req: &pbProduct.GetProductRequest{
					ProductId: "123",
				},
			},
			mock: func(usecase *mocks.MockIProductUsecase) {},
			expectedProduct: nil,
			expectedErr:     status.Errorf(codes.InvalidArgument, "Invalid id: invalid UUID length: 3"),
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase := productHelper(t)
			testcase.mock(productUsecase)

			actualProduct, actualErr := controller.GetProduct(
				testcase.args.ctx,
				productUsecase,
				testcase.args.req,
			)

			assert.Equal(t, testcase.expectedProduct, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
