package postgres

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) DeleteOneProduct(ctx context.Context, id string) error {
	ds := goqu.
		From("product").
		Where(goqu.Ex{"id": id}).
		Delete()

	sql, _, err := ds.ToSQL()
	if err != nil {
		return base_errors.WithPath("ds.ToSQL", err)
	}

	_, err = p.pool.Exec(ctx, sql)
	if err != nil {
		return base_errors.WithPath("p.pool.Exec", err)
	}

	return nil
}
