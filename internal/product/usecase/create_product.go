package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
)

func (u *UseCase) CreateProduct(ctx context.Context, input dto.CreateProductInput) (dto.CreateProductOutput, error) {
	var output dto.CreateProductOutput

	_, err := u.postgres.SelectOneCategory(ctx, input.CategoryCode)
	if err != nil {
		if errors.Is(err, base_errors.NotFound) {
			return output, entity.ErrCategoryDoesNotExist
		}

		return output, base_errors.WithPath("u.postgres.SelectOneCategory", err)
	}

	createdProduct, err := u.postgres.InsertOneProduct(ctx, input)
	if err != nil {
		return output, base_errors.WithPath("u.postgres.InsertOneProduct", err)
	}

	output.ID = createdProduct.ID

	return output, nil
}
