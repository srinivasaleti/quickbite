package handler

import (
	"net/http"

	productdb "github.com/srinivasaleti/quickbite/server/internal/product/db"
	"github.com/srinivasaleti/quickbite/server/pkg/httputils"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
)

type ProductHandler struct {
	Logger    logger.ILogger
	ProductDB productdb.IProductDB
}

func (c *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("request recieved to fetch all products")
	products, err := c.ProductDB.GetProducts()
	if err != nil {
		httputils.WriteError(w, "unable to get products", httputils.InternalServerError, http.StatusInternalServerError)
		return
	}
	c.Logger.Info("successfully recieved all products")
	httputils.WriteJSONResponse(w, products, http.StatusOK)
}

func NewProductHandler(logger logger.ILogger) ProductHandler {
	return ProductHandler{
		Logger:    logger,
		ProductDB: productdb.NewProductDB(),
	}
}
