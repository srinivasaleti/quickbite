package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	couponService "github.com/srinivasaleti/quickbite/server/internal/domain/coupon/service"
	orderdb "github.com/srinivasaleti/quickbite/server/internal/domain/order/db"
	ordermodel "github.com/srinivasaleti/quickbite/server/internal/domain/order/model"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
	productmodel "github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/httputils"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type OrderHandler struct {
	Logger        logger.ILogger
	OrderDB       orderdb.IOrderDB
	ProductDB     productdb.IProductDB
	CouponService couponService.ICouponService
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("request received to create new order")

	orderPayload, err := h.prepareOrderSummary(r, w)
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

func (h *OrderHandler) CalculateOrderSummary(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("request received to create new order")
	orderPayload, err := h.prepareOrderSummary(r, w)
	if err != nil {
		return
	}
	httputils.WriteJSONResponse(w, orderPayload, http.StatusCreated)
}

func (h *OrderHandler) prepareOrderSummary(r *http.Request, w http.ResponseWriter) (*ordermodel.Order, error) {
	var payload ordermodel.CreateOrderPayload
	// validate request
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

	// update order prices
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

	// validate coupons
	err = h.validateCouponCode(payload.CouponCode)
	if err == couponService.ErrCouponsNotLoaded {
		h.Logger.Error(err, "coupons not loaded yet")
		httputils.WriteError(w, "Coupons are not loaded yet. Plese try after some time or remove coupon", httputils.InvalidOrder)
		return nil, err
	}
	if err != nil {
		h.Logger.Error(err, "invalid coupon code")
		httputils.WriteError(w, "Invalid coupon code. Either submit a valid coupon code or remove coupon code", httputils.InvalidOrder)
		return nil, errors.New("invalid coupon code")
	}

	// prepare order
	orderPayload := ordermodel.Order{
		OrderItems: payload.OrderItems,
		CouponCode: payload.CouponCode,
		Products:   products,
	}
	totalPrice, err := getTotalPrice(orderPayload, true)
	if err != nil {
		h.Logger.Error(err, "invalid order")
		httputils.WriteError(w, err.Error(), httputils.InvalidOrder)
		return nil, err
	}
	orderPayload.TotalPriceInCents = totalPrice

	return &orderPayload, nil
}

func (h *OrderHandler) validateCouponCode(coupon *string) error {
	if coupon == nil {
		return nil
	}
	// HAPPYHOURS, BUYGETONE are not valid coupon codes. They are not present in any one of the coupon file.
	// However https://github.com/oolio-group/kart-challenge says HAPPYHOURS, BUYGETONE as a valid coupons.
	// So explicity adding these two as valid coupons.
	if *coupon == "HAPPYHOURS" || *coupon == "BUYGETONE" {
		return nil
	}
	return h.CouponService.ValidateCoupon(*coupon)
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
		Logger:        logger,
		OrderDB:       orderdb.NewOrderDB(db),
		ProductDB:     productdb.NewProductDB(db),
		CouponService: couponService.NewCouponService(logger),
	}
}
