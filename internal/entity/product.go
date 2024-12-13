package entity

import "time"

// Product представляет товар в каталоге
type Product struct {
	ID          int       `json:"id"`          // Уникальный идентификатор товара
	Name        string    `json:"name"`        // Название товара
	Description string    `json:"description"` // Описание товара
	Price       float64   `json:"price"`       // Цена товара
	CountryID   int       `json:"country_id"`  // Внешний ключ на страну
	Country     Country   `json:"country"`     // Связь с таблицей стран
	CategoryID  int       `json:"category_id"` // Внешний ключ на категорию
	Category    Category  `json:"category"`    // Связь с таблицей категорий
	CreatedAt   time.Time `json:"created_at"`  // Время создания записи
	UpdatedAt   time.Time `json:"updated_at"`  // Время обновления записи
}
