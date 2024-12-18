package models

import "time"

// Category представляет категорию товара
type Category struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`     // Уникальный идентификатор категории
	Name        string    `gorm:"type:varchar(255);not null" json:"name"` // Название категории
	Description string    `gorm:"type:text" json:"description"`           // Описание категории
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`       // Время создания записи
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`       // Время обновления записи
}
