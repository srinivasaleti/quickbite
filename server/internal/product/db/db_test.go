package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	t.Run("should get products", func(t *testing.T) {
		db := ProductDB{}

		products, err := db.GetProducts()

		assert.NoError(t, err)
		assert.Equal(t, products, products)
	})
}
