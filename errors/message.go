package errors

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Message(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Please, fill in this field"
	case "len":
		return lenError(err)
	case "min":
		return minError(err)
	default:
		return ""
	}
}

func lenError(err validator.FieldError) string {
	requiredValue := err.Param()

	return fmt.Sprintf("this field must have exactly %s characters", requiredValue)
}

func minError(err validator.FieldError) string {
	requiredValue := err.Param()
	return fmt.Sprintf("this field must have at least %s characters", requiredValue)
}
