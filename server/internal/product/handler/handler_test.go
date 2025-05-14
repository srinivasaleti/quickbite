package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/srinivasaleti/planner/server/internal/product/db"
	"github.com/srinivasaleti/planner/server/internal/product/model"
	"github.com/srinivasaleti/planner/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var products = []model.Product{
	{
		ID:       "10",
		Name:     "Chicken Waffle",
		Price:    1,
		Category: "Waffle",
	},
}

var productDBMock = db.MockProductDB{}

func getHandler() ProductHandler {
	handler := NewProductHandler(&logger.Logger{})
	handler.ProductDB = &productDBMock
	return handler
}

func reset() {
	productDBMock.Reset()
}

func TestGetProduct(t *testing.T) {
	t.Run("should get products", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProducts").
			Return(products, nil)

		rr := getProducts()

		assert.Equal(t, http.StatusOK, rr.Code)
		var actualProducts []model.Product
		_ = json.Unmarshal(rr.Body.Bytes(), &actualProducts)
		assert.Equal(t, products, actualProducts)
	})

	t.Run("should return erorr if unable to get products", func(t *testing.T) {
		reset()

		productDBMock.
			On("GetProducts").
			Return(nil, errors.New("DB Error"))
		rr := getProducts()

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

}

func getProducts() *httptest.ResponseRecorder {
	handler := getHandler()
	req, _ := http.NewRequest("GET", "/api/products", nil)
	rr := httptest.NewRecorder()
	handler.GetProducts(rr, req)
	return rr
}
