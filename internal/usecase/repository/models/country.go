package models

import "time"

// Country представляет страну, связанную с товаром
type Country struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`          // Уникальный идентификатор страны
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`      // Название страны
	Code      string    `gorm:"type:varchar(3);not null;unique" json:"code"` // Код страны (ISO 3166-1 Alpha-3)
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`            // Время создания записи
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`            // Время обновления записи
}
