package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/error_type"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
)

func (u *UseCase) DeleteProduct(ctx context.Context, id string) error {
	_, err := u.postgres.SelectOneProduct(ctx, id)
	if err != nil {
		if errors.Is(err, common_error.NotFound) {
			return error_type.ProductDoesNotExist
		}

		return common_error.WithPath("u.postgres.SelectOneProduct", err)
	}

	err = u.postgres.DeleteOneProduct(ctx, id)
	if err != nil {
		return common_error.WithPath("u.postgres.DeleteOneProduct", err)
	}

	return nil
}
