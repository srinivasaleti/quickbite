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
	InsertOrder(order ordermodel.Order) (*ordermodel.Order, error)
}

type OrderDB struct {
	DB database.DB
}

func (db *OrderDB) InsertOrder(order ordermodel.Order) (*ordermodel.Order, error) {
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
		INSERT INTO orders (coupon_code, total_price)
		VALUES (@couponCode, @totalPrice)
		RETURNING id
	`, pgx.NamedArgs{"couponCode": order.CouponCode, "totalPrice": order.TotalPriceInCents}).Scan(&orderID)
	if err != nil {
		return nil, err
	}

	// create order items
	for _, item := range order.OrderItems {
		_, err := tx.Exec(ctx, `
			INSERT INTO order_items (order_id, product_id, quantity, price)
			VALUES (@orderId, @productId, @quantity, @price)
		`, pgx.NamedArgs{
			"productId": item.ProductID,
			"quantity":  item.Quantity,
			"price":     item.PriceInCents,
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

	order.ID = orderID
	return &order, nil
}

func NewOrderDB(db database.DB) IOrderDB {
	return &OrderDB{
		DB: db,
	}
}
