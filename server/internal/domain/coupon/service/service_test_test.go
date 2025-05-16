package service

import (
	"errors"
	"testing"

	"github.com/srinivasaleti/quickbite/server/pkg/bloomfilters"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

var bloomfilterMock = &bloomfilters.BloomFilterMock{}

func TestInit(t *testing.T) {
	t.Run("should load", func(t *testing.T) {
		bloomfilterMock.Reset()
		bloomfilterMock.On("Load", couponFiles).Return(nil)

		service := CouponService{logger: &logger.Logger{}, filter: bloomfilterMock}

		assert.Equal(t, service.init(), nil)
		bloomfilterMock.AssertExpectations(t)
	})

	t.Run("should return error if unable to load", func(t *testing.T) {
		bloomfilterMock.Reset()
		bloomfilterMock.On("Load", couponFiles).Return(errors.New("unable to load"))

		service := CouponService{logger: &logger.Logger{}, filter: bloomfilterMock}

		assert.Equal(t, service.init(), errors.New("unable to load"))
		bloomfilterMock.AssertExpectations(t)
	})
}

func TestIsValidCoupon(t *testing.T) {
	t.Run("should return ErrCouponsNotLoaded if coupons not loaded", func(t *testing.T) {
		bloomfilterMock.Reset()
		bloomfilterMock.On("IsLoaded").Return(false)

		service := CouponService{logger: &logger.Logger{}, filter: bloomfilterMock}
		coupon := "123"

		assert.Equal(t, service.ValidateCoupon(coupon), ErrCouponsNotLoaded)
		bloomfilterMock.AssertExpectations(t)
	})

	t.Run("should return ErrInvalidCouponCode if coupons not valid", func(t *testing.T) {
		coupon := "123"
		bloomfilterMock.Reset()
		bloomfilterMock.On("IsLoaded").Return(true)
		bloomfilterMock.
			On("ElmentExistsInWhichFiles", coupon).
			Return([]string{couponFiles[0]})

		service := CouponService{logger: &logger.Logger{}, filter: bloomfilterMock}

		assert.Equal(t, service.ValidateCoupon(coupon), ErrInvalidCouponCode)
		bloomfilterMock.AssertExpectations(t)
	})

	t.Run("should return ErrInvalidCouponCode if coupon exists atleast in two files", func(t *testing.T) {
		coupon := "123"
		bloomfilterMock.Reset()
		bloomfilterMock.On("IsLoaded").Return(true)
		bloomfilterMock.
			On("ElmentExistsInWhichFiles", coupon).
			Return([]string{couponFiles[0], couponFiles[1]})

		service := CouponService{logger: &logger.Logger{}, filter: bloomfilterMock}

		assert.NoError(t, service.ValidateCoupon(coupon))
		bloomfilterMock.AssertExpectations(t)
	})

}
