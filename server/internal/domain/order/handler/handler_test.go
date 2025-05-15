package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	orderdb "github.com/srinivasaleti/quickbite/server/internal/domain/order/db"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
	productmodel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/srinivasaleti/quickbite/server/pkg/price"
	"github.com/stretchr/testify/assert"
)

var orderDBMock = orderdb.MockOrderDB{}
var productDBMock = productdb.MockProductDB{}

var validOrderPayload = ordermodel.CreateOrderPayload{
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{{ProductID: "prod-1", Quantity: 2}, {ProductID: "prod-2", Quantity: 5}},
}

var products = []productmodel.Product{
	{ID: "prod-1", Price: price.Price(10)},
	{ID: "prod-2", Price: price.Price(20)},
}

var validOrderPayloadWithPrices = ordermodel.CreateOrderPayload{
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{
		{ProductID: "prod-1", Quantity: 2, PriceInCents: products[0].Price.ToCents()},
		{ProductID: "prod-2", Quantity: 5, PriceInCents: products[1].Price.ToCents()}},
}

var orderPayload = orderdb.OrderPayload{
	CreateOrderPayload: validOrderPayloadWithPrices,
	TotalPriceInCents:  12000,
}

var returnedOrder = ordermodel.Order{
	ID:         "order-123",
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{
		{ID: "item-1", ProductID: "prod-1", Quantity: 2, PriceInCents: products[0].Price.ToCents()},
		{ID: "item-1", ProductID: "prod-2", Quantity: 2, PriceInCents: products[1].Price.ToCents()},
	},
}

func getHandler() OrderHandler {
	handler := OrderHandler{
		Logger: &logger.Logger{},
	}
	handler.OrderDB = &orderDBMock
	handler.ProductDB = &productDBMock
	return handler
}

func reset() {
	orderDBMock.Reset()
	productDBMock.Reset()
}

func TestCreateOrder(t *testing.T) {
	t.Run("should return bad request if unable decode body", func(t *testing.T) {
		reset()
		rr := createOrderRequest("invalid body")
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return bad request if payload is invalid", func(t *testing.T) {
		reset()
		invalidPayload := map[string]interface{}{
			"items": []interface{}{},
		}
		rr := createOrderRequest(invalidPayload)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return internal server error if unable to get products", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return(nil, errors.New("unabel to get products"))

		rr := createOrderRequest(validOrderPayload)

		productDBMock.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should return bad request if all produts not found", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return([]productmodel.Product{products[0]}, nil)
		rr := createOrderRequest(validOrderPayload)

		productDBMock.AssertExpectations(t)
		orderDBMock.AssertExpectations(t)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return internal server error on db failure", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return(products, nil)
		orderDBMock.
			On("InsertOrder", orderPayload).
			Return(ordermodel.Order{}, errors.New("DB error"))

		rr := createOrderRequest(validOrderPayload)

		productDBMock.AssertExpectations(t)
		orderDBMock.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should create and return order", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return(products, nil)
		orderDBMock.
			On("InsertOrder", orderPayload).
			Return(returnedOrder, nil)

		rr := createOrderRequest(validOrderPayload)

		orderDBMock.AssertExpectations(t)
		productDBMock.AssertExpectations(t)
		assert.Equal(t, rr.Code, http.StatusCreated)

		var actualOrder ordermodel.Order
		_ = json.Unmarshal(rr.Body.Bytes(), &actualOrder)
		returnedOrder.Products = products
		assert.Equal(t, returnedOrder, actualOrder)
		// Prices should be converted to main price
		assert.Equal(t, returnedOrder.Products[0].Price, price.Price(10))
		assert.Equal(t, returnedOrder.OrderItems[0].PriceInCents, price.Cent(1000))
	})
}

func TestUpdateOrderItemPrices(t *testing.T) {
	t.Run("should update price for matching product IDs", func(t *testing.T) {
		payload := ordermodel.CreateOrderPayload{OrderItems: []ordermodel.OrderItem{{ProductID: "1"}, {ProductID: "2"}}}
		products := []productmodel.Product{{ID: "1", Price: price.Price(100)}, {ID: "2", Price: 200}}

		updateOrderItemPrices(payload, products)

		assert.Equal(t, price.Cent(10000), payload.OrderItems[0].PriceInCents)
		assert.Equal(t, price.Cent(20000), payload.OrderItems[1].PriceInCents)
	})

	t.Run("should not change price if product ID does not match", func(t *testing.T) {
		payload := ordermodel.CreateOrderPayload{OrderItems: []ordermodel.OrderItem{{ProductID: "1"}, {ProductID: "3"}}}

		products := []productmodel.Product{{ID: "1", Price: price.Price(100)}, {ID: "2", Price: price.Price(200)}}

		updateOrderItemPrices(payload, products)

		assert.Equal(t, price.Cent(10000), payload.OrderItems[0].PriceInCents)
		assert.Equal(t, price.Cent(0), payload.OrderItems[1].PriceInCents)
	})
}

func createOrderRequest(payload interface{}) *httptest.ResponseRecorder {
	handler := getHandler()
	router := chi.NewRouter()
	router.Post("/api/order", handler.CreateOrder)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/order", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
