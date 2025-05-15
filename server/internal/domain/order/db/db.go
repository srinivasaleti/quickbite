package db

import (
	"github.com/srinivasaleti/quickbite/server/internal/database"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
)

type IOrderDB interface {
	InsertOrder(createOrderPayload ordermodel.CreateOrderPayload) (*ordermodel.Order, error)
}

type OrderDB struct {
	DB database.DB
}

func (db *OrderDB) InsertOrder(createOrderPayload ordermodel.CreateOrderPayload) (*ordermodel.Order, error) {
	return &ordermodel.Order{}, nil
}

func NewOrderDB(db database.DB) IOrderDB {
	return &OrderDB{
		DB: db,
	}
}
