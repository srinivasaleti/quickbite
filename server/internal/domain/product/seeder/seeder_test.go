package seeder

import (
	"testing"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestSeeProducts(t *testing.T) {
	var products []model.Product
	err := readData(ProductsFile, &products)
	assert.NoError(t, err)

	var categories []model.Category
	err = readData(CategoriesFile, &categories)
	assert.NoError(t, err)

	testContainer, err := database.SetupTestDatabase()
	assert.NoError(t, err)

	t.Run("should seed both categories & products", func(t *testing.T) {
		seeder := NewProductSeeder(&logger.Logger{}, testContainer.DB)

		result, err := seeder.SeedProducts()

		assert.NoError(t, err)
		assert.Equal(t, len(result.Products), len(products))
		assert.Equal(t, len(result.Categories), len(categories))
		assert.Equal(t, result.Products[0].Name, products[0].Name)
		assert.Equal(t, result.Categories[0].Name, categories[0].Name)
	})
}
