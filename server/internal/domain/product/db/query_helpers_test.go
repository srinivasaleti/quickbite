package db

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestConstructGetProductsWhereQuery(t *testing.T) {
	t.Run("no filters", func(t *testing.T) {
		filters := GetProductFilters{}
		whereClause, args := constructGetProductsWhereQuery(filters)

		assert.Equal(t, "", whereClause)
		assert.Equal(t, pgx.NamedArgs{}, args)
	})

	t.Run("with single ID", func(t *testing.T) {
		filters := GetProductFilters{
			IDs: []string{"10", "20"},
		}
		whereClause, args := constructGetProductsWhereQuery(filters)

		assert.Equal(t, "WHERE p.id = ANY(@ids)", whereClause)
		assert.Equal(t, pgx.NamedArgs{"ids": []string{"10", "20"}}, args)
	})

}
