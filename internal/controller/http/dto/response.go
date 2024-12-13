package dto

import (
	"encoding/json"
	"github.com/CracherX/catalog_hist/internal/entity"
	"net/http"
)

type GetProductsResponse struct {
	Products []*GetProductResponse `json:"products"`
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"pageSize"`
	Category string                `json:"category"`
	Country  string                `json:"country"`
}

func ToProductsDTO(products []*GetProductResponse, total int64, page, pageSize int, category, country string) *GetProductsResponse {
	return &GetProductsResponse{
		Products: products,
		Category: category,
		Country:  country,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

type GetProductResponse struct {
	ID           int     `json:"id"`          // Уникальный идентификатор товара
	Name         string  `json:"name"`        // Название товара
	Description  string  `json:"description"` // Описание товара
	Price        float64 `json:"price"`       // Цена товара
	CountryName  string  `json:"country"`     // Название страны
	CategoryName string  `json:"category"`    // Название категории
}

func ToProductDTO(product *entity.Product) *GetProductResponse {
	return &GetProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		CountryName:  product.Country.Name,
		CategoryName: product.Category.Name,
	}
}

func ToProductDTOs(products []entity.Product) []*GetProductResponse {
	var responses []*GetProductResponse
	for _, product := range products {
		responses = append(responses, ToProductDTO(&product))
	}
	return responses
}

type e struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Response возвращает сообщение об успехе или ошибке клиенту в json формате.
func Response(w http.ResponseWriter, status int, errMsg string, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := e{
		Status:  status,
		Error:   http.StatusText(status),
		Message: errMsg,
	}
	if len(details) > 0 {
		errorResponse.Details = details[0]
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errorResponse)
}
