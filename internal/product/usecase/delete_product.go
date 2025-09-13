package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
)

func (u *UseCase) DeleteProduct(ctx context.Context, id string) error {
	_, err := u.postgres.SelectOneProduct(ctx, id)
	if err != nil {
		if errors.Is(err, base_errors.NotFound) {
			return entity.ErrProductDoesNotExist
		}

		return base_errors.WithPath("u.postgres.SelectOneProduct", err)
	}

	err = u.postgres.DeleteOneProduct(ctx, id)
	if err != nil {
		return base_errors.WithPath("u.postgres.DeleteOneProduct", err)
	}

	return nil
}
