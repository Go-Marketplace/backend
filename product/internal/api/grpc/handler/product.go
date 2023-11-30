package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
)

type productRoutes struct {
	pbProduct.UnimplementedProductServer

	productUsecase *usecase.ProductUsecase
	logger         *logger.Logger
}

func NewProductRoutes(productUsecase *usecase.ProductUsecase, logger *logger.Logger) *productRoutes {
	return &productRoutes{
		productUsecase: productUsecase,
		logger:         logger,
	}
}

func (routes *productRoutes) GetProduct(ctx context.Context, req *pbProduct.GetProductRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.GetProduct(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

func (routes *productRoutes) GetAllProducts(ctx context.Context, req *pbProduct.GetAllProductsRequest) (*pbProduct.ProductsResponse, error) {
	products, err := controller.GetAllProducts(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pbProduct.ProductResponse, 0, len(products))
	for _, product := range products {
		protoProducts = append(protoProducts, product.ToProto())
	}

	return &pbProduct.ProductsResponse{
		Products: protoProducts,
	}, nil
}

func (routes *productRoutes) GetAllUserProducts(ctx context.Context, req *pbProduct.GetAllUserProductsRequest) (*pbProduct.ProductsResponse, error) {
	products, err := controller.GetAllUserProducts(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pbProduct.ProductResponse, 0, len(products))
	for _, product := range products {
		protoProducts = append(protoProducts, product.ToProto())
	}

	return &pbProduct.ProductsResponse{
		Products: protoProducts,
	}, nil
}

func (routes *productRoutes) GetAllCategoriesProducts(ctx context.Context, req *pbProduct.GetAllCategoryProductsRequest) (*pbProduct.ProductsResponse, error) {
	products, err := controller.GetAllCategoryProducts(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pbProduct.ProductResponse, 0, len(products))
	for _, product := range products {
		protoProducts = append(protoProducts, product.ToProto())
	}

	return &pbProduct.ProductsResponse{
		Products: protoProducts,
	}, nil
}

func (routes *productRoutes) GetAllCategoryProducts(ctx context.Context, req *pbProduct.GetAllCategoryProductsRequest) (*pbProduct.ProductsResponse, error) {
	products, err := controller.GetAllCategoryProducts(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pbProduct.ProductResponse, 0, len(products))
	for _, product := range products {
		protoProducts = append(protoProducts, product.ToProto())
	}

	return &pbProduct.ProductsResponse{
		Products: protoProducts,
	}, nil
}

func (routes *productRoutes) CreateProduct(ctx context.Context, req *pbProduct.CreateProductRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.CreateProduct(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

func (routes *productRoutes) UpdateProduct(ctx context.Context, req *pbProduct.UpdateProductRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.UpdateProduct(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

func (routes *productRoutes) DeleteProduct(ctx context.Context, req *pbProduct.DeleteProductRequest) (*pbProduct.DeleteProductResponse, error) {
	err := controller.DeleteProduct(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return &pbProduct.DeleteProductResponse{}, nil
}

func (routes *productRoutes) ModerateProduct(ctx context.Context, req *pbProduct.ModerateProductRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.ModerateProduct(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

func (routes *productRoutes) GetCategory(ctx context.Context, req *pbProduct.GetCategoryRequest) (*pbProduct.CategoryResponse, error) {
	category, err := controller.GetCategory(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return category.ToProto(), nil
}

func (routes *productRoutes) GetAllCategories(ctx context.Context, req *pbProduct.GetAllCategoriesRequest) (*pbProduct.CategoriesResponse, error) {
	categories, err := controller.GetAllCategories(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoCategories := make([]*pbProduct.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		protoCategories = append(protoCategories, category.ToProto())
	}

	return &pbProduct.CategoriesResponse{
		Categories: protoCategories,
	}, nil
}

func (routes *productRoutes) CreateDiscount(ctx context.Context, req *pbProduct.CreateDiscountRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.CreateDiscount(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

func (routes *productRoutes) DeleteDiscount(ctx context.Context, req *pbProduct.DeleteDiscountRequest) (*pbProduct.ProductResponse, error) {
	product, err := controller.DeleteDiscount(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}
