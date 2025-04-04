package dto

type GetProductsRequest struct {
	PageSize   string   `validate:"numeric"`
	Page       string   `validate:"numeric"`
	From       string   `validate:"required,len=4,numeric"`
	Until      string   `validate:"required,len=4,numeric"`
	Categories []string `json:"categories"`
	Countries  []string `json:"countries"`
}

type GetProductRequest struct {
	ID string `validate:"required,numeric"`
}

type PatchRequest struct {
	JWT     string                 `validate:"required,jwt"`
	ID      string                 `validate:"required,numeric"`
	Updates map[string]interface{} `validate:"required"`
}

type DeleteRequest struct {
	JWT string `validate:"required,jwt"`
	ID  string `validate:"required,numeric"`
}

type AddProductRequest struct {
	JWT         string  `validate:"required,jwt"`
	Name        string  `json:"name" validate:"required"`          // Название товара
	Description string  `json:"description"`                       // Описание товара
	Price       float64 `json:"price" validate:"required,numeric"` // Цена товара
	Year        int     `json:"year" validate:"required"`
	Picture     string  `json:"picture"`
	CountryId   int     `json:"country_id" validate:"required"`
	CategoryId  int     `json:"category_id" validate:"required"`
}

type AddCategoryRequest struct {
	JWT         string `validate:"required,jwt"`
	Name        string `json:"name" validate:"required"` // Название товара
	Description string `json:"description"`              // Описание товара
}

type AddCountryRequest struct {
	JWT  string `validate:"required,jwt"`
	Name string `json:"name" validate:"required"` // Название товара
	Code string `json:"code"`                     // Описание товара
}

type AddPicturesRequest struct {
	JWT        string   `validate:"required,jwt" json:"jwt"`
	PictureURL []string `validate:"required,min=1,dive" json:"pictureURL"`
	ProductID  int      `validate:"required" json:"productID"`
}

type GetPicturesRequest struct {
	ProductID string `validate:"required,numeric"`
}
