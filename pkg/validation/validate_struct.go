package validation

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(input any) (map[string]string, error) {
	inputReflect := reflect.TypeOf(input)

	fields := map[string]string{}

	err := validate.Struct(input)
	if err != nil {
		var validateErrs validator.ValidationErrors

		if !errors.As(err, &validateErrs) {
			return nil, errWithPath("validate.Struct", err)
		}

		for _, e := range validateErrs {
			msg := messages[e.Tag()]

			if len(msg) == 0 {
				msg = unknownValidation
			}

			fields[getJsonTagByName(inputReflect, e.Field())] = msg
		}

		return fields, err
	}

	return nil, nil
}

func ValidateStructWithDecodeJSONBody(body io.ReadCloser, input any) (map[string]string, error) {
	fields := map[string]string{}

	err := json.NewDecoder(body).Decode(input)
	if err != nil {
		var unmarshalTypeErr *json.UnmarshalTypeError

		if !errors.As(err, &unmarshalTypeErr) {
			return nil, errWithPath("json.NewDecoder", err)
		}

		fields[unmarshalTypeErr.Field] = messages.WrongType(unmarshalTypeErr.Type, unmarshalTypeErr.Value)

		return fields, err
	}

	fields, err = ValidateStruct(input)
	if err != nil {
		return fields, errWithPath("validation.ValidateStruct", err)
	}

	return nil, nil
}

func getJsonTagByName(t reflect.Type, name string) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	field, _ := t.FieldByName(name)

	return field.Tag.Get("json")
}
