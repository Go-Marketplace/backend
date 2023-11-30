package model

import (
	"encoding/json"
	"fmt"
	"time"

	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Discount struct {
	ProductID uuid.UUID `json:"product_id"`
	Percent   float32   `json:"percent" validate:"required,min=0,max=100"`
	CreatedAt time.Time `json:"created_at"`
	EndedAt   time.Time `json:"ended_at"`
}

func (discount *Discount) Validate() error {
	if discount.EndedAt.Before(discount.CreatedAt) {
		return fmt.Errorf("ended_at cannot be less than created_at")
	}

	validate := validator.New()
	return validate.Struct(discount)
}

func (discount *Discount) ToProto() *pbProduct.DiscountResponse {
	return &pbProduct.DiscountResponse{
		ProductId: discount.ProductID.String(),
		Percent:   discount.Percent,
		CreatedAt: timestamppb.New(discount.CreatedAt),
		EndedAt:   timestamppb.New(discount.EndedAt),
	}
}

func (discount *Discount) MarshalBinary() ([]byte, error) {
	return json.Marshal(discount)
}

func (discount *Discount) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &discount); err != nil {
		return fmt.Errorf("cannot unmarshal binary to discount: %w", err)
	}

	return nil
}
