package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	couponservice "github.com/srinivasaleti/quickbite/server/internal/domain/coupon/service"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var couponServiceMock = couponservice.MockCouponService{}

func getHandler() CouponHandler {
	handler := CouponHandler{
		Logger:        &logger.Logger{},
		CouponSerivce: &couponServiceMock,
	}
	return handler
}

func reset() {
	couponServiceMock.Reset()
}

func TestCreateOrder(t *testing.T) {
	t.Run("should return 404 if code is invalid", func(t *testing.T) {
		reset()
		coupon := "123"
		couponServiceMock.On("IsValidCoupon", coupon).Return(couponservice.ErrInvalidCouponCode)
		rr := createValidateCouponRequest(coupon)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})

	t.Run("should return 500 if unable to verify", func(t *testing.T) {
		reset()
		coupon := "123"
		couponServiceMock.On("IsValidCoupon", coupon).Return(errors.New("something goes wrong"))
		rr := createValidateCouponRequest(coupon)
		assert.Equal(t, rr.Code, http.StatusInternalServerError)
	})

	t.Run("should return 200 on validating code", func(t *testing.T) {
		reset()
		coupon := "123"
		couponServiceMock.On("IsValidCoupon", coupon).Return(nil)
		rr := createValidateCouponRequest(coupon)
		assert.Equal(t, rr.Code, http.StatusOK)
	})

}

func createValidateCouponRequest(coupon string) *httptest.ResponseRecorder {
	handler := getHandler()
	router := chi.NewRouter()
	router.Post("/api/coupon/{couponCode}/validate", handler.ValidateCoupon)
	req, _ := http.NewRequest("POST", "/api/coupon/"+coupon+"/validate", bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
