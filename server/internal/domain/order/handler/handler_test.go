package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	couponservice "github.com/srinivasaleti/quickbite/server/internal/domain/coupon/service"
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
var couponServiceMock = couponservice.MockCouponService{}
var validOrderPayload = ordermodel.CreateOrderPayload{
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{{ProductID: "prod-1", Quantity: 2}, {ProductID: "prod-2", Quantity: 5}},
}

var products = []productmodel.Product{
	{ID: "prod-1", Price: price.Price(10)},
	{ID: "prod-2", Price: price.Price(20)},
}

var order = ordermodel.Order{
	TotalPriceInCents: 12000,
	OrderItems: []ordermodel.OrderItem{
		{ProductID: "prod-1", Quantity: 2, PriceInCents: products[0].Price.ToCents()},
		{ProductID: "prod-2", Quantity: 5, PriceInCents: products[1].Price.ToCents()},
	},
	Products: products,
}

func getHandler() OrderHandler {
	handler := OrderHandler{
		Logger:        &logger.Logger{},
		OrderDB:       &orderDBMock,
		ProductDB:     &productDBMock,
		CouponService: &couponServiceMock,
	}
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
			On("InsertOrder", order).
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
			On("InsertOrder", order).
			Return(order, nil)

		rr := createOrderRequest(validOrderPayload)

		orderDBMock.AssertExpectations(t)
		productDBMock.AssertExpectations(t)
		assert.Equal(t, rr.Code, http.StatusCreated)

		var actualOrder ordermodel.Order
		_ = json.Unmarshal(rr.Body.Bytes(), &actualOrder)
		assert.Equal(t, order, actualOrder)
		// Prices should be converted to main price
		assert.Equal(t, order.Products[0].Price, price.Price(10))
		assert.Equal(t, order.OrderItems[0].PriceInCents, price.Cent(1000))
	})

	t.Run("should validate coupon if exits", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return(products, nil)
		couponServiceMock.
			On("ValidateCoupon", "Coupon1").
			Return(nil)

		// Apply coupon code
		validOrderPayload.CouponCode = ToPtr("Coupon1")
		order.TotalPriceInCents = order.TotalPriceInCents - order.TotalPriceInCents.Percentize(10)
		order.CouponCode = ToPtr("Coupon1")
		orderDBMock.
			On("InsertOrder", order).
			Return(order, nil)

		rr := createOrderRequest(validOrderPayload)

		orderDBMock.AssertExpectations(t)
		productDBMock.AssertExpectations(t)
		assert.Equal(t, rr.Code, http.StatusCreated)

		var actualOrder ordermodel.Order
		_ = json.Unmarshal(rr.Body.Bytes(), &actualOrder)
		assert.Equal(t, order, actualOrder)
		// Prices should be converted to main price
		assert.Equal(t, order.Products[0].Price, price.Price(10))
		assert.Equal(t, order.OrderItems[0].PriceInCents, price.Cent(1000))
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

func TestCalculateOrderSummary(t *testing.T) {
	t.Run("should return bad request if unable decode body", func(t *testing.T) {
		reset()
		rr := orderSummaryRequest("invalid body")
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should create and return order", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts", productdb.GetProductFilters{IDs: []string{
				validOrderPayload.OrderItems[0].ProductID,
				validOrderPayload.OrderItems[1].ProductID,
			}}).Return(products, nil)

		rr := orderSummaryRequest(validOrderPayload)

		orderDBMock.AssertExpectations(t)
		productDBMock.AssertExpectations(t)
		assert.Equal(t, rr.Code, http.StatusCreated)

		var actualOrder ordermodel.Order
		_ = json.Unmarshal(rr.Body.Bytes(), &actualOrder)
		assert.Equal(t, order, actualOrder)
		// Prices should be converted to main price
		assert.Equal(t, order.Products[0].Price, price.Price(10))
		assert.Equal(t, order.OrderItems[0].PriceInCents, price.Cent(1000))
		orderDBMock.AssertNotCalled(t, "InsertOrder")
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

func orderSummaryRequest(payload interface{}) *httptest.ResponseRecorder {
	handler := getHandler()
	router := chi.NewRouter()
	router.Post("/api/order/summary", handler.CalculateOrderSummary)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/order/summary", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
