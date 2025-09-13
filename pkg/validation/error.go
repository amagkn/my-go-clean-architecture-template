package validation

import "fmt"

func errWithPath(path string, err error) error {
	return fmt.Errorf("%s: %w", path, err)
}
