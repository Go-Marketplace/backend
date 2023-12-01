package model

import (
	"time"

	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the cart structure is stored in the database
type Cart struct {
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Cartlines []*CartLine
}

func (cart *Cart) ToProto() *pbCart.CartResponse {
	protoCartlines := make([]*pbCart.CartlineResponse, 0, len(cart.Cartlines))
	for _, cartline := range cart.Cartlines {
		protoCartlines = append(protoCartlines, cartline.ToProto())
	}

	return &pbCart.CartResponse{
		UserId:    cart.UserID.String(),
		Cartlines: protoCartlines,
		CreatedAt: timestamppb.New(cart.CreatedAt),
		UpdatedAt: timestamppb.New(cart.UpdatedAt),
	}
}

// Represents one line with a product in a shopping cart in the database
type CartLine struct {
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name" validate:"max=128"`
	Quantity  int64     `json:"quantity" validate:"min=0,max=10000000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cartline *CartLine) Validate() error {
	validate := validator.New()
	return validate.Struct(cartline)
}

func (cartline *CartLine) ToProto() *pbCart.CartlineResponse {
	return &pbCart.CartlineResponse{
		UserId:     cartline.UserID.String(),
		ProductId:  cartline.ProductID.String(),
		Name:       cartline.Name,
		Quantity:   cartline.Quantity,
		CreatedAt:  timestamppb.New(cartline.CreatedAt),
		UpdatedAt:  timestamppb.New(cartline.UpdatedAt),
	}
}
