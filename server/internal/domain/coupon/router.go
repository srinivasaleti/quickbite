package coupon

import (
	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/quickbite/server/internal/domain/coupon/handler"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type CouponRouter struct {
	Handler handler.CouponHandler
}

func (config *CouponRouter) AddRoutesToAppRouter(appRouter chi.Router) {
	appRouter.Post("/coupon/{couponCode}/validate", config.Handler.ValidateCoupon)
}

func NewCouponRouter(logger logger.ILogger) CouponRouter {
	return CouponRouter{
		Handler: handler.NewCouponHandler(logger),
	}
}
