package handler

import (
	"errors"
	"strings"

	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	"github.com/srinivasaleti/quickbite/server/pkg/price"
)

func ToPtr[T any](v T) *T {
	return &v
}

func getTotalPrice(order ordermodel.Order, isValidCoupon bool) (price.Cent, error) {
	total := totalPriceInCents(order)
	if order.CouponCode == nil || !isValidCoupon {
		return total, nil
	}
	code := strings.ToUpper(*order.CouponCode)

	switch code {
	// As per https://github.com/oolio-group/kart-challenge
	// HAPPYHOURS applies 18% discount to the order total
	case "HAPPYHOURS":
		return price.Cent(total).Subtract(price.Cent(total).Percentize(18)), nil
	// BUYGETONE gives the lowest priced item for free
	case "BUYGETONE":
		if len(order.OrderItems) < 2 {
			return total, errors.New("add atleast 2 items to apply the coupon")
		}
		lowest := findLowestUnitPrice(order.OrderItems)
		return price.Cent(total).Subtract(lowest), nil
	default:
		// There is not mention about rest of the coupons. So going with a 10% discusount for now.
		return price.Cent(total).Subtract(price.Cent(total).Percentize(10)), nil
	}
}

func totalPriceInCents(payload ordermodel.Order) price.Cent {
	total := price.Cent(0)
	for _, item := range payload.OrderItems {
		total = total.Add(item.PriceInCents.Multiply(item.Quantity))
	}
	return total
}

func findLowestUnitPrice(items []ordermodel.OrderItem) price.Cent {
	lowest := items[0].PriceInCents

	for _, item := range items {
		if item.PriceInCents < lowest {
			lowest = item.PriceInCents
		}
	}

	return lowest
}
