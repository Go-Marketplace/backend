package model

import (
	"time"

	"github.com/google/uuid"
)

// Represents how the cart structure is stored in the database
type Cart struct {
	ID         uuid.UUID `json:"cart_id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Represents one line with a product in a shopping cart in the database
type CartLine struct {
	ID        uuid.UUID `json:"cartline_id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
