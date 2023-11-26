package usecase_test

import (
	"context"
	"testing"

	mocks "github.com/Go-Marketplace/backend/order/internal/mocks/repo"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func orderHelper(t *testing.T) (*usecase.OrderUseCase, *mocks.MockOrderRepo) {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockOrderRepo(mockCtrl)
	order := usecase.NewOrderUseCase(repo)

	return order, repo
}

func TestGetOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	ctx := context.Background()
	var id uuid.UUID
	var expectedOrder model.Order
	var expectedErr error

	testcases := []struct {
		name          string
		args          args
		mock          func(repo *mocks.MockOrderRepo)
		expectedOrder *model.Order
		expectedErr   error
	}{
		{
			name: "Successfully get order",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrder(ctx, id).Return(&expectedOrder, nil).Times(1)
			},
			expectedOrder: &expectedOrder,
			expectedErr:   nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrder(ctx, id).Return(nil, expectedErr).Times(1)
			},
			expectedOrder: nil,
			expectedErr:   expectedErr,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrder, actualErr := orderUseCase.GetOrder(testcase.args.ctx, testcase.args.id)
			assert.Equal(t, actualOrder, testcase.expectedOrder)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestGetAllUserOrders(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	ctx := context.Background()
	var userID uuid.UUID
	var expectedOrders []*model.Order
	var expectedErr error

	testcases := []struct {
		name           string
		args           args
		mock           func(repo *mocks.MockOrderRepo)
		expectedOrders []*model.Order
		expectedErr    error
	}{
		{
			name: "Successfully get user orders",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetAllUserOrders(ctx, userID).Return(expectedOrders, nil).Times(1)
			},
			expectedOrders: expectedOrders,
			expectedErr:    nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetAllUserOrders(ctx, userID).Return(nil, expectedErr).Times(1)
			},
			expectedOrders: nil,
			expectedErr:    expectedErr,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrder, actualErr := orderUseCase.GetAllUserOrders(testcase.args.ctx, testcase.args.userID)
			assert.Equal(t, actualOrder, testcase.expectedOrders)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx   context.Context
		order model.Order
	}

	ctx := context.Background()
	var order model.Order
	var expectedErr error

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully create order",
			args: args{
				ctx:   ctx,
				order: order,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrder(ctx, order).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx:   ctx,
				order: order,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrder(ctx, order).Return(expectedErr).Times(1)
			},
			expectedErr: expectedErr,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.CreateOrder(testcase.args.ctx, testcase.args.order)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestCancelOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	ctx := context.Background()
	var id uuid.UUID
	var expectedErr error

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully create order",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrder(ctx, id).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrder(ctx, id).Return(expectedErr).Times(1)
			},
			expectedErr: expectedErr,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.DeleteOrder(testcase.args.ctx, testcase.args.id)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}
