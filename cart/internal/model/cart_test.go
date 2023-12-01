package model_test

import (
	"strings"
	"testing"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		cartline *model.CartLine
		wasError bool
	}{
		{
			name: "Cartline is valid",
			cartline: &model.CartLine{
				Name:     "test",
				Quantity: 10,
			},
			wasError: false,
		},
		{
			name: "Too long name",
			cartline: &model.CartLine{
				Name:     strings.Repeat("t", 129),
				Quantity: 10,
			},
			wasError: true,
		},
		{
			name: "Too big quantity",
			cartline: &model.CartLine{
				Name:     "test",
				Quantity: 100000000,
			},
			wasError: true,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			actualErr := testcase.cartline.Validate()

			assert.Equal(t, testcase.wasError, actualErr != nil)
		})
	}
}
