package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
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
		c.Logger.Error(err, "unable to get products")
		httputils.WriteError(w, "unable to get products", httputils.InternalServerError, http.StatusInternalServerError)
		return
	}
	c.Logger.Info("successfully recieved all products")
	httputils.WriteJSONResponse(w, products, http.StatusOK)
}

func (c *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "productId")
	c.Logger.Info("request recieved to get product by id", "productId", productId)
	product, err := c.ProductDB.GetProductById(productId)
	if err == productdb.ErrNoProductFound {
		c.Logger.Error(err, "product not found", "productId", productId)
		httputils.WriteError(w, "product not found", httputils.NotFound, http.StatusNotFound)
		return
	}
	if err != nil {
		c.Logger.Error(err, "unable to get product by id", "productId", productId)
		httputils.WriteError(w, "unable to get product by id", httputils.NotFound, http.StatusInternalServerError)
		return
	}
	c.Logger.Info("successfully recieved product by id", "productId", productId)
	httputils.WriteJSONResponse(w, product, http.StatusOK)
}

func NewProductHandler(logger logger.ILogger, db database.DB) ProductHandler {
	return ProductHandler{
		Logger:    logger,
		ProductDB: productdb.NewProductDB(db),
	}
}
