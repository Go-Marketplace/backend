package model_test

import (
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
				Quantity: 10,
			},
			wasError: false,
		},
		{
			name: "Too big quantity",
			cartline: &model.CartLine{
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
