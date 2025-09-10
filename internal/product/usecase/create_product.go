package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/error_type"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
)

func (u *UseCase) CreateProduct(ctx context.Context, input dto.CreateProductInput) (dto.CreateProductOutput, error) {
	var output dto.CreateProductOutput

	_, err := u.postgres.SelectOneCategory(ctx, input.CategoryCode)
	if err != nil {
		if errors.Is(err, common_error.NotFound) {
			return output, error_type.CategoryDoesNotExist
		}

		return output, common_error.WithPath("u.postgres.SelectOneCategory", err)
	}

	createdProduct, err := u.postgres.InsertOneProduct(ctx, input)
	if err != nil {
		return output, common_error.WithPath("u.postgres.InsertOneProduct", err)
	}

	output.ID = createdProduct.ID

	return output, nil
}
