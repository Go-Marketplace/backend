package model_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Go-Marketplace/backend/order/internal/model"
)

func TestValidateOrderline(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name      string
		orderline *model.Orderline
		wasError  bool
	}{
		{
			name: "Order is valid",
			orderline: &model.Orderline{
				Name:     "test",
				Price:    100,
				Quantity: 100,
				Status:   1,
			},
			wasError: false,
		},
		{
			name: "Too long name",
			orderline: &model.Orderline{
				Name:     strings.Repeat("t", 129),
				Price:    100,
				Quantity: 100,
				Status:   1,
			},
			wasError: true,
		},
		{
			name: "Too big price",
			orderline: &model.Orderline{
				Name:     "test",
				Price:    10000000000,
				Quantity: 100,
				Status:   1,
			},
			wasError: true,
		},
		{
			name: "Too big quantity",
			orderline: &model.Orderline{
				Name:     "test",
				Price:    100,
				Quantity: 100000000,
				Status:   1,
			},
			wasError: true,
		},
		{
			name: "Invalid status",
			orderline: &model.Orderline{
				Name:     "test",
				Price:    100,
				Quantity: 100,
				Status:   5,
			},
			wasError: true,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			actualErr := testcase.orderline.Validate()

			assert.Equal(t, testcase.wasError, actualErr != nil)
		})
	}
}
