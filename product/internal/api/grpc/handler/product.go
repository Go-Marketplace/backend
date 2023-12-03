package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
)

type productRoutes struct {
	pbProduct.UnimplementedProductServer

	productUsecase *usecase.ProductUsecase
	cartClient     pbCart.CartClient
	logger         *logger.Logger
}

func NewProductRoutes(
	productUsecase *usecase.ProductUsecase,
	cartClient pbCart.CartClient,
	logger *logger.Logger,
) *productRoutes {
	return &productRoutes{
		productUsecase: productUsecase,
		cartClient:     cartClient,
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

func (routes *productRoutes) GetProducts(ctx context.Context, req *pbProduct.GetProductsRequest) (*pbProduct.ProductsResponse, error) {
	products, err := controller.GetProducts(ctx, routes.productUsecase, req)
	if err != nil {
		return nil, err
	}

	protoProducts := make([]*pbProduct.ProductResponse, 0, len(products))
	for _, product := range products {
		if product != nil {
			protoProducts = append(protoProducts, product.ToProto())
		} else {
			routes.logger.Warn("Get nil product in GetProducts handler")
		}
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

func (routes *productRoutes) UpdateProducts(ctx context.Context, req *pbProduct.UpdateProductsRequest) (*pbProduct.UpdateProductsResponse, error) {
	if err := controller.UpdateProducts(ctx, routes.productUsecase, req); err != nil {
		return nil, err
	}

	return &pbProduct.UpdateProductsResponse{}, nil
}

func (routes *productRoutes) DeleteProduct(ctx context.Context, req *pbProduct.DeleteProductRequest) (*pbProduct.DeleteProductResponse, error) {
	if err := controller.DeleteProduct(ctx, routes.productUsecase, routes.cartClient, req); err != nil {
		return nil, err
	}

	return &pbProduct.DeleteProductResponse{}, nil
}

func (routes *productRoutes) DeleteUserProducts(ctx context.Context, req *pbProduct.DeleteUserProductsRequest) (*pbProduct.DeleteUserProductsResponse, error) {
	if err := controller.DeleteUserProducts(ctx, routes.productUsecase, routes.cartClient, req); err != nil {
		return nil, err
	}

	return &pbProduct.DeleteUserProductsResponse{}, nil
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
