package model

import (
	"time"

	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the order structure is stored in the database
type Order struct {
	ID        uuid.UUID `json:"order_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Orderlines []*Orderline `json:"orderlines"`
}

func (order *Order) ToProto() *pbOrder.OrderResponse {
	var pbOrderlines []*pbOrder.OrderlineResponse
	if order.Orderlines != nil {
		pbOrderlines = make([]*pbOrder.OrderlineResponse, 0, len(order.Orderlines))
		for _, orderline := range order.Orderlines {
			pbOrderlines = append(pbOrderlines, orderline.ToProto())
		}
	} else {
		pbOrderlines = make([]*pbOrder.OrderlineResponse, 0)
	}

	return &pbOrder.OrderResponse{
		OrderId:    order.ID.String(),
		UserId:     order.UserID.String(),
		Orderlines: pbOrderlines,
		CreatedAt:  timestamppb.New(order.CreatedAt),
		UpdatedAt:  timestamppb.New(order.UpdatedAt),
	}
}

type OrderlineStatus int32

const (
	Canceled OrderlineStatus = iota
	PendingPayment
	Delivery
	Recieved
)

// Represents one line with a product in a order in the database
type Orderline struct {
	OrderID   uuid.UUID       `json:"order_id"`
	ProductID uuid.UUID       `json:"product_id"`
	Name      string          `json:"name" validate:"max=128"`
	Price     int64           `json:"price" validate:"min=0,max=1000000000"`
	Quantity  int64           `json:"quantity" validate:"min=0,max=10000000"`
	Status    OrderlineStatus `json:"status" validate:"min=0,max=3"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (orderline *Orderline) Validate() error {
	validate := validator.New()
	return validate.Struct(orderline)
}

func (orderline *Orderline) ToProto() *pbOrder.OrderlineResponse {
	return &pbOrder.OrderlineResponse{
		OrderId:   orderline.OrderID.String(),
		ProductId: orderline.ProductID.String(),
		Name:      orderline.Name,
		Price:     orderline.Price,
		Quantity:  orderline.Quantity,
		Status:    pbOrder.OrderlineStatus(orderline.Status),
		CreatedAt: timestamppb.New(orderline.CreatedAt),
		UpdatedAt: timestamppb.New(orderline.UpdatedAt),
	}
}
