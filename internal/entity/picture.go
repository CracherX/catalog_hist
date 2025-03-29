package entity

type Picture struct {
	ID         int    `json:"id"`
	PictureURL string `json:"pictureURL"`
	ProductID  int    `json:"productID"`
}
