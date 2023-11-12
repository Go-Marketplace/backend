package model

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

// Represents how the user structure is stored in the database
type User struct {
	ID        uuid.UUID    `json:"user_id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Password  string       `json:"password"`
	Email     mail.Address `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
