package router

import "net/http"

type ProductHandler interface {
	GetProductsHandler(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	PatchProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
}

type CategoryHandler interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	PatchCategory(w http.ResponseWriter, r *http.Request)
	AddCategory(w http.ResponseWriter, r *http.Request)
}

type CountryHandler interface {
	GetCountries(w http.ResponseWriter, r *http.Request)
	DeleteCountry(w http.ResponseWriter, r *http.Request)
	PatchCountry(w http.ResponseWriter, r *http.Request)
	AddCountry(w http.ResponseWriter, r *http.Request)
}

type PictureHandler interface {
	AddPictures(w http.ResponseWriter, r *http.Request)
	DeletePicture(w http.ResponseWriter, r *http.Request)
	GetAllProductPictures(w http.ResponseWriter, r *http.Request)
}
