package errors

import (
	"echo-test/types"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Message(err validator.FieldError, step types.IStep) string {
	switch err.Tag() {
	case "required":
		return "Please, fill in this field"
	case "len":
		return lenError(err)
	case "min":
		return minError(err)
	case "ltefield":
		return step.LteFieldError(err.Field())
	case "gt":
		return gtError(err)
	default:
		return fmt.Sprintf("This field must be valid %s", err.Tag())
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

func gtError(err validator.FieldError) string {
	if val, ok := err.Value().(int); ok {
		if val < 0 {
			return "This field cannot be negative"
		}
		return fmt.Sprintf("This field must be greater than %s", err.Param())
	}
	if err.Param() == "0" {
		return "This field cannot be empty"
	}

	return fmt.Sprintf("This field must have length greater than %s", err.Param())
}
