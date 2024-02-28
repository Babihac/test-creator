package errors

import (
	"echo-test/types"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func PopulateErrorMap(errorsMap *(map[string]string), err error, form types.IForm) *map[string]string {

	t := reflect.TypeOf(form).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		(*errorsMap)[jsonTag] = ""
	}

	for _, err := range err.(validator.ValidationErrors) {
		(*errorsMap)[err.Field()] = Message(err, form)
	}

	return errorsMap
}
