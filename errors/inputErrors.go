package errors

import (
	"echo-test/types"

	"github.com/go-playground/validator/v10"
)

func PopulateErrorMap(errorsMap *(map[string]string), err error, validatedField types.IStep) *map[string]string {

	for _, err := range err.(validator.ValidationErrors) {
		(*errorsMap)[err.Field()] = Message(err, validatedField)
	}

	return errorsMap
}
