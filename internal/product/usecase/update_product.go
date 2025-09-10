package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/error_type"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
)

func (u *UseCase) UpdateProduct(ctx context.Context, input dto.UpdateProductInput) error {
	_, err := u.postgres.SelectOneCategory(ctx, *input.CategoryCode)
	if err != nil {
		if errors.Is(err, common_error.NotFound) {
			return error_type.CategoryDoesNotExist
		}

		return common_error.WithPath("u.postgres.SelectOneCategory", err)
	}

	_, err = u.postgres.SelectOneProduct(ctx, *input.ID)
	if err != nil {
		if errors.Is(err, common_error.NotFound) {
			return error_type.ProductDoesNotExist
		}

		return common_error.WithPath("u.postgres.SelectOneProduct", err)
	}

	err = u.postgres.UpdateOneProduct(ctx, input)
	if err != nil {
		return common_error.WithPath("u.postgres.UpdateOneProduct", err)
	}

	return nil
}
