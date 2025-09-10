package dto

type DeleteProductInput struct {
	ID string `json:"id" validate:"required,uuid"`
}
