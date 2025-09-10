package common_error

import (
	"errors"
	"fmt"
)

var (
	Validation     = errors.New("validation_error")
	InternalServer = errors.New("internal_server_error")
	InvalidJSON    = errors.New("json_body_error")
	NotFound       = errors.New("not_found")
)

func WithPath(path string, err error) error {
	return fmt.Errorf("%s: %w", path, err)
}
