package usecase

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
)

func (u *UseCase) GetProducts(ctx context.Context, input dto.GetProductsInput) ([]dto.GetProductsOutput, error) {
	var output []dto.GetProductsOutput

	products, err := u.postgres.SelectManyProducts(ctx, input.CategoryCode)
	if err != nil {
		return output, common_error.WithPath("u.postgres.SelectManyProducts", err)
	}

	for _, p := range products {
		output = append(output, dto.GetProductsOutput{
			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			ImageUrl:     p.ImageUrl,
			CategoryCode: p.CategoryCode,
		})
	}

	return output, nil
}
