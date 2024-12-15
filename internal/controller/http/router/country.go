package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Country(mr *mux.Router, ch CountryHandler) *mux.Router {
	r := mr.PathPrefix("/country").Subrouter()
	r.HandleFunc("", ch.GetCountries).Methods(http.MethodGet)
	r.HandleFunc("", ch.PatchCountry).Methods(http.MethodPatch)
	r.HandleFunc("", ch.DeleteCountry).Methods(http.MethodDelete)
	r.HandleFunc("", ch.AddCountry).Methods(http.MethodPost)
	return r
}
