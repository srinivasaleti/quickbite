package handler

import (
	"encoding/json"
	"net/http"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	orderdb "github.com/srinivasaleti/quickbite/server/internal/domain/order/db"
	"github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	"github.com/srinivasaleti/quickbite/server/pkg/httputils"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type OrderHandler struct {
	Logger  logger.ILogger
	OrderDB orderdb.IOrderDB
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("request received to create new order")

	var payload model.CreateOrderPayload

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

	order, err := h.OrderDB.InsertOrder(payload)
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

func NewOrderHanlder(logger logger.ILogger, db database.DB) OrderHandler {
	return OrderHandler{
		Logger:  logger,
		OrderDB: orderdb.NewOrderDB(db),
	}
}
