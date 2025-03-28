package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Pictures(mr *mux.Router, ch PictureHandler) *mux.Router {
	r := mr.PathPrefix("/pictures").Subrouter()
	r.HandleFunc("", ch.GetAllProductPictures).Methods(http.MethodGet)
	r.HandleFunc("", ch.DeletePicture).Methods(http.MethodDelete)
	r.HandleFunc("", ch.AddPictures).Methods(http.MethodPost)
	return r
}
