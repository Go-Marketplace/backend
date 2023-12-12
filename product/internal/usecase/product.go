package usecase

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/product/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/google/uuid"
)

type IProductUsecase interface {
	// Product
	GetProducts(ctx context.Context, searchParams dto.SearchProductsDTO) ([]*model.Product, error)
	GetProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, product model.Product) (*model.Product, error)
	UpdateProducts(ctx context.Context, products []model.Product) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
	DeleteUserProducts(ctx context.Context, userID uuid.UUID) error

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
		if product != nil {
			product.Discount, err = usecase.discountRepo.GetDiscount(ctx, product.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (usecase *ProductUsecase) GetProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	product, err := usecase.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	err = usecase.setProductsDiscount(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (usecase *ProductUsecase) GetProducts(ctx context.Context, searchParams dto.SearchProductsDTO) ([]*model.Product, error) {
	products, err := usecase.productRepo.GetProducts(ctx, searchParams)
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

func (usecase *ProductUsecase) UpdateProducts(ctx context.Context, products []model.Product) error {
	return usecase.productRepo.UpdateProducts(ctx, products)
}

func (usecase *ProductUsecase) UpdateProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	if err := usecase.productRepo.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}

	return usecase.GetProduct(ctx, product.ID)
}

func (usecase *ProductUsecase) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	return usecase.productRepo.DeleteProduct(ctx, productID)
}

func (usecase *ProductUsecase) DeleteUserProducts(ctx context.Context, userID uuid.UUID) error {
	return usecase.productRepo.DeleteUserProducts(ctx, userID)
}

func (usecase *ProductUsecase) GetCategory(ctx context.Context, categoryID int32) (*model.Category, error) {
	return usecase.productRepo.GetCategory(ctx, categoryID)
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
