package dto

type GetProductsRequest struct {
	PageSize   string   `validate:"required,numeric"`
	Page       string   `validate:"required,numeric"`
	From       string   `validate:"required,len=4,numeric"`
	Until      string   `validate:"required,len=4,numeric"`
	Categories []string `validate:"required" json:"categories"`
	Countries  []string `validate:"required" json:"countries"`
}

type GetProductRequest struct {
	ID string `validate:"required,numeric"`
}
