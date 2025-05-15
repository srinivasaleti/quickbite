package db

import (
	"strings"

	"github.com/jackc/pgx/v5"
)

// constructGetProductsWhereQuery builds a WHERE clause and pgx.NamedArgs
// for filtering products.
func constructGetProductsWhereQuery(filters GetProductFilters) (string, pgx.NamedArgs) {
	args := pgx.NamedArgs{}
	whereFilters := []string{}

	if len(filters.IDs) > 0 {
		args["ids"] = filters.IDs
		whereFilters = append(whereFilters, "p.id = ANY(@ids)")
	}
	if len(whereFilters) > 0 {
		return "WHERE " + strings.Join(whereFilters, " AND "), args
	}
	return "", args
}
