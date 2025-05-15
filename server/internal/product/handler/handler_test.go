package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/quickbite/server/internal/product/db"
	productdb "github.com/srinivasaleti/quickbite/server/internal/product/db"
	"github.com/srinivasaleti/quickbite/server/internal/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var products = []model.Product{
	{
		ID:    "10",
		Name:  "Chicken Waffle",
		Price: 1,
	},
}

var productDBMock = db.MockProductDB{}

func getHandler() ProductHandler {
	handler := ProductHandler{
		Logger: &logger.Logger{},
	}
	handler.ProductDB = &productDBMock
	return handler
}

func reset() {
	productDBMock.Reset()
}

func TestGetProducts(t *testing.T) {
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

func TestGetProductByID(t *testing.T) {
	t.Run("should get product by id", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProductById", products[0].ID).
			Return(products[0], nil)

		rr := getProductById(products[0].ID)

		assert.Equal(t, http.StatusOK, rr.Code)
		productDBMock.AssertExpectations(t)
		var actualProduct model.Product
		_ = json.Unmarshal(rr.Body.Bytes(), &actualProduct)
		assert.Equal(t, products[0], actualProduct)
	})

	t.Run("should return not found if the product not exists", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProductById", products[0].ID).
			Return(nil, productdb.ErrNoProductFound)

		rr := getProductById(products[0].ID)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("should return internal server error if something goes wrong", func(t *testing.T) {
		reset()
		productDBMock.
			On("GetProductById", products[0].ID).
			Return(nil, errors.New("Db Error"))

		rr := getProductById(products[0].ID)

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

func getProductById(id string) *httptest.ResponseRecorder {
	handler := getHandler()
	router := chi.NewRouter()
	router.Get("/api/product/{productId}", handler.GetProduct)
	req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
