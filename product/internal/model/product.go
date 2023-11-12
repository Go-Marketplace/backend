package model

import (
	"time"

	"github.com/google/uuid"
)

// Represents how the product structure is stored in the database
type Product struct {
	ID          uuid.UUID `json:"product_id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Weight      int       `json:"weight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
