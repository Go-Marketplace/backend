package model

import (
	"time"

	"github.com/google/uuid"
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
	ProductIDs      []uuid.UUID       `json:"product_ids"`
	Status          OrderStatus       `json:"status"`
	TotalPrice      int               `json:"total_price"`
	ShippingCost    int               `json:"shipping_cost"`
	DeliveryAddress string            `json:"delivery_address"`
	DeliveryType    OrderDeliveryType `json:"delivery_type"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       time.Time         `json:"deleted_at"`
	CancelledAt     time.Time         `json:"cancelled_at"`
}
