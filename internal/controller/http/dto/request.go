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
