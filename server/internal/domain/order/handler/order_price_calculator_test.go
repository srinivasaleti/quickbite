package handler

import (
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	"github.com/srinivasaleti/quickbite/server/pkg/price"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalPrice(t *testing.T) {
	items := []model.OrderItem{
		{ProductID: "p1", PriceInCents: price.Cent(650), Quantity: 1},
		{ProductID: "p2", PriceInCents: price.Cent(700), Quantity: 10},
		{ProductID: "p3", PriceInCents: price.Cent(800), Quantity: 3},
	}

	t.Run("No coupon", func(t *testing.T) {
		order := ordermodel.CreateOrderPayload{
			OrderItems: items,
			CouponCode: nil,
		}
		total, err := getTotalPrice(order, true)
		assert.NoError(t, err)
		assert.Equal(t, price.Cent(10050), total)
	})

	t.Run("No discount if coupon not valid", func(t *testing.T) {
		coupon := "HAPPYHRS"
		order := ordermodel.CreateOrderPayload{
			OrderItems: items,
			CouponCode: &coupon,
		}
		total, err := getTotalPrice(order, false)
		assert.NoError(t, err)
		assert.Equal(t, price.Cent(10050), total)
	})

	t.Run("HAPPYHOURS coupon - 18% discount", func(t *testing.T) {
		coupon := "HAPPYHOURS"
		order := ordermodel.CreateOrderPayload{
			OrderItems: items,
			CouponCode: &coupon,
		}
		total, err := getTotalPrice(order, true)
		assert.NoError(t, err)
		expectedDiscount := price.Cent(10050).Percentize(18)
		expectedTotal := price.Cent(10050).Subtract(expectedDiscount)
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("BUYGETONE coupon - lowest priced item free", func(t *testing.T) {
		coupon := "BUYGETONE"
		order := ordermodel.CreateOrderPayload{
			OrderItems: items,
			CouponCode: &coupon,
		}
		total, err := getTotalPrice(order, true)
		assert.NoError(t, err)
		lowest := findLowestUnitPrice(items)
		expectedTotal := price.Cent(10050).Subtract(lowest)
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("BUYGETONE coupon with less than 2 items should error", func(t *testing.T) {
		coupon := "BUYGETONE"
		order := ordermodel.CreateOrderPayload{
			OrderItems: []model.OrderItem{
				{ProductID: "p1", PriceInCents: price.Cent(650), Quantity: 1},
			},
			CouponCode: &coupon,
		}
		total, err := getTotalPrice(order, true)
		assert.Error(t, err)
		assert.EqualError(t, err, "add atleast 2 items to apply the coupon")
		assert.Equal(t, price.Cent(650), total)
	})

	t.Run("123456 coupon - 10% discount", func(t *testing.T) {
		coupon := "123456"
		order := ordermodel.CreateOrderPayload{
			OrderItems: items,
			CouponCode: &coupon,
		}
		total, err := getTotalPrice(order, true)
		assert.NoError(t, err)
		assert.Equal(t, price.Cent(9045), total)
	})
}
