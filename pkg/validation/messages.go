package validation

import (
	"fmt"
	"reflect"
)

type msg map[string]string

const unknownValidation = "unknown validation error"

var messages = msg{
	"numeric":  "must be numeric",
	"min":      "value is too short",
	"max":      "value is too long",
	"required": "is required",
	"url":      "must be a valid URL",
	"uuid":     "must be a valid uuid",
}

func (m *msg) WrongType(expected reflect.Type, got string) string {
	return fmt.Sprintf("expects a %s but got %s", expected, got)
}
