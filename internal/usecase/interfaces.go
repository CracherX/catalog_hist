package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	GetProducts(limit, offset, from, untill int, countries, categories []string) ([]entity.Product, error)
	CountRecords(from, untill int, countries, categories []string) (int64, error)
	GetProduct(id int) (*entity.Product, error)
	UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error)
	DeleteProduct(id int) (*entity.Product, error)
	AddProduct(product *entity.Product) (*entity.Product, error)
}

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
	UpdateCategory(id int, updates map[string]interface{}) (*entity.Category, error)
	DeleteCategory(id int) error
	AddCategory(product *entity.Category) (*entity.Category, error)
}

type CountryRepository interface {
	GetAllCountries() ([]entity.Country, error)
	AddCountry(product *entity.Country) (*entity.Country, error)
	DeleteCountry(id int) error
	UpdateCountry(id int, updates map[string]interface{}) (*entity.Country, error)
}

type PictureRepository interface {
	AddPictures(prodID int, url ...string) error
	DeletePicture(id int) error
	GetAllPictures(prodId int) ([]entity.Picture, error)
}
