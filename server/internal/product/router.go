package product

import (
	"quickbite/server/internal/product/handler"

	"github.com/go-chi/chi/v5"
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
