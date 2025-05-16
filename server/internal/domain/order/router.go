package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/order/handler"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type OrderRouter struct {
	Handler handler.OrderHandler
}

func (config *OrderRouter) AddRoutesToAppRouter(appRouter chi.Router) {
	appRouter.Post("/order", config.Handler.CreateOrder)
	appRouter.Post("/order/summary", config.Handler.CalculateOrderSummary)
}

func NewOrderRouter(logger logger.ILogger, db database.DB) OrderRouter {
	return OrderRouter{
		Handler: handler.NewOrderHanlder(logger, db),
	}
}
