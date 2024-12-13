package entity

import "time"

// Category представляет категорию товара
type Category struct {
	ID          int       `json:"id"`          // Уникальный идентификатор категории
	Name        string    `json:"name"`        // Название категории
	Description string    `json:"description"` // Описание категории
	CreatedAt   time.Time `json:"created_at"`  // Время создания записи
	UpdatedAt   time.Time `json:"updated_at"`  // Время обновления записи
}
