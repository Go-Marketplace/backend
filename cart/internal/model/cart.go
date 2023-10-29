package model

import "github.com/google/uuid"

// Represents how the cart structure is stored in the database
type Cart struct {
	ID         uuid.UUID   `json:"cart_id"`
	UserID     uuid.UUID   `json:"user_id"`
	ProductIDs []uuid.UUID `json:"product_ids"`
	TotalPrice int         `json:"total_price"`
}
