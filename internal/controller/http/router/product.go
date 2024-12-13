package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Product(mr *mux.Router, ph ProductHandler) *mux.Router {
	r := mr.PathPrefix("/product").Subrouter()
	r.HandleFunc("", ph.GetProductsHandler).Methods(http.MethodGet)
	r.HandleFunc("/concrete", ph.GetProductById).Methods(http.MethodGet)
	return r
}
