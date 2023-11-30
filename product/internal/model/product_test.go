package model_test

import (
	"strings"
	"testing"

	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateProduct(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		product  *model.Product
		wasError bool
	}{
		{
			name: "Product is valid",
			product: &model.Product{
				Name:        "test",
				Description: "test",
				Price:       100,
				Quantity:    100,
			},
			wasError: false,
		},
		{
			name: "Too long name",
			product: &model.Product{
				Name:        strings.Repeat("t", 129),
				Description: "test",
				Price:       100,
				Quantity:    100,
			},
			wasError: true,
		},
		{
			name: "Too long description",
			product: &model.Product{
				Name:        "test",
				Description: strings.Repeat("t", 1025),
				Price:       100,
				Quantity:    100,
			},
			wasError: true,
		},
		{
			name: "Too big price",
			product: &model.Product{
				Name:        "test",
				Description: "test",
				Price:       10000000000,
				Quantity:    100,
			},
			wasError: true,
		},
		{
			name: "Too big quantity",
			product: &model.Product{
				Name:        "test",
				Description: "test",
				Price:       100,
				Quantity:    100000000,
			},
			wasError: true,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			actualErr := testcase.product.Validate()

			assert.Equal(t, testcase.wasError, actualErr != nil)
		})
	}
}
