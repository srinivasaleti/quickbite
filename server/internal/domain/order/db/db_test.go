package db

import (
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
	productmodel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/stretchr/testify/assert"
)

var products = []productmodel.Product{
	{Name: "Product1", ExternalID: ToPtr("p-1"), PriceInCents: 1000, ImageURL: "url1"},
	{Name: "Product2", ExternalID: ToPtr("p-2"), PriceInCents: 2000, ImageURL: "url2"},
}

func TestInsertOrder(t *testing.T) {
	testContainer, err := database.SetupTestDatabase()
	assert.NoError(t, err)

	productdb := productdb.ProductDB{DB: testContainer.DB}
	orderDb := OrderDB{DB: testContainer.DB}

	insertedProducts, err := productdb.InsertOrUpdateProducts(products)
	assert.NoError(t, err)

	var orderPayload = OrderPayload{
		TotalPriceInCents: 5000,
		CreateOrderPayload: ordermodel.CreateOrderPayload{
			OrderItems: []ordermodel.OrderItem{
				{ProductID: insertedProducts[0].ID, PriceInCents: 1000, Quantity: 3},
				{ProductID: insertedProducts[1].ID, PriceInCents: 2000, Quantity: 1},
			},
		},
	}
	order, err := orderDb.InsertOrder(orderPayload)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.NotNil(t, order.OrderItems[0].ID)
	assert.NotNil(t, order.OrderItems[1].ID)
	assert.Equal(t, order.OrderItems, orderPayload.CreateOrderPayload.OrderItems)
	assert.Equal(t, order.CouponCode, orderPayload.CreateOrderPayload.CouponCode)
}

func ToPtr[T any](v T) *T {
	return &v
}
