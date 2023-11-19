package model

import (
	"encoding/json"
	"fmt"
	"time"

	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Discount struct {
	ProductID uuid.UUID `json:"product_id"`
	Percent   float32   `json:"percent"`
	CreatedAt time.Time `json:"created_at"`
	EndedAt   time.Time `json:"ended_at"`
}

func (discount *Discount) ToProto() *pbProduct.DiscountModel {
	return &pbProduct.DiscountModel{
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
