package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrderPayloadValidate(t *testing.T) {
	t.Run("valid payload", func(t *testing.T) {
		payload := CreateOrderPayload{
			CouponCode: nil,
			OrderItems: []OrderItem{
				{ProductID: "prod-1", Quantity: 2},
			},
		}
		err := payload.Validate()
		assert.NoError(t, err)
	})

	t.Run("return error if no items in payload", func(t *testing.T) {
		payload := CreateOrderPayload{
			CouponCode: nil,
			OrderItems: []OrderItem{},
		}
		err := payload.Validate()
		assert.Error(t, err)
		assert.Equal(t, "order must have at least one item", err.Error())
	})

	t.Run("return error if empty product ID", func(t *testing.T) {
		payload := CreateOrderPayload{
			OrderItems: []OrderItem{
				{ProductID: "", Quantity: 1},
			},
		}
		err := payload.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "productId is required")
	})

	t.Run("return eror if quantity is zero", func(t *testing.T) {
		payload := CreateOrderPayload{
			OrderItems: []OrderItem{
				{ProductID: "prod-2", Quantity: 0},
			},
		}
		err := payload.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "quantity must be greater than 0")
	})
}
