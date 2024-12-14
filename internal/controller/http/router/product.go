package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Product(mr *mux.Router, ph ProductHandler) *mux.Router {
	r := mr.PathPrefix("/product").Subrouter()
	r.HandleFunc("", ph.GetProductsHandler).Methods(http.MethodPost)
	r.HandleFunc("/concrete", ph.GetProductById).Methods(http.MethodGet)
	r.HandleFunc("/patch", ph.PatchProduct).Methods(http.MethodPatch)
	return r
}
