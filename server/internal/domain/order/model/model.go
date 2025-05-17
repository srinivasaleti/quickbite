package model

import (
	"errors"
	"fmt"
	"time"

	productModel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/price"
)

type OrderItem struct {
	ID           string       `json:"id,omitempty"`
	ProductID    string       `json:"productId"`
	PriceInCents price.Cent   `json:"priceInCents,omitempty"`
	Price        price.Dollar `json:"price,omitempty"`
	Quantity     int          `json:"quantity"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time   `json:"updatedAt,omitempty"`
}

type Order struct {
	ID                string                 `json:"id"`
	TotalPriceInCents price.Cent             `json:"totalPriceInCents,omitempty"`
	CouponCode        *string                `json:"couponCode,omitempty"`
	CreatedAt         *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time             `json:"updatedAt,omitempty"`
	OrderItems        []OrderItem            `json:"items,omitempty"`
	Products          []productModel.Product `json:"products,omitempty"`
}

type CreateOrderPayload struct {
	CouponCode *string     `json:"couponCode"`
	OrderItems []OrderItem `json:"items,omitempty"`
}

func (p *CreateOrderPayload) Validate() error {
	if len(p.OrderItems) == 0 {
		return errors.New("order must have at least one item")
	}

	for i, item := range p.OrderItems {
		if item.ProductID == "" {
			return fmt.Errorf("productId is required in item at index %d", i)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("quantity must be greater than 0 in item at index %d", i)
		}
	}

	return nil
}

func (payload *CreateOrderPayload) GetProductIds() []string {
	ids := []string{}
	for _, item := range payload.OrderItems {
		ids = append(ids, item.ProductID)
	}
	return ids
}
