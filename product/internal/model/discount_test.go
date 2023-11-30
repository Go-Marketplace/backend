package model_test

import (
	"testing"

	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateDiscount(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		discount *model.Discount
		wasError bool
	}{
		{
			name: "Discount is valid",
			discount: &model.Discount{
				Percent: 20,
			},
			wasError: false,
		},
		{
			name: "Too big percent",
			discount: &model.Discount{
				Percent: 120,
			},
			wasError: true,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			actualErr := testcase.discount.Validate()

			assert.Equal(t, testcase.wasError, actualErr != nil)
		})
	}
}
