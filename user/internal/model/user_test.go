package model_test

import (
	"strings"
	"testing"

	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		user     *model.User
		wasError bool
	}{
		{
			name: "User is valid",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  "test",
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   "test",
				Phone:     "+79991234567",
				Role:      1,
			},
			wasError: false,
		},
		{
			name: "Too long first name",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: strings.Repeat("t", 129),
				LastName:  "test",
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   "test",
				Phone:     "+79991234567",
				Role:      1,
			},
			wasError: true,
		},
		{
			name: "Too long last name",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  strings.Repeat("t", 129),
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   "test",
				Phone:     "+79991234567",
				Role:      1,
			},
			wasError: true,
		},
		{
			name: "Invalid email",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  "test",
				Password:  "testtest",
				Email:     "testmail@.ru",
				Address:   "test",
				Phone:     "+79991234567",
				Role:      1,
			},
			wasError: true,
		},
		{
			name: "Too long address",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  "test",
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   strings.Repeat("t", 129),
				Phone:     "+79991234567",
				Role:      1,
			},
			wasError: true,
		},
		{
			name: "Invalid phone",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  "test",
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   "test",
				Phone:     "09999123456",
				Role:      1,
			},
			wasError: true,
		},
		{
			name: "Invalid role",
			user: &model.User{
				ID:        uuid.New(),
				FirstName: "test",
				LastName:  "test",
				Password:  "testtest",
				Email:     "test@mail.ru",
				Address:   "test",
				Phone:     "+79991234567",
				Role:      7,
			},
			wasError: true,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			actualErr := testcase.user.Validate()

			assert.Equal(t, testcase.wasError, actualErr != nil)
		})
	}
}
