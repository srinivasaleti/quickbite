package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
)

var ErrInvalidProductID = errors.New("invalid product id")

type IOrderDB interface {
	InsertOrder(createOrderPayload ordermodel.CreateOrderPayload) (*ordermodel.Order, error)
}

type OrderDB struct {
	DB database.DB
}

func (db *OrderDB) InsertOrder(payload ordermodel.CreateOrderPayload) (*ordermodel.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// create order
	var orderID string
	err = tx.QueryRow(ctx, `
		INSERT INTO orders (coupon_code)
		VALUES (@couponCode)
		RETURNING id
	`, pgx.NamedArgs{"couponCode": payload.CouponCode}).Scan(&orderID)
	if err != nil {
		return nil, err
	}

	// create order items
	for _, item := range payload.OrderItems {
		_, err := tx.Exec(ctx, `
			INSERT INTO order_items (order_id, product_id, quantity, price)
			VALUES (@orderId, @productId, @quantity, @price)
		`, pgx.NamedArgs{
			"productId": item.ProductID,
			"quantity":  item.Quantity,
			"price":     item.Price,
			"orderId":   orderID,
		})
		if database.ErrIsConstraint(err, "order_items_product_id_fkey") {
			return nil, ErrInvalidProductID
		}
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &ordermodel.Order{
		ID:         orderID,
		CouponCode: payload.CouponCode,
		OrderItems: payload.OrderItems,
	}, nil
}

func NewOrderDB(db database.DB) IOrderDB {
	return &OrderDB{
		DB: db,
	}
}
