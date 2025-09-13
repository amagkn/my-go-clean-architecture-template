package postgres

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) UpdateOneProduct(ctx context.Context, input dto.UpdateProductInput) error {
	record := goqu.Record{}

	if input.Name != nil {
		record["name"] = *input.Name
	}
	if input.Description != nil {
		record["description"] = *input.Description
	}
	if input.ImageUrl != nil {
		record["image_url"] = *input.ImageUrl
	}
	if input.CategoryCode != nil {
		record["category_code"] = *input.CategoryCode
	}

	ds := goqu.Update("product").Set(record).Where(goqu.Ex{"id": *input.ID})

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
