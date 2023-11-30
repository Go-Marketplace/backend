package usecase

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/product/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
)

type IProductUsecase interface {
	// Product
	GetProduct(ctx context.Context, id uuid.UUID) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]*model.Product, error)
	GetAllUserProducts(ctx context.Context, userID uuid.UUID) ([]*model.Product, error)
	GetAllCategoryProducts(ctx context.Context, categoryID int32) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, product model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ModerateProduct(ctx context.Context, product model.Product) (*model.Product, error)

	// Category
	GetCategory(ctx context.Context, id int32) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]*model.Category, error)

	// Discount
	CreateDiscount(ctx context.Context, discount model.Discount) (*model.Product, error)
	DeleteDiscount(ctx context.Context, productID uuid.UUID) (*model.Product, error)
}

type ProductUsecase struct {
	productRepo  interfaces.ProductRepo
	discountRepo interfaces.DiscountRepo
	logger       *logger.Logger
}

func NewProductUsecase(productRepo interfaces.ProductRepo, discountRepo interfaces.DiscountRepo, logger *logger.Logger) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  productRepo,
		discountRepo: discountRepo,
		logger:       logger,
	}
}

func (usecase *ProductUsecase) setProductsDiscount(ctx context.Context, products ...*model.Product) error {
	var err error
	for _, product := range products {
		product.Discount, err = usecase.discountRepo.GetDiscount(ctx, product.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (usecase *ProductUsecase) GetProduct(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	product, err := usecase.productRepo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	err = usecase.setProductsDiscount(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (usecase *ProductUsecase) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := usecase.productRepo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	err = usecase.setProductsDiscount(ctx, products...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (usecase *ProductUsecase) GetAllUserProducts(ctx context.Context, userID uuid.UUID) ([]*model.Product, error) {
	products, err := usecase.productRepo.GetAllUserProducts(ctx, userID)
	if err != nil {
		return nil, err
	}

	err = usecase.setProductsDiscount(ctx, products...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (usecase *ProductUsecase) GetAllCategoryProducts(ctx context.Context, categoryID int32) ([]*model.Product, error) {
	products, err := usecase.productRepo.GetAllCategoryProducts(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	err = usecase.setProductsDiscount(ctx, products...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (usecase *ProductUsecase) CreateProduct(ctx context.Context, product model.Product) error {
	return usecase.productRepo.CreateProduct(ctx, product)
}

func (usecase *ProductUsecase) UpdateProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	if err := usecase.productRepo.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}

	return usecase.GetProduct(ctx, product.ID)
}

func (usecase *ProductUsecase) ModerateProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	if err := usecase.productRepo.ModerateProduct(ctx, product); err != nil {
		return nil, err
	}

	return usecase.GetProduct(ctx, product.ID)
}

func (usecase *ProductUsecase) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return usecase.productRepo.DeleteProduct(ctx, id)
}

func (usecase *ProductUsecase) GetCategory(ctx context.Context, id int32) (*model.Category, error) {
	return usecase.productRepo.GetCategory(ctx, id)
}

func (usecase *ProductUsecase) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	return usecase.productRepo.GetAllCategories(ctx)
}

func (usecase *ProductUsecase) CreateDiscount(ctx context.Context, discount model.Discount) (*model.Product, error) {
	if err := usecase.discountRepo.CreateDiscount(ctx, discount); err != nil {
		return nil, err
	}

	return usecase.GetProduct(ctx, discount.ProductID)
}

func (usecase *ProductUsecase) DeleteDiscount(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	if err := usecase.discountRepo.DeleteDiscount(ctx, productID); err != nil {
		return nil, err
	}

	return usecase.GetProduct(ctx, productID)
}
