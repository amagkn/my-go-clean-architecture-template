package dto

type GetProductsInput struct {
	CategoryCode string `json:"categoryCode" validate:"required,numeric"`
}

type GetProductsOutput struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ImageUrl     string `json:"imageUrl"`
	CategoryCode int    `json:"categoryCode"`
}
