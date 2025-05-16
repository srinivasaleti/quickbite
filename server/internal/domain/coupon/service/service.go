package service

import (
	"errors"

	"github.com/srinivasaleti/quickbite/server/pkg/bloomfilters"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

var ErrInvalidCouponCode = errors.New("error Invalid Coupcode")
var ErrCouponsNotLoaded = errors.New("error coupons not loaded")
var couponFiles = []string{"./coupons/couponbase1.gz", "./coupons/couponbase2.gz", "./coupons/couponbase3.gz"}

type ICouponService interface {
	ValidateCoupon(coupon string) error
}

type CouponService struct {
	filter bloomfilters.IBloomFilter
	logger logger.ILogger
}

func (c *CouponService) ValidateCoupon(coupon string) error {
	if !c.filter.IsLoaded() {
		return ErrCouponsNotLoaded
	}
	if len(c.filter.ElmentExistsInWhichFiles(coupon)) >= 2 {
		return nil
	}
	return ErrInvalidCouponCode
}

func (c *CouponService) init() error {
	c.logger.Info("⏳ Loading coupons... This may take a few seconds")
	err := c.filter.Load(couponFiles)
	if err != nil {
		c.logger.Error(err, "❌ unable to load files")
		return err
	}
	c.logger.Info("✅ successfully loaded coupons")
	return nil
}

func NewCouponService(logger logger.ILogger) ICouponService {
	couponService := &CouponService{logger: logger, filter: bloomfilters.NewGzipBloomFilter()}
	go couponService.init()
	return couponService
}
