package service

import "errors"

var ErrInvalidCouponCode = errors.New("error Invalid Coupcode")

type ICouponService interface {
	IsValidCoupon(coupon string) error
}

type CouponService struct{}

func (c *CouponService) IsValidCoupon(coupon string) error {
	return ErrInvalidCouponCode
}

func NewCouponService() ICouponService {
	return &CouponService{}
}
