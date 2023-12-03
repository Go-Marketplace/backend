package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	mocks "github.com/Go-Marketplace/backend/product/internal/mocks/repo"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func productHelper(t *testing.T) (*usecase.ProductUsecase, *mocks.MockProductRepo, *mocks.MockDiscountRepo) {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := logger.New("debug")

	productRepo := mocks.NewMockProductRepo(mockCtrl)
	discountRepo := mocks.NewMockDiscountRepo(mockCtrl)
	productUsecase := usecase.NewProductUsecase(productRepo, discountRepo, logger)

	return productUsecase, productRepo, discountRepo
}

func TestGetProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}

	ctx := context.Background()

	productID := uuid.New()

	expectedProductFromRepo := &model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}
	expectedDiscountFromRepo := &model.Discount{
		ProductID: productID,
		Percent:   20,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name            string
		args            args
		mock            func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo)
		expectedProduct *model.Product
		expectedErr     error
	}{
		{
			name: "Successfully get product",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(expectedDiscountFromRepo, nil).Times(1)
			},
			expectedProduct: expectedProductFromRepo,
			expectedErr:     nil,
		},
		{
			name: "Got error when get product",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProduct(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
		{
			name: "Got error when get discount",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			productUsecase, productRepo, discountRepo := productHelper(t)
			testcase.mock(productRepo, discountRepo)

			actualProduct, actualErr := productUsecase.GetProduct(
				testcase.args.ctx,
				testcase.args.productID,
			)

			assert.Equal(t, testcase.expectedProduct, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetProducts(t *testing.T) {
	type args struct {
		ctx          context.Context
		searchParams dto.SearchProductsDTO
	}

	ctx := context.Background()
	productID := uuid.New()

	expectedProductsFromRepo := []*model.Product{
		{
			ID:          productID,
			Name:        "test",
			Description: "test",
		},
	}
	expectedDiscountFromRepo := &model.Discount{
		ProductID: productID,
		Percent:   20,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name             string
		args             args
		mock             func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo)
		expectedProducts []*model.Product
		expectedErr      error
	}{
		{
			name: "Successfully get all products",
			args: args{
				ctx:          ctx,
				searchParams: dto.SearchProductsDTO{},
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProducts(ctx, dto.SearchProductsDTO{}).Return(expectedProductsFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(expectedDiscountFromRepo, nil).Times(1)
			},
			expectedProducts: expectedProductsFromRepo,
			expectedErr:      nil,
		},
		{
			name: "Got error when get products",
			args: args{
				ctx:          ctx,
				searchParams: dto.SearchProductsDTO{},
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProducts(ctx, dto.SearchProductsDTO{}).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProducts: nil,
			expectedErr:      expectedErrFromRepo,
		},
		{
			name: "Got error when get discount",
			args: args{
				ctx:          ctx,
				searchParams: dto.SearchProductsDTO{},
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().GetProducts(ctx, dto.SearchProductsDTO{}).Return(expectedProductsFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProducts: nil,
			expectedErr:      expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			productUsecase, productRepo, discountRepo := productHelper(t)
			testcase.mock(productRepo, discountRepo)

			actualProducts, actualErr := productUsecase.GetProducts(
				testcase.args.ctx,
				testcase.args.searchParams,
			)

			assert.Equal(t, testcase.expectedProducts, actualProducts)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx     context.Context
		product model.Product
	}

	ctx := context.Background()

	productID := uuid.New()
	testProduct := model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(productRepo *mocks.MockProductRepo)
		expectedErr error
	}{
		{
			name: "Successfully create product",
			args: args{
				ctx:     ctx,
				product: testProduct,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().CreateProduct(ctx, testProduct).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when create product",
			args: args{
				ctx:     ctx,
				product: testProduct,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().CreateProduct(ctx, testProduct).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, _ := productHelper(t)
			testcase.mock(productRepo)

			actualErr := productUsecase.CreateProduct(
				testcase.args.ctx,
				testcase.args.product,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx     context.Context
		product model.Product
	}

	ctx := context.Background()

	productID := uuid.New()
	testProduct := model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}

	expectedProductFromRepo := &model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}
	expectedDiscountFromRepo := &model.Discount{
		ProductID: productID,
		Percent:   20,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name            string
		args            args
		mock            func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo)
		expectedProduct *model.Product
		expectedErr     error
	}{
		{
			name: "Successfully update product",
			args: args{
				ctx:     ctx,
				product: testProduct,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().UpdateProduct(ctx, testProduct).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(expectedDiscountFromRepo, nil).Times(1)
			},
			expectedProduct: expectedProductFromRepo,
			expectedErr:     nil,
		},
		{
			name: "Got error when update product",
			args: args{
				ctx:     ctx,
				product: testProduct,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().UpdateProduct(ctx, testProduct).Return(expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
		{
			name: "Got error when get product",
			args: args{
				ctx:     ctx,
				product: testProduct,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				productRepo.EXPECT().UpdateProduct(ctx, testProduct).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, discountRepo := productHelper(t)
			testcase.mock(productRepo, discountRepo)

			actualProduct, actualErr := productUsecase.UpdateProduct(
				testcase.args.ctx,
				testcase.args.product,
			)

			assert.Equal(t, testcase.expectedProduct, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

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
		mock        func(productRepo *mocks.MockProductRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete product",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().DeleteProduct(ctx, productID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete product",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().DeleteProduct(ctx, productID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, _ := productHelper(t)
			testcase.mock(productRepo)

			actualErr := productUsecase.DeleteProduct(
				testcase.args.ctx,
				testcase.args.productID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetCategory(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  int32
	}

	ctx := context.Background()

	var categoryID int32 = 1

	expectedCategoryFromRepo := &model.Category{
		ID:   int32(categoryID),
		Name: "test",
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name             string
		args             args
		mock             func(productRepo *mocks.MockProductRepo)
		expectedCategory *model.Category
		expectedErr      error
	}{
		{
			name: "Successfully get category",
			args: args{
				ctx: ctx,
				id:  categoryID,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().GetCategory(ctx, categoryID).Return(expectedCategoryFromRepo, nil).Times(1)
			},
			expectedCategory: expectedCategoryFromRepo,
			expectedErr:      nil,
		},
		{
			name: "Got error when get category",
			args: args{
				ctx: ctx,
				id:  categoryID,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().GetCategory(ctx, categoryID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedCategory: nil,
			expectedErr:      expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, _ := productHelper(t)
			testcase.mock(productRepo)

			actualCategory, actualErr := productUsecase.GetCategory(
				testcase.args.ctx,
				testcase.args.id,
			)

			assert.Equal(t, testcase.expectedCategory, actualCategory)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetAllCategories(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	ctx := context.Background()

	var categoryID int32 = 1

	expectedCategoryFromRepo := []*model.Category{
		{
			ID:   categoryID,
			Name: "test",
		},
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name               string
		args               args
		mock               func(productRepo *mocks.MockProductRepo)
		expectedCategories []*model.Category
		expectedErr        error
	}{
		{
			name: "Successfully get all categories",
			args: args{
				ctx: ctx,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().GetAllCategories(ctx).Return(expectedCategoryFromRepo, nil).Times(1)
			},
			expectedCategories: expectedCategoryFromRepo,
			expectedErr:        nil,
		},
		{
			name: "Got error when get all categories",
			args: args{
				ctx: ctx,
			},
			mock: func(productRepo *mocks.MockProductRepo) {
				productRepo.EXPECT().GetAllCategories(ctx).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedCategories: nil,
			expectedErr:        expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, _ := productHelper(t)
			testcase.mock(productRepo)

			actualCategory, actualErr := productUsecase.GetAllCategories(
				testcase.args.ctx,
			)

			assert.Equal(t, testcase.expectedCategories, actualCategory)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestCreateDiscount(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		discount model.Discount
	}

	ctx := context.Background()

	productID := uuid.New()
	testDiscount := model.Discount{
		ProductID: productID,
		Percent:   20,
	}

	expectedProductFromRepo := &model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}
	expectedDiscountFromRepo := &model.Discount{
		ProductID: productID,
		Percent:   20,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name            string
		args            args
		mock            func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo)
		expectedProduct *model.Product
		expectedErr     error
	}{
		{
			name: "Successfully create discount",
			args: args{
				ctx:      ctx,
				discount: testDiscount,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().CreateDiscount(ctx, testDiscount).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(expectedDiscountFromRepo, nil).Times(1)
			},
			expectedProduct: expectedProductFromRepo,
			expectedErr:     nil,
		},
		{
			name: "Got error when create discount",
			args: args{
				ctx:      ctx,
				discount: testDiscount,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().CreateDiscount(ctx, testDiscount).Return(expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
		{
			name: "Got error when get product",
			args: args{
				ctx:      ctx,
				discount: testDiscount,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().CreateDiscount(ctx, testDiscount).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, discountRepo := productHelper(t)
			testcase.mock(productRepo, discountRepo)

			actualProduct, actualErr := productUsecase.CreateDiscount(
				testcase.args.ctx,
				testcase.args.discount,
			)

			assert.Equal(t, testcase.expectedProduct, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteDiscount(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx        context.Context
		discountID uuid.UUID
	}

	ctx := context.Background()

	productID := uuid.New()

	expectedProductFromRepo := &model.Product{
		ID:          productID,
		Name:        "test",
		Description: "test",
	}
	expectedDiscountFromRepo := &model.Discount{
		ProductID: productID,
		Percent:   20,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name            string
		args            args
		mock            func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo)
		expectedProduct *model.Product
		expectedErr     error
	}{
		{
			name: "Successfully delete discount",
			args: args{
				ctx:        ctx,
				discountID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().DeleteDiscount(ctx, productID).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(expectedProductFromRepo, nil).Times(1)
				discountRepo.EXPECT().GetDiscount(ctx, productID).Return(expectedDiscountFromRepo, nil).Times(1)
			},
			expectedProduct: expectedProductFromRepo,
			expectedErr:     nil,
		},
		{
			name: "Got error when delete discount",
			args: args{
				ctx:        ctx,
				discountID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().DeleteDiscount(ctx, productID).Return(expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
		{
			name: "Got error when get product",
			args: args{
				ctx:        ctx,
				discountID: productID,
			},
			mock: func(productRepo *mocks.MockProductRepo, discountRepo *mocks.MockDiscountRepo) {
				discountRepo.EXPECT().DeleteDiscount(ctx, productID).Return(nil).Times(1)
				productRepo.EXPECT().GetProduct(ctx, productID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedProduct: nil,
			expectedErr:     expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			productUsecase, productRepo, discountRepo := productHelper(t)
			testcase.mock(productRepo, discountRepo)

			actualProduct, actualErr := productUsecase.DeleteDiscount(
				testcase.args.ctx,
				testcase.args.discountID,
			)

			assert.Equal(t, testcase.expectedProduct, actualProduct)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}
