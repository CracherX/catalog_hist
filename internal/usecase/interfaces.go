package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	GetProducts(limit, offset, from, untill int, countries, categories []string) ([]entity.Product, error)
	CountRecords(from, untill int, countries, categories []string) (int64, error)
	GetProduct(id int) (*entity.Product, error)
}

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
}
