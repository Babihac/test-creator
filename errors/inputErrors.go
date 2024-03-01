package errors

import (
	"echo-test/types"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func PopulateErrorMap(errorsMap *(map[string]string), err error, form types.IForm) *map[string]string {

	for _, err := range err.(validator.ValidationErrors) {
		inputId := strings.NewReplacer("[", "-", "]", "").Replace(err.Field())
		(*errorsMap)[inputId] = Message(err, form)
	}

	return errorsMap
}

func ClearErrors(errorsMap *(map[string]string), request *http.Request) {
	for key, value := range request.PostForm {
		if len(value) == 1 {
			(*errorsMap)[key] = ""
			continue
		}

		for i := range value {
			inputId := strings.NewReplacer("[", "-", "]", fmt.Sprintf("-%d", i)).Replace(key)
			(*errorsMap)[inputId] = ""
		}
	}
}
