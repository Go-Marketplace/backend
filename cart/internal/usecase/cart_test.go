package usecase_test

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/Go-Marketplace/backend/cart/internal/mocks/repo"
	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func cartHelper(t *testing.T) (*usecase.CartUsecase, *mocks.MockCartRepo, *mocks.MockCartTaskRepo) {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := logger.New("debug")

	cartRepo := mocks.NewMockCartRepo(mockCtrl)
	cartTaskRepo := mocks.NewMockCartTaskRepo(mockCtrl)
	cartUsecase := usecase.NewCartUsecase(cartRepo, cartTaskRepo, logger)

	return cartUsecase, cartRepo, cartTaskRepo
}

func TestGetUserCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	ctx := context.Background()

	userID := uuid.New()

	expectedCartFromRepo := &model.Cart{
		UserID: userID,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(cartRepo *mocks.MockCartRepo)
		expectedCart *model.Cart
		expectedErr  error
	}{
		{
			name: "Successfully get user cart",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().GetUserCart(ctx, userID).Return(expectedCartFromRepo, nil).Times(1)
			},
			expectedCart: expectedCartFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Got error when get user cart",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().GetUserCart(ctx, userID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedCart: nil,
			expectedErr:  expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualCart, actualErr := cartUsecase.GetUserCart(
				testcase.args.ctx,
				testcase.args.userID,
			)

			assert.Equal(t, testcase.expectedCart, actualCart)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestCreateCart(t *testing.T) {
	type args struct {
		ctx  context.Context
		cart model.Cart
	}

	ctx := context.Background()

	userID := uuid.New()

	testCart := model.Cart{
		UserID: userID,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(cartRepo *mocks.MockCartRepo, cartTaskRepo *mocks.MockCartTaskRepo)
		expectedErr error
	}{
		{
			name: "Successfully create new cart",
			args: args{
				ctx:  ctx,
				cart: testCart,
			},
			mock: func(cartRepo *mocks.MockCartRepo, cartTaskRepo *mocks.MockCartTaskRepo) {
				cartRepo.EXPECT().CreateCart(ctx, testCart).Return(nil).Times(1)
				cartTaskRepo.EXPECT().CreateCartTask(ctx, gomock.Any()).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when create cart",
			args: args{
				ctx:  ctx,
				cart: testCart,
			},
			mock: func(cartRepo *mocks.MockCartRepo, cartTaskRepo *mocks.MockCartTaskRepo) {
				cartRepo.EXPECT().CreateCart(ctx, testCart).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
		{
			name: "Got error when create cart task",
			args: args{
				ctx:  ctx,
				cart: testCart,
			},
			mock: func(cartRepo *mocks.MockCartRepo, cartTaskRepo *mocks.MockCartTaskRepo) {
				cartRepo.EXPECT().CreateCart(ctx, testCart).Return(nil).Times(1)
				cartTaskRepo.EXPECT().CreateCartTask(ctx, gomock.Any()).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, cartTaskRepo := cartHelper(t)
			testcase.mock(cartRepo, cartTaskRepo)

			actualErr := cartUsecase.CreateCart(
				testcase.args.ctx,
				testcase.args.cart,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetCartline(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    uuid.UUID
		productID uuid.UUID
	}

	ctx := context.Background()

	userID := uuid.New()
	productID := uuid.New()

	expectedCartlineFromRepo := &model.CartLine{
		UserID:    userID,
		ProductID: productID,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name             string
		args             args
		mock             func(cartRepo *mocks.MockCartRepo)
		expectedCartline *model.CartLine
		expectedErr      error
	}{
		{
			name: "Successfully get cartline",
			args: args{
				ctx:       ctx,
				userID:    userID,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().GetCartline(ctx, userID, productID).Return(expectedCartlineFromRepo, nil).Times(1)
			},
			expectedCartline: expectedCartlineFromRepo,
			expectedErr:      nil,
		},
		{
			name: "Got error when get cartline",
			args: args{
				ctx:       ctx,
				userID:    userID,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().GetCartline(ctx, userID, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedCartline: nil,
			expectedErr:      expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualCartline, actualErr := cartUsecase.GetCartline(
				testcase.args.ctx,
				testcase.args.userID,
				testcase.args.productID,
			)

			assert.Equal(t, testcase.expectedCartline, actualCartline)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestCreateCartline(t *testing.T) {
	type args struct {
		ctx      context.Context
		cartline *model.CartLine
	}

	ctx := context.Background()

	userID := uuid.New()
	productID := uuid.New()

	testCartline := &model.CartLine{
		UserID:    userID,
		ProductID: productID,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(cartRepo *mocks.MockCartRepo)
		expectedErr error
	}{
		{
			name: "Successfully create new cartline",
			args: args{
				ctx:      ctx,
				cartline: testCartline,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().CreateCartline(ctx, testCartline).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when create cartline",
			args: args{
				ctx:      ctx,
				cartline: testCartline,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().CreateCartline(ctx, testCartline).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualErr := cartUsecase.CreateCartline(
				testcase.args.ctx,
				testcase.args.cartline,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestUpdateCartline(t *testing.T) {
	type args struct {
		ctx      context.Context
		cartline model.CartLine
	}

	ctx := context.Background()

	userID := uuid.New()
	productID := uuid.New()

	testCartline := model.CartLine{
		UserID:    userID,
		ProductID: productID,
	}

	expectedCartlineFromRepo := &testCartline
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name                     string
		args                     args
		mock                     func(cartRepo *mocks.MockCartRepo)
		expectedCartlineFromRepo *model.CartLine
		expectedErr              error
	}{
		{
			name: "Successfully update new cartline",
			args: args{
				ctx:      ctx,
				cartline: testCartline,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().UpdateCartline(ctx, testCartline).Return(nil).Times(1)
				cartRepo.EXPECT().GetCartline(ctx, userID, productID).Return(expectedCartlineFromRepo, nil)
			},
			expectedCartlineFromRepo: expectedCartlineFromRepo,
			expectedErr:              nil,
		},
		{
			name: "Got error when update cartline",
			args: args{
				ctx:      ctx,
				cartline: testCartline,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().UpdateCartline(ctx, testCartline).Return(expectedErrFromRepo).Times(1)
			},
			expectedCartlineFromRepo: nil,
			expectedErr:              expectedErrFromRepo,
		},
		{
			name: "Got error when get cartline",
			args: args{
				ctx:      ctx,
				cartline: testCartline,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().UpdateCartline(ctx, testCartline).Return(expectedErrFromRepo).Times(1)
				cartRepo.EXPECT().GetCartline(ctx, userID, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedCartlineFromRepo: nil,
			expectedErr:              expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualCartline, actualErr := cartUsecase.UpdateCartline(
				testcase.args.ctx,
				testcase.args.cartline,
			)

			assert.Equal(t, testcase.expectedCartlineFromRepo, actualCartline)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteCart(t *testing.T) {
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
		mock        func(cartRepo *mocks.MockCartRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete cart",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCart(ctx, userID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete cart",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCart(ctx, userID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualErr := cartUsecase.DeleteCart(
				testcase.args.ctx,
				testcase.args.userID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteCartline(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    uuid.UUID
		productID uuid.UUID
	}

	ctx := context.Background()

	userID := uuid.New()
	productID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(cartRepo *mocks.MockCartRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete cartline",
			args: args{
				ctx:       ctx,
				userID:    userID,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCartline(ctx, userID, productID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete cartline",
			args: args{
				ctx:       ctx,
				userID:    userID,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCartline(ctx, userID, productID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualErr := cartUsecase.DeleteCartline(
				testcase.args.ctx,
				testcase.args.userID,
				testcase.args.productID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteProductCartlines(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}

	ctx := context.Background()

	productID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(cartRepo *mocks.MockCartRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete product cartlines",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteProductCartlines(ctx, productID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete product cartlines",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteProductCartlines(ctx, productID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualErr := cartUsecase.DeleteProductCartlines(
				testcase.args.ctx,
				testcase.args.productID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteCartCartlines(t *testing.T) {
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
		mock        func(cartRepo *mocks.MockCartRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete cart cartlines",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCartCartlines(ctx, userID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete cart cartlines",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(cartRepo *mocks.MockCartRepo) {
				cartRepo.EXPECT().DeleteCartCartlines(ctx, userID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			cartUsecase, cartRepo, _ := cartHelper(t)
			testcase.mock(cartRepo)

			actualErr := cartUsecase.DeleteCartCartlines(
				testcase.args.ctx,
				testcase.args.userID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
