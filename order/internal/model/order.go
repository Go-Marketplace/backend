package model

import (
	"time"

	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the order structure is stored in the database
type Order struct {
	ID         uuid.UUID `json:"order_id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice int64     `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Orderlines []*Orderline `json:"orderlines"`
}

func (order *Order) ToProto() *pbOrder.OrderResponse {
	var pbOrderlines []*pbOrder.Orderline
	if order.Orderlines != nil {
		pbOrderlines = make([]*pbOrder.Orderline, 0, len(order.Orderlines))
		for _, orderline := range order.Orderlines {
			pbOrderlines = append(pbOrderlines, orderline.ToProto())
		}
	} else {
		pbOrderlines = make([]*pbOrder.Orderline, 0)
	}

	return &pbOrder.OrderResponse{
		OrderId:    order.ID.String(),
		UserId:     order.UserID.String(),
		Orderlines: pbOrderlines,
		TotalPrice: order.TotalPrice,
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

// Represents one line with a product in a shopping cart in the database
type Orderline struct {
	ID        uuid.UUID       `json:"orderline_id"`
	OrderID   uuid.UUID       `json:"order_id"`
	ProductID uuid.UUID       `json:"product_id"`
	Name      string          `json:"name"`
	Price     int64           `json:"price"`
	Quantity  int64           `json:"quantity"`
	Status    OrderlineStatus `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (orderline *Orderline) ToProto() *pbOrder.Orderline {
	return &pbOrder.Orderline{
		OrderlineId: orderline.ID.String(),
		OrderId:     orderline.OrderID.String(),
		ProductId:   orderline.ProductID.String(),
		Name:        orderline.Name,
		Price:       orderline.Price,
		Quantity:    orderline.Quantity,
		Status:      pbOrder.OrderlineStatus(orderline.Status),
		CreatedAt:   timestamppb.New(orderline.CreatedAt),
		UpdatedAt:   timestamppb.New(orderline.UpdatedAt),
	}
}
