package entity

import "time"

// Product представляет товар в каталоге
type Product struct {
	ID          int       `json:"id"`          // Уникальный идентификатор товара
	Name        string    `json:"name"`        // Название товара
	Description string    `json:"description"` // Описание товара
	Price       float64   `json:"price"`       // Цена товара
	CreatedAt   time.Time `json:"created_at"`  // Время создания записи
	UpdatedAt   time.Time `json:"updated_at"`  // Время обновления записи
}

// Category представляет категорию товара
type Category struct {
	ID          int       `json:"id"`          // Уникальный идентификатор категории
	Name        string    `json:"name"`        // Название категории
	Description string    `json:"description"` // Описание категории
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
