package dto

type GetProductsRequest struct {
	PageSize string `validate:"required,numeric"`
	Page     string `validate:"required,numeric"`
	Category string `validate:"required"`
	Country  string `validate:"required"`
}
