package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
)

func (u *UseCase) UpdateProduct(ctx context.Context, input dto.UpdateProductInput) error {
	_, err := u.postgres.SelectOneCategory(ctx, *input.CategoryCode)
	if err != nil {
		if errors.Is(err, base_errors.NotFound) {
			return entity.ErrCategoryDoesNotExist
		}

		return base_errors.WithPath("u.postgres.SelectOneCategory", err)
	}

	_, err = u.postgres.SelectOneProduct(ctx, *input.ID)
	if err != nil {
		if errors.Is(err, base_errors.NotFound) {
			return entity.ErrProductDoesNotExist
		}

		return base_errors.WithPath("u.postgres.SelectOneProduct", err)
	}

	err = u.postgres.UpdateOneProduct(ctx, input)
	if err != nil {
		return base_errors.WithPath("u.postgres.UpdateOneProduct", err)
	}

	return nil
}
