package router

import "net/http"

type ProductHandler interface {
	GetProductsHandler(w http.ResponseWriter, r *http.Request)
}
