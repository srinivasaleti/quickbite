package db

import "github.com/srinivasaleti/planner/server/internal/product/model"

type IProductDB interface {
	GetProducts() ([]model.Product, error)
}

type ProductDB struct{}

func (db *ProductDB) GetProducts() ([]model.Product, error) {
	return products, nil
}

func NewProductDB() IProductDB {
	return &ProductDB{}
}

var products = []model.Product{
	{
		ID:       "10",
		Name:     "Chicken Waffle",
		Price:    1,
		Category: "Waffle",
	},
}
