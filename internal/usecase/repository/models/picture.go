package models

type Picture struct {
	ID         int     `gorm:"primaryKey;autoIncrement" json:"id"`
	PictureURL string  `gorm:"type:varchar(255);not null" json:"pictureURL"`
	Product    Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	ProductID  int     `gorm:"not null" json:"productID"`
}
