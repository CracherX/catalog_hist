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

type PatchProductRequest struct {
	JWT     string                 `validate:"required,jwt"`
	ID      string                 `validate:"required,numeric"`
	Updates map[string]interface{} `validate:"required"`
}

type DeleteProductRequest struct {
	JWT string `validate:"required,jwt"`
	ID  string `validate:"required,numeric"`
}

type AddProductRequest struct {
	JWT         string  `validate:"required,jwt"`
	Name        string  `json:"name" validate:"required"`          // Название товара
	Description string  `json:"description"`                       // Описание товара
	Price       float64 `json:"price" validate:"required,numeric"` // Цена товара
	Year        int     `json:"year" validate:"required"`
	CountryId   int     `json:"country_id" validate:"required"`
	CategoryId  int     `json:"category_id" validate:"required"`
}
