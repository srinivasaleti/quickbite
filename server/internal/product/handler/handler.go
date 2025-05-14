package handler

import "net/http"

type ProductHandler struct{}

func (c *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get products!"))
}

func NewProductHandler() ProductHandler {
	return ProductHandler{}
}
