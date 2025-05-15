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
	productmodel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var orderDBMock = orderdb.MockOrderDB{}
var validPayload = ordermodel.CreateOrderPayload{
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{
		{ProductID: "prod-123", Quantity: 2},
	},
}

var returnedOrder = ordermodel.Order{
	ID:         "order-123",
	CouponCode: nil,
	OrderItems: []ordermodel.OrderItem{
		{
			ID:        "item-1",
			ProductID: "prod-123",
			Quantity:  2,
		},
	},
	Products: []productmodel.Product{
		{
			ID:    "prod-123",
			Name:  "Test Product",
			Price: 100,
		},
	},
}

func getHandler() OrderHandler {
	handler := OrderHandler{
		Logger: &logger.Logger{},
	}
	handler.OrderDB = &orderDBMock
	return handler
}

func reset() {
	orderDBMock.Reset()
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

	t.Run("should return internal server error on db failure", func(t *testing.T) {
		reset()

		orderDBMock.
			On("InsertOrder", validPayload).
			Return(ordermodel.Order{}, errors.New("DB error"))

		rr := createOrderRequest(validPayload)

		orderDBMock.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should create and return order", func(t *testing.T) {
		reset()
		orderDBMock.
			On("InsertOrder", validPayload).
			Return(returnedOrder, nil)
		rr := createOrderRequest(validPayload)

		orderDBMock.AssertExpectations(t)
		assert.Equal(t, rr.Code, http.StatusCreated)

		var actualOrder ordermodel.Order
		_ = json.Unmarshal(rr.Body.Bytes(), &actualOrder)
		assert.Equal(t, returnedOrder, actualOrder)
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
