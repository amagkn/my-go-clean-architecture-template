package postgres

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectManyProducts(ctx context.Context, categoryCode string) ([]entity.Product, error) {
	var products []entity.Product

	ds := goqu.
		From("product").
		Select("id", "name", "description", "image_url", "category_code").
		Where(goqu.Ex{"category_code": categoryCode})

	sql, _, err := ds.ToSQL()
	if err != nil {
		return products, err
	}

	rows, err := p.pool.Query(ctx, sql)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl, &product.CategoryCode)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return products, err
	}

	return products, nil
}
