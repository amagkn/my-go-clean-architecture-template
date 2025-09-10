package postgres

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (p *Postgres) InsertOneProduct(ctx context.Context, input dto.CreateProductInput) (entity.Product, error) {
	var product entity.Product

	ds := goqu.Insert("product").
		Rows(goqu.Record{
			//todo: Сделать генерацию id на уровне базы данных
			"id":            uuid.New(),
			"name":          input.Name,
			"description":   input.Description,
			"image_url":     input.ImageUrl,
			"category_code": input.CategoryCode,
		}).
		Returning("id", "name", "description", "image_url", "category_code")

	sql, args, err := ds.ToSQL()
	if err != nil {
		return product, common_error.WithPath("ds.ToSQL", err)
	}

	row := p.pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl, &product.CategoryCode)
	if err != nil {
		return product, common_error.WithPath("row.Scan", err)
	}

	return product, nil
}
