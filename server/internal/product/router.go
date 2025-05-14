package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/planner/server/internal/product/handler"
)

type ProductRouter struct {
	Handler handler.ProductHandler
}

func (config *ProductRouter) AddRoutesToAppRouter(appRouter chi.Router) {
	appRouter.Get("/products", config.Handler.GetProducts)
}

func NewProductRouter() ProductRouter {
	return ProductRouter{
		Handler: handler.NewProductHandler(),
	}
}
