package model

import (
	"time"

	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Represents how the user structure is stored in the database
type User struct {
	ID        uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) ToProto() *pbUser.UserResponse {
	return &pbUser.UserResponse{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func FromProtoToUser(protoUser *pbUser.UserRequest) (*User, error) {
	return &User{
		ID:        uuid.New(),
		FirstName: protoUser.FirstName,
		LastName:  protoUser.LastName,
		Email:     protoUser.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
