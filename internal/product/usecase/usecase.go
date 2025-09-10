package usecase

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
)

type Postgres interface {
	SelectManyProducts(ctx context.Context, categoryCode string) ([]entity.Product, error)
	SelectOneProduct(ctx context.Context, id string) (entity.Product, error)
	SelectOneCategory(ctx context.Context, categoryCode string) (entity.Category, error)
	InsertOneProduct(ctx context.Context, input dto.CreateProductInput) (entity.Product, error)
	UpdateOneProduct(ctx context.Context, input dto.UpdateProductInput) error
	DeleteOneProduct(ctx context.Context, id string) error
}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}
