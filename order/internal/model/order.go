package model

import (
	"time"

	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderStatus int

const (
	WaitingForPaymentStatus OrderStatus = iota
	WaitingForDeliveryStatus
	DeliveryStatus
	WaitingForRecieveStatus
	RecievedStatus
)

type OrderDeliveryType int

const (
	SelfDelivery OrderDeliveryType = iota
	CourierDelivery
)

// Represents how the order structure is stored in the database
type Order struct {
	ID              uuid.UUID         `json:"order_id"`
	UserID          uuid.UUID         `json:"user_id"`
	Status          OrderStatus       `json:"status"`
	TotalPrice      int64             `json:"total_price"`
	ShippingCost    int64             `json:"shipping_cost"`
	DeliveryAddress string            `json:"delivery_address"`
	DeliveryType    OrderDeliveryType `json:"delivery_type"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`

	Cartlines []*Cartline `json:"cartlines"`
}

func (order *Order) ToProto() *pbOrder.GetOrderResponse {
	var pbCartlines []*pbOrder.Cartline
	if order.Cartlines != nil {
		pbCartlines = make([]*pbOrder.Cartline, 0, len(order.Cartlines))
		for _, cartline := range order.Cartlines {
			pbCartlines = append(pbCartlines, cartline.ToProto())
		}
	} else {
		pbCartlines = make([]*pbOrder.Cartline, 0)
	}

	return &pbOrder.GetOrderResponse{
		Id:              order.ID.String(),
		UserId:          order.UserID.String(),
		Cartlines:       pbCartlines,
		Status:          pbOrder.OrderStatus(order.Status),
		TotalPrice:      order.TotalPrice,
		ShippingCost:    order.ShippingCost,
		DeliveryAddress: order.DeliveryAddress,
		DeliveryType:    pbOrder.OrderDeliveryType(order.DeliveryType),
		CreatedAt:       timestamppb.New(order.CreatedAt),
		UpdatedAt:       timestamppb.New(order.UpdatedAt),
	}
}

// Represents one line with a product in a shopping cart in the database
type Cartline struct {
	ID       uuid.UUID `json:"cartline_id"`
	OrderID  uuid.UUID `json:"order_id"`
	Quantity int64     `json:"quantity"`

	Product *Product `json:"product"`
}

func (cartline *Cartline) ToProto() *pbOrder.Cartline {
	var pbProduct *pbOrder.Product
	if cartline.Product != nil {
		pbProduct = cartline.Product.ToProto()
	}

	return &pbOrder.Cartline{
		Id:       cartline.ID.String(),
		OrderId:  cartline.OrderID.String(),
		Product:  pbProduct,
		Quantity: cartline.Quantity,
	}
}

// Represents a product in a shopping cart in the database
type Product struct {
	ID          uuid.UUID `json:"product_id"`
	CartlineID  uuid.UUID `json:"cartline_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
}

func (product *Product) ToProto() *pbOrder.Product {
	return &pbOrder.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}
