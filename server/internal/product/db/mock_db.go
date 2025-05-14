package db

import (
	"github.com/srinivasaleti/quickbite/server/internal/product/model"
	"github.com/stretchr/testify/mock"
)

type MockProductDB struct {
	mock.Mock
}

func (m *MockProductDB) GetProducts() ([]model.Product, error) {
	args := m.Mock.Called()
	result, _ := args.Get(0).([]model.Product)
	return result, args.Error(1)
}

func (m *MockProductDB) Reset() {
	m.ExpectedCalls = nil
}
