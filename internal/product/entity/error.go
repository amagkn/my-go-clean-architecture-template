package entity

import "errors"

var (
	ErrCategoryDoesNotExist = errors.New("category_does_not_exist")
	ErrProductDoesNotExist  = errors.New("product_does_not_exist")
)
