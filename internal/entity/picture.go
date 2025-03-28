package entity

type Picture struct {
	ID         int     `json:"id"`
	PictureURL string  `json:"pictureURL"`
	Product    Product `json:"product"`
	ProductID  int     `json:"productID"`
}
