package model

import (
	"time"

	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the product structure is stored in the database
type Product struct {
	ID          uuid.UUID `json:"product_id"`
	UserID      uuid.UUID `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
	Moderated   bool      `json:"moderated"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Discount *Discount
}

func (product *Product) ToProto() *pbProduct.ProductModel {
	var discount *pbProduct.DiscountModel
	if product.Discount != nil {
		discount = product.Discount.ToProto()
	}

	return &pbProduct.ProductModel{
		ProductId:   product.ID.String(),
		UserId:      product.UserID.String(),
		CategoryId:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Moderated:   product.Moderated,
		Discount:    discount,
		CreatedAt:   timestamppb.New(product.CreatedAt),
		UpdatedAt:   timestamppb.New(product.UpdatedAt),
	}
}

type Category struct {
	ID          int32  `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (category *Category) ToProto() *pbProduct.CategoryModel {
	return &pbProduct.CategoryModel{
		CategoryId:  category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
