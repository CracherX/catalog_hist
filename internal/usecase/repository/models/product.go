package models

import "time"

// Product представляет товар в каталоге
type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`       // Уникальный идентификатор товара
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`   // Название товара
	Description string    `gorm:"type:text;not null" json:"description"`    // Описание товара
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"` // Цена товара
	CountryID   int       `gorm:"not null" json:"country_id"`               // Внешний ключ на страну
	Country     Country   `gorm:"foreignKey:CountryID" json:"country"`      // Связь с таблицей стран
	CategoryID  int       `gorm:"not null" json:"category_id"`              // Внешний ключ на категорию
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`    // Связь с таблицей категорий
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`         // Время создания записи
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`         // Время обновления записи
}
