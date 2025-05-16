package db

import (
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	"github.com/stretchr/testify/mock"
)

type MockOrderDB struct {
	mock.Mock
}

func (m *MockOrderDB) InsertOrder(createOrderPayload ordermodel.Order) (*ordermodel.Order, error) {
	args := m.Mock.Called(createOrderPayload)
	result, _ := args.Get(0).(ordermodel.Order)
	return &result, args.Error(1)
}

func (m *MockOrderDB) Reset() {
	m.ExpectedCalls = nil
}
