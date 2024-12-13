package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Category(mr *mux.Router, ch CategoryHandler) *mux.Router {
	r := mr.PathPrefix("/categories").Subrouter()
	r.HandleFunc("", ch.GetCategories).Methods(http.MethodGet)
	return r
}
