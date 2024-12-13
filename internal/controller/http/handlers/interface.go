package handlers

import "github.com/CracherX/catalog_hist/internal/entity"

type ProductUseCase interface {
	GetProducts(page, pageSize int, country, category string) ([]entity.Product, int64, error)
	GetConcreteProduct(id int) (*entity.Product, error)
}

type Validator interface {
	Validate(dto interface{}) error
}

type Logger interface {
	Info(msg string, field ...any)
	Error(msg string, field ...any)
	Debug(msg string, field ...any)
}
