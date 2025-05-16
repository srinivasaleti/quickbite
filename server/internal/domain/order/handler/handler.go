package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	orderdb "github.com/srinivasaleti/quickbite/server/internal/domain/order/db"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
	productmodel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/httputils"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type OrderHandler struct {
	Logger    logger.ILogger
	OrderDB   orderdb.IOrderDB
	ProductDB productdb.IProductDB
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("request received to create new order")

	orderPayload, err := h.preapreOrderPayload(r, w)
	if err != nil {
		return
	}
	// Create order
	order, err := h.OrderDB.InsertOrder(*orderPayload)
	if err == orderdb.ErrInvalidProductID {
		h.Logger.Error(err, "invalid product id")
		httputils.WriteError(w, "invalid product id", httputils.BadRquest)
		return
	}
	if err != nil {
		h.Logger.Error(err, "failed to create order")
		httputils.WriteError(w, "unable to create order", httputils.InternalServerError)
		return
	}
	h.Logger.Info("order created successfully", "orderId", order.ID)

	httputils.WriteJSONResponse(w, order, http.StatusCreated)
}

func (h *OrderHandler) preapreOrderPayload(r *http.Request, w http.ResponseWriter) (*ordermodel.Order, error) {
	var payload ordermodel.CreateOrderPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, "failed to decode order payload")
		httputils.WriteError(w, "invalid request body", httputils.BadRquest)
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		h.Logger.Error(err, "order payload validation failed")
		httputils.WriteError(w, err.Error(), httputils.BadRquest)
		return nil, err
	}

	productsIds := payload.GetProductIds()
	products, err := h.ProductDB.GetProducts(productdb.GetProductFilters{IDs: productsIds})
	if err != nil {
		h.Logger.Error(err, "unable to get products")
		httputils.WriteError(w, "unable to get products", httputils.InternalServerError)
		return nil, err
	}
	if len(products) != len(productsIds) {
		h.Logger.Error(err, "invalid product ids")
		httputils.WriteError(w, "invalid product ids", httputils.BadRquest)
		return nil, errors.New("invalid product ids")
	}
	updateOrderItemPrices(payload, products)

	totalPrice, err := getTotalPrice(payload)
	if err != nil {
		h.Logger.Error(err, "invalid product id")
		httputils.WriteError(w, "invalid product id", "INVALID_COUPON")
		return nil, err
	}
	orderPayload := ordermodel.Order{
		OrderItems:        payload.OrderItems,
		CouponCode:        payload.CouponCode,
		TotalPriceInCents: totalPrice,
		Products:          products,
	}
	return &orderPayload, nil
}

func updateOrderItemPrices(payload ordermodel.CreateOrderPayload, products []productmodel.Product) {
	for index, item := range payload.OrderItems {
		for _, product := range products {
			if product.ID == item.ProductID {
				payload.OrderItems[index].PriceInCents = product.Price.ToCents()
			}
		}
	}
}

func NewOrderHanlder(logger logger.ILogger, db database.DB) OrderHandler {
	return OrderHandler{
		Logger:    logger,
		OrderDB:   orderdb.NewOrderDB(db),
		ProductDB: productdb.NewProductDB(db),
	}
}
