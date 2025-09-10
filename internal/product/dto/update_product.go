package dto

type UpdateProductInput struct {
	ID           *string `json:"id" validate:"required,uuid"`
	Name         *string `json:"name" validate:"omitempty,min=3"`
	Description  *string `json:"description" validate:"omitempty,min=3"`
	ImageUrl     *string `json:"imageUrl" validate:"omitempty,url"`
	CategoryCode *string `json:"categoryCode" validate:"required,numeric"`
}
