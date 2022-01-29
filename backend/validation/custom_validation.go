package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorToText(e *validator.FieldError) string {
	err := *e
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", err.Field(), err.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", err.Field(), err.Param())
	case "email":
		return "invalid email"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", err.Field(), err.Param())
	}
	return fmt.Sprintf("%s is not valid", err.Field())
}

func FormatErrors(errors validator.ValidationErrors) (fErrors map[string][]string) {
	fErrors = make(map[string][]string)
	for _, err := range errors {
		name := strings.SplitN(err.Namespace(), ".", 2)[1]
		fErrors[name] = []string{ValidationErrorToText(&err)}
	}
	return
}
