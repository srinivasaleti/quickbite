package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product/handler"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type ProductRouter struct {
	Handler handler.ProductHandler
}

func (config *ProductRouter) AddRoutesToAppRouter(appRouter chi.Router) {
	appRouter.Get("/product", config.Handler.GetProducts)
	appRouter.Get("/product/{productId}", config.Handler.GetProduct)
}

func NewProductRouter(logger logger.ILogger, db database.DB) ProductRouter {
	return ProductRouter{
		Handler: handler.NewProductHandler(logger, db),
	}
}
