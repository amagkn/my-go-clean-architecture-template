package error_type

import "errors"

var (
	CategoryDoesNotExist = errors.New("category_does_not_exist")
	ProductDoesNotExist  = errors.New("product_does_not_exist")
)
