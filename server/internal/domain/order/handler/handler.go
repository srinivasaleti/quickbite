package handler

import (
	"encoding/json"
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

	var payload ordermodel.CreateOrderPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, "failed to decode order payload")
		httputils.WriteError(w, "invalid request body", httputils.BadRquest)
		return
	}

	if err := payload.Validate(); err != nil {
		h.Logger.Error(err, "order payload validation failed")
		httputils.WriteError(w, err.Error(), httputils.BadRquest)
		return
	}

	// Fetch products and update the order item prices
	productsIds := payload.GetProductIds()
	products, err := h.ProductDB.GetProducts(productdb.GetProductFilters{IDs: productsIds})
	if err != nil {
		h.Logger.Error(err, "unable to get products")
		httputils.WriteError(w, "unable to get products", httputils.InternalServerError)
		return
	}
	if len(products) != len(productsIds) {
		h.Logger.Error(err, "invalid product ids")
		httputils.WriteError(w, "invalid product ids", httputils.BadRquest)
		return
	}
	updateOrderItemPrices(payload, products)

	// Create order
	order, err := h.OrderDB.InsertOrder(createDbOrderPayload(payload))
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

	order.Products = products

	httputils.WriteJSONResponse(w, order, http.StatusCreated)
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

func createDbOrderPayload(payload ordermodel.CreateOrderPayload) orderdb.OrderPayload {
	orderPayload := orderdb.OrderPayload{}
	orderPayload.CreateOrderPayload = payload
	orderPayload.TotalPriceInCents = 0
	for _, item := range payload.OrderItems {
		orderPayload.TotalPriceInCents = orderPayload.TotalPriceInCents + item.PriceInCents.Multiply(item.Quantity)
	}
	return orderPayload
}

func NewOrderHanlder(logger logger.ILogger, db database.DB) OrderHandler {
	return OrderHandler{
		Logger:    logger,
		OrderDB:   orderdb.NewOrderDB(db),
		ProductDB: productdb.NewProductDB(db),
	}
}
