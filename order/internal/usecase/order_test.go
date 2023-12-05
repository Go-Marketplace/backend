package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
	mocks "github.com/Go-Marketplace/backend/order/internal/mocks/repo"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func orderHelper(t *testing.T) (*usecase.OrderUsecase, *mocks.MockOrderRepo) {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockOrderRepo(mockCtrl)
	order := usecase.NewOrderUsecase(repo)

	return order, repo
}

func TestGetOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	ctx := context.Background()
	orderID := uuid.New()

	expectedOrderFromRepo := &model.Order{
		ID: orderID,
	}
	expectedErrFromRepo := errors.New("test error")

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
				id:  orderID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrder(ctx, orderID).Return(expectedOrderFromRepo, nil).Times(1)
			},
			expectedOrder: expectedOrderFromRepo,
			expectedErr:   nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx: ctx,
				id:  orderID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrder(ctx, orderID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedOrder: nil,
			expectedErr:   expectedErrFromRepo,
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

func TestGetOrders(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx          context.Context
		searchParams dto.SearchOrderDTO
	}

	ctx := context.Background()
	orderID := uuid.New()
	searchParams := dto.SearchOrderDTO{}

	expectedOrdersFromRepo := []*model.Order{
		{
			ID: orderID,
		},
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name           string
		args           args
		mock           func(repo *mocks.MockOrderRepo)
		expectedOrders []*model.Order
		expectedErr    error
	}{
		{
			name: "Successfully get orders",
			args: args{
				ctx:          ctx,
				searchParams: searchParams,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrders(ctx, searchParams).Return(expectedOrdersFromRepo, nil).Times(1)
			},
			expectedOrders: expectedOrdersFromRepo,
			expectedErr:    nil,
		},
		{
			name: "Get repo error",
			args: args{
				ctx:          ctx,
				searchParams: searchParams,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrders(ctx, searchParams).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedOrders: nil,
			expectedErr:    expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrder, actualErr := orderUseCase.GetOrders(testcase.args.ctx, testcase.args.searchParams)
			assert.Equal(t, actualOrder, testcase.expectedOrders)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx   context.Context
		order *model.Order
	}

	ctx := context.Background()
	orderID := uuid.New()
	testOrder := &model.Order{
		ID: orderID,
	}

	expectedOrderFromRepo := &model.Order{
		ID: orderID,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name          string
		args          args
		mock          func(repo *mocks.MockOrderRepo)
		expectedOrder *model.Order
		expectedErr   error
	}{
		{
			name: "Successfully create order",
			args: args{
				ctx:   ctx,
				order: testOrder,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrder(ctx, testOrder).Return(nil).Times(1)
				repo.EXPECT().GetOrder(ctx, orderID).Return(expectedOrderFromRepo, nil).Times(1)
			},
			expectedOrder: expectedOrderFromRepo,
			expectedErr:   nil,
		},
		{
			name: "Get error when create order",
			args: args{
				ctx:   ctx,
				order: testOrder,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrder(ctx, testOrder).Return(expectedErrFromRepo).Times(1)
			},
			expectedOrder: nil,
			expectedErr:   expectedErrFromRepo,
		},
		{
			name: "Get error when get order",
			args: args{
				ctx:   ctx,
				order: testOrder,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrder(ctx, testOrder).Return(nil).Times(1)
				repo.EXPECT().GetOrder(ctx, orderID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedOrder: nil,
			expectedErr:   expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrder, actualErr := orderUseCase.CreateOrder(testcase.args.ctx, testcase.args.order)

			assert.Equal(t, testcase.expectedOrder, actualOrder)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx     context.Context
		orderID uuid.UUID
	}

	ctx := context.Background()
	orderID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete order",
			args: args{
				ctx:     ctx,
				orderID: orderID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrder(ctx, orderID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get error when delete order",
			args: args{
				ctx:     ctx,
				orderID: orderID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrder(ctx, orderID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.DeleteOrder(testcase.args.ctx, testcase.args.orderID)

			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestDeleteUserOrders(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	ctx := context.Background()
	userID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete user orders",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteUserOrders(ctx, userID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get error when delete user orders",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteUserOrders(ctx, userID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.DeleteUserOrders(testcase.args.ctx, testcase.args.userID)

			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestGetOrderline(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context
		orderID   uuid.UUID
		productID uuid.UUID
	}

	ctx := context.Background()
	orderID := uuid.New()
	productID := uuid.New()

	expectedOrderlineFromRepo := &model.Orderline{
		OrderID:   orderID,
		ProductID: productID,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name              string
		args              args
		mock              func(repo *mocks.MockOrderRepo)
		expectedOrderline *model.Orderline
		expectedErr       error
	}{
		{
			name: "Successfully get orderline",
			args: args{
				ctx:       ctx,
				orderID:   orderID,
				productID: productID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrderline(ctx, orderID, productID).Return(expectedOrderlineFromRepo, nil).Times(1)
			},
			expectedOrderline: expectedOrderlineFromRepo,
			expectedErr:       nil,
		},
		{
			name: "Get error when get orderline",
			args: args{
				ctx:       ctx,
				orderID:   orderID,
				productID: productID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().GetOrderline(ctx, orderID, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedOrderline: nil,
			expectedErr:       expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrderline, actualErr := orderUseCase.GetOrderline(
				testcase.args.ctx,
				testcase.args.orderID,
				testcase.args.productID,
			)

			assert.Equal(t, actualOrderline, testcase.expectedOrderline)
			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}

func TestCreateOrderline(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context
		orderline *model.Orderline
	}

	ctx := context.Background()
	orderID := uuid.New()
	testOrderline := &model.Orderline{
		OrderID: orderID,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully create orderline",
			args: args{
				ctx:       ctx,
				orderline: testOrderline,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrderline(ctx, testOrderline).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get error when create orderline",
			args: args{
				ctx:       ctx,
				orderline: testOrderline,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().CreateOrderline(ctx, testOrderline).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.CreateOrderline(testcase.args.ctx, testcase.args.orderline)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestUpdateOrderline(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context
		orderline *model.Orderline
	}

	ctx := context.Background()
	orderID := uuid.New()
	productID := uuid.New()
	testOrderline := &model.Orderline{
		OrderID:   orderID,
		ProductID: productID,
	}

	expectedOrderlineFromRepo := &model.Orderline{
		OrderID:   orderID,
		ProductID: productID,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name              string
		args              args
		mock              func(repo *mocks.MockOrderRepo)
		expectedOrderline *model.Orderline
		expectedErr       error
	}{
		{
			name: "Successfully update orderline",
			args: args{
				ctx:       ctx,
				orderline: testOrderline,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().UpdateOrderline(ctx, testOrderline).Return(nil).Times(1)
				repo.EXPECT().GetOrderline(ctx, orderID, productID).Return(expectedOrderlineFromRepo, nil).Times(1)
			},
			expectedOrderline: expectedOrderlineFromRepo,
			expectedErr:       nil,
		},
		{
			name: "Get error when update orderline",
			args: args{
				ctx:       ctx,
				orderline: testOrderline,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().UpdateOrderline(ctx, testOrderline).Return(expectedErrFromRepo).Times(1)
			},
			expectedOrderline: nil,
			expectedErr:       expectedErrFromRepo,
		},
		{
			name: "Get error when get orderline",
			args: args{
				ctx:       ctx,
				orderline: testOrderline,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().UpdateOrderline(ctx, testOrderline).Return(nil).Times(1)
				repo.EXPECT().GetOrderline(ctx, orderID, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedOrderline: nil,
			expectedErr:       expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualOrderline, actualErr := orderUseCase.UpdateOrderline(testcase.args.ctx, testcase.args.orderline)

			assert.Equal(t, testcase.expectedOrderline, actualOrderline)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteOrderline(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context
		orderID   uuid.UUID
		productID uuid.UUID
	}

	ctx := context.Background()
	orderID := uuid.New()
	productID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockOrderRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete orderline",
			args: args{
				ctx:       ctx,
				orderID:   orderID,
				productID: productID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrderline(ctx, orderID, productID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Get error when delete orderline",
			args: args{
				ctx:       ctx,
				orderID:   orderID,
				productID: productID,
			},
			mock: func(repo *mocks.MockOrderRepo) {
				repo.EXPECT().DeleteOrderline(ctx, orderID, productID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			orderUseCase, orderRepo := orderHelper(t)
			testcase.mock(orderRepo)

			actualErr := orderUseCase.DeleteOrderline(
				testcase.args.ctx,
				testcase.args.orderID,
				testcase.args.productID,
			)

			assert.Equal(t, actualErr, testcase.expectedErr)
		})
	}
}
