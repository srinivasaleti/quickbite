package db

import (
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/price"
	"github.com/stretchr/testify/assert"
)

var productsData = []model.Product{
	{
		Name:       "Test Product 1",
		ExternalID: ToPtr("test-prod-001"),
		Price:      10.5,
		ImageURL:   "http://example.com/image1.jpg",
	},
	{
		Name:       "Test Product 2",
		ExternalID: ToPtr("test-prod-002"),
		Price:      15.0,
		ImageURL:   "http://example.com/image2.jpg",
	},
}
var categoriesData = []model.Category{
	{Name: "Cake"},
	{Name: "Pie"},
	{Name: "Brownie"},
}

func TestProductDBOperations(t *testing.T) {
	testContainer, err := database.SetupTestDatabase()
	assert.NoError(t, err)
	productDB := ProductDB{DB: testContainer.DB}

	// should insert categories on first go
	insertCategories, err := productDB.InsertOrUpdateCategories(categoriesData)
	assert.NoError(t, err)
	assert.Equal(t, len(insertCategories), len(categoriesData))
	for _, c := range insertCategories {
		assert.NotEmpty(t, c.ID)
	}

	// should insert products on first go
	insertedProducts, err := productDB.InsertOrUpdateProducts(productsData)
	assert.Equal(t, len(insertedProducts), len(productsData))
	assert.NoError(t, err)
	for _, p := range insertedProducts {
		assert.NotEmpty(t, p.ID)
	}

	t.Run("insertOrUpdateCategories", func(t *testing.T) {
		updatedCategories, err := productDB.InsertOrUpdateCategories(categoriesData)
		assert.NoError(t, err)
		assert.Equal(t, len(updatedCategories), len(categoriesData))
		for index, c := range updatedCategories {
			assert.Equal(t, c.ID, insertCategories[index].ID)
		}

	})

	t.Run("insertOrUpdateProducts", func(t *testing.T) {
		updatedProducts, err := productDB.InsertOrUpdateProducts(productsData)
		productsData[0].Price = 100
		assert.Equal(t, len(updatedProducts), len(productsData))
		assert.NoError(t, err)
		assert.Equal(t, updatedProducts[0].Price, price.Price(100))
		for index, p := range updatedProducts {
			assert.Equal(t, p.ID, insertedProducts[index].ID)
		}
	})

	t.Run("get products", func(t *testing.T) {
		// Add first product to a category
		productsData[0].CategoryID = &insertCategories[0].ID
		_, err := productDB.InsertOrUpdateProducts(productsData)
		assert.NoError(t, err)

		products, err := productDB.GetProducts(GetProductFilters{})
		assert.NoError(t, err)
		assert.Equal(t, len(products), len(productsData))
		assert.Equal(t, products[0].Name, insertedProducts[0].Name)
		assert.Equal(t, products[0].CategoryName, &insertCategories[0].Name)

		// Filter by ids
		products, err = productDB.GetProducts(GetProductFilters{IDs: []string{insertedProducts[0].ID}})
		assert.NoError(t, err)
		assert.Equal(t, len(products), 1)
		assert.Equal(t, products[0].ID, insertedProducts[0].ID)
	})

	t.Run("get product by id", func(t *testing.T) {
		// Add first product to a category
		productsData[0].CategoryID = &insertCategories[0].ID
		_, err := productDB.InsertOrUpdateProducts(productsData)
		assert.NoError(t, err)

		// get product by id
		product, err := productDB.GetProductById(insertedProducts[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, product.Name, insertedProducts[0].Name)
		assert.Equal(t, product.CategoryName, &insertCategories[0].Name)

		// product not found
		product, err = productDB.GetProductById("2a7a8dc2-fdcd-4849-b64e-dde873b0de43")
		assert.Equal(t, err, ErrNoProductFound)
		assert.Nil(t, product)

		// unknown error
		product, err = productDB.GetProductById("invalid-id")
		assert.NotEqual(t, err, ErrNoProductFound)
		assert.Nil(t, product)
		assert.Error(t, err)
	})

	testContainer.TearDown()
}

func ToPtr[T any](v T) *T {
	return &v
}
