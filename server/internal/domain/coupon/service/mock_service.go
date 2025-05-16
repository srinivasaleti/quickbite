package service

import "github.com/stretchr/testify/mock"

type MockCouponService struct {
	mock.Mock
}

func (m *MockCouponService) ValidateCoupon(coupon string) error {
	args := m.Called(coupon)
	return args.Error(0)
}

func (m *MockCouponService) Reset() {
	m.ExpectedCalls = nil
}
