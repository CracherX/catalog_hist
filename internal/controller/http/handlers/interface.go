package handlers

import (
	"github.com/CracherX/catalog_hist/internal/entity"
	"net/http"
)

type ProductUseCase interface {
	GetProducts(page, pageSize, from, untill int, countries, categories []string) ([]entity.Product, int64, error)
	GetConcreteProduct(id int) (*entity.Product, error)
	UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error)
	DeleteProduct(id int) error
	AddProduct(product *entity.Product) (*entity.Product, error)
}

type CategoryUseCase interface {
	GetCategories() ([]entity.Category, error)
	UpdateCategory(id int, updates map[string]interface{}) (*entity.Category, error)
	DeleteCategory(id int) error
	AddCategory(product *entity.Category) (*entity.Category, error)
}

type CountryUseCase interface {
	GetCountries() ([]entity.Country, error)
	AddCountry(product *entity.Country) (*entity.Country, error)
	DeleteCountry(id int) error
	UpdateCategory(id int, updates map[string]interface{}) (*entity.Country, error)
}

type PictureUseCase interface {
	AddPictures(prodID int, url ...string) error
	DeletePicture(id int) error
	GetPictures(prodID int) ([]entity.Picture, error)
}

type Validator interface {
	Validate(dto interface{}) error
}

type Logger interface {
	Info(msg string, field ...any)
	Error(msg string, field ...any)
	Debug(msg string, field ...any)
}

type Client interface {
	Get(path string, queryParams ...map[string]string) (*http.Response, error)
}
