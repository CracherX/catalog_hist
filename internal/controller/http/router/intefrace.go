package router

import "net/http"

type ProductHandler interface {
	GetProductsHandler(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	PatchProduct(w http.ResponseWriter, r *http.Request)
}

type CategoryHandler interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type CountryHandler interface {
	GetCountries(w http.ResponseWriter, r *http.Request)
}
