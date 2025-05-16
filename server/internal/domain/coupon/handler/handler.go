package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	couponservice "github.com/srinivasaleti/quickbite/server/internal/domain/coupon/service"
	"github.com/srinivasaleti/quickbite/server/pkg/httputils"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type CouponHandler struct {
	Logger        logger.ILogger
	CouponSerivce couponservice.ICouponService
}

func (h *CouponHandler) ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	couponCode := chi.URLParam(r, "couponCode")
	h.Logger.Info("request received to validate coupon code")
	err := h.CouponSerivce.IsValidCoupon(couponCode)
	if err == couponservice.ErrInvalidCouponCode {
		h.Logger.Error(err, "invalid coupon")
		httputils.WriteError(w, "invalid coupon", httputils.BadRquest)
		return
	}
	if err != nil {
		h.Logger.Error(err, "unable to validate coupon")
		httputils.WriteError(w, "unable to validate coupon", httputils.InternalServerError)
		return
	}
	h.Logger.Info("successfully validated couponCode")
	httputils.WriteJSONResponse(w, nil, http.StatusOK)
}

func NewCouponHandler(logger logger.ILogger) CouponHandler {
	return CouponHandler{
		Logger:        logger,
		CouponSerivce: couponservice.NewCouponService(),
	}
}
