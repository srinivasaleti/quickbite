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

func (m *MockProductDB) InsertOrUpdateCategories(categories []model.Category) ([]model.Category, error) {
	args := m.Mock.Called(categories)
	result, _ := args.Get(0).([]model.Category)
	return result, args.Error(1)
}

func (m *MockProductDB) InsertOrUpdateProducts(products []model.Product) ([]model.Product, error) {
	args := m.Mock.Called(products)
	result, _ := args.Get(0).([]model.Product)
	return result, args.Error(1)
}

func (m *MockProductDB) GetProductById(id string) (*model.Product, error) {
	args := m.Mock.Called(id)
	result, _ := args.Get(0).(model.Product)
	return &result, args.Error(1)
}

func (m *MockProductDB) Reset() {
	m.ExpectedCalls = nil
}
