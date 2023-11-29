package model

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var phoneValidator = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

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
	FirstName string    `json:"first_name" validate:"max=128"`
	LastName  string    `json:"last_name" validate:"max=128"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Address   string    `json:"address" validate:"max=128"`
	Phone     string    `json:"phone"`
	Role      UserRoles `json:"role" validate:"oneof=0 1 2 3"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func (user *User) Validate() error {
	if user.Email != "" {
		if err := validateEmail(user.Email); err != nil {
			return fmt.Errorf("invalid email: %w", err)
		}
	}

	if user.Phone != "" && !phoneValidator.Match([]byte(user.Phone)) {
		return fmt.Errorf("invalid phone number format")
	}

	validate := validator.New()
	return validate.Struct(user)
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
