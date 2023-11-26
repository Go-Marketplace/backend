package model

import (
	"time"

	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the cart structure is stored in the database
type Cart struct {
	ID        uuid.UUID `json:"cart_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Cartlines []*CartLine
}

func (cart *Cart) ToProto() *pbCart.CartModel {
	protoCartlines := make([]*pbCart.CartlineModel, 0, len(cart.Cartlines))
	for _, cartline := range cart.Cartlines {
		protoCartlines = append(protoCartlines, cartline.ToProto())
	}

	return &pbCart.CartModel{
		CartId:    cart.ID.String(),
		UserId:    cart.UserID.String(),
		Cartlines: protoCartlines,
		CreatedAt: timestamppb.New(cart.CreatedAt),
		UpdatedAt: timestamppb.New(cart.UpdatedAt),
	}
}

// Represents one line with a product in a shopping cart in the database
type CartLine struct {
	ID        uuid.UUID `json:"cartline_id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cartline *CartLine) ToProto() *pbCart.CartlineModel {
	return &pbCart.CartlineModel{
		CartlineId: cartline.ID.String(),
		CartId:     cartline.CartID.String(),
		ProductId:  cartline.ProductID.String(),
		Quantity:   cartline.Quantity,
		CreatedAt:  timestamppb.New(cartline.CreatedAt),
		UpdatedAt:  timestamppb.New(cartline.UpdatedAt),
	}
}
