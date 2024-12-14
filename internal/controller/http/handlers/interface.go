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
}

type CountryUseCase interface {
	GetCountries() ([]entity.Country, error)
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
