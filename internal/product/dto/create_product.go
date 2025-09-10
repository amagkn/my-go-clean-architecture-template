package dto

type CreateProductInput struct {
	Name         string `json:"name" validate:"required,min=3"`
	Description  string `json:"description" validate:"required,min=3"`
	ImageUrl     string `json:"imageUrl" validate:"required,url"`
	CategoryCode string `json:"categoryCode" validate:"required,numeric"`
}

type CreateProductOutput struct {
	ID string `json:"id"`
}
