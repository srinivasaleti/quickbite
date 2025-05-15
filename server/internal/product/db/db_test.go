package db

import (
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/product/model"
	"github.com/stretchr/testify/assert"
)

var productsData = []model.Product{
	{
		Name:       "Test Product 1",
		ExternalID: "test-prod-001",
		Price:      10.5,
		ImageURL:   "http://example.com/image1.jpg",
	},
	{
		Name:       "Test Product 2",
		ExternalID: "test-prod-002",
		Price:      15.0,
		ImageURL:   "http://example.com/image2.jpg",
	},
}
var categoriesData = []model.Category{
	{Name: "Cake"},
	{Name: "Pie"},
	{Name: "Brownie"},
}

func TestGetProducts(t *testing.T) {
	t.Run("should get products", func(t *testing.T) {
		db := ProductDB{}

		products, err := db.GetProducts()

		assert.NoError(t, err)
		assert.Equal(t, products, products)
	})
}

func TestProductDBOperations(t *testing.T) {
	testContainer, err := database.SetupTestDatabase()
	assert.NoError(t, err)
	productDB := ProductDB{DB: testContainer.DB}

	t.Run("insertOrUpdateCategories", func(t *testing.T) {
		// should insert on first go
		insertCategoreis, err := productDB.InsertOrUpdateCategories(categoriesData)
		assert.NoError(t, err)
		assert.Equal(t, len(insertCategoreis), len(categoriesData))
		for _, c := range insertCategoreis {
			assert.NotEmpty(t, c.ID)
		}

		// should update on first go
		updatedCategories, err := productDB.InsertOrUpdateCategories(categoriesData)
		assert.NoError(t, err)
		assert.Equal(t, len(updatedCategories), len(categoriesData))
		for index, c := range updatedCategories {
			assert.Equal(t, c.ID, insertCategoreis[index].ID)
		}

	})

	t.Run("insertOrUpdateProducts", func(t *testing.T) {
		// should insert on first go
		insertedProducts, err := productDB.InsertOrUpdateProducts(productsData)
		assert.Equal(t, len(insertedProducts), len(productsData))
		assert.NoError(t, err)
		for _, p := range insertedProducts {
			assert.NotEmpty(t, p.ID)
		}

		// should update on first go
		updatedProducts, err := productDB.InsertOrUpdateProducts(productsData)
		productsData[0].Price = 100
		assert.Equal(t, len(updatedProducts), len(productsData))
		assert.NoError(t, err)
		assert.Equal(t, updatedProducts[0].Price, float64(100))
		for index, p := range updatedProducts {
			assert.Equal(t, p.ID, insertedProducts[index].ID)
		}
	})
}
