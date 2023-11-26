package model

import (
	"time"

	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRoles int

const (
	Guest UserRoles = iota
	RegisteredUser
	Admin
	SuperAdmin
)

// Represents how the user structure is stored in the database
type User struct {
	ID        uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Role      UserRoles `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) ToProto() *pbUser.UserResponse {
	return &pbUser.UserResponse{
		UserId:    user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		Role:      pbUser.UserRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
