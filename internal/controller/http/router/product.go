package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Product(mr *mux.Router, ph ProductHandler) *mux.Router {
	r := mr.PathPrefix("/product").Subrouter()
	r.HandleFunc("/list", ph.GetProductsHandler).Methods(http.MethodPost)
	r.HandleFunc("", ph.GetProductById).Methods(http.MethodGet)
	r.HandleFunc("", ph.PatchProduct).Methods(http.MethodPatch)
	r.HandleFunc("", ph.DeleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("", ph.AddProduct).Methods(http.MethodPost)
	return r
}
