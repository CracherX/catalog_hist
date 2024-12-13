package entity

import "time"

// Country представляет страну, связанную с товаром
type Country struct {
	ID        int       `json:"id"`         // Уникальный идентификатор страны
	Name      string    `json:"name"`       // Название страны
	Code      string    `json:"code"`       // Код страны (ISO 3166-1 Alpha-3)
	CreatedAt time.Time `json:"created_at"` // Время создания записи
	UpdatedAt time.Time `json:"updated_at"` // Время обновления записи
}
