package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectOneProduct(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product

	ds := goqu.
		From("product").
		Select("id", "name", "description", "image_url").
		Where(goqu.Ex{"id": id}).
		Limit(1)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return product, base_errors.WithPath("ds.ToSQL", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return product, base_errors.NotFound
		}

		return product, base_errors.WithPath("p.pool.QueryRow.Scan", err)
	}

	return product, nil
}
