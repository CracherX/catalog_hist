package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Category(mr *mux.Router, ch CategoryHandler) *mux.Router {
	r := mr.PathPrefix("/categories").Subrouter()
	r.HandleFunc("", ch.GetCategories).Methods(http.MethodGet)
	r.HandleFunc("", ch.AddCategory).Methods(http.MethodPost)
	r.HandleFunc("", ch.DeleteCategory).Methods(http.MethodDelete)
	r.HandleFunc("", ch.PatchCategory).Methods(http.MethodPatch)
	return r
}
