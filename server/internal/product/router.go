package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/planner/server/internal/product/handler"
	"github.com/srinivasaleti/planner/server/pkg/logger"
)

type ProductRouter struct {
	Handler handler.ProductHandler
}

func (config *ProductRouter) AddRoutesToAppRouter(appRouter chi.Router) {
	appRouter.Get("/products", config.Handler.GetProducts)
}

func NewProductRouter(logger logger.ILogger) ProductRouter {
	return ProductRouter{
		Handler: handler.NewProductHandler(logger),
	}
}
