package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) SelectOneCategory(ctx context.Context, categoryCode string) (entity.Category, error) {
	var category entity.Category

	ds := goqu.
		From("category").
		Select("name", "code").
		Where(goqu.Ex{"code": categoryCode}).
		Limit(1)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return category, common_error.WithPath("ds.ToSQL", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&category.Name, &category.Code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return category, common_error.NotFound
		}

		return category, common_error.WithPath("p.pool.QueryRow.Scan", err)
	}

	return category, nil
}
