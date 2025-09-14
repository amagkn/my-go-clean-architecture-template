package base_errors

import (
	"errors"
	"fmt"
)

var (
	Validation     = errors.New("validation_error")
	InternalServer = errors.New("internal_server_error")
	InvalidJSON    = errors.New("invalid_json_body")
	NotFound       = errors.New("not_found")
)

func WithPath(path string, err error) error {
	return fmt.Errorf("%s: %w", path, err)
}
