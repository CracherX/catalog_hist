package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	GetProducts(limit, offset, from, untill int, countries, categories []string) ([]entity.Product, error)
	CountRecords(from, untill int, countries, categories []string) (int64, error)
	GetProduct(id int) (*entity.Product, error)
	UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error)
	DeleteProduct(id int) error
	AddProduct(product *entity.Product) (*entity.Product, error)
}

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
}

type CountryRepository interface {
	GetAllCountries() ([]entity.Country, error)
}
