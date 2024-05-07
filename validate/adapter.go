package validate

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"

	"github.com/neutrinocorp/nolan/exception"
)

// AdaptValidatorError parses go-playground/validator errors into its domain exception counterparts as a single
// error instance using errors.Join routine.
func AdaptValidatorError(src error) error {
	var validateErrs validator.ValidationErrors
	if !errors.As(src, &validateErrs) {
		return src
	}

	errs := make([]error, 0, len(validateErrs))
	for _, srcErr := range validateErrs {
		field := strcase.ToSnake(srcErr.Field())
		switch srcErr.Tag() {
		case "required":
			errs = append(errs, exception.ArgumentRequired{ArgName: field})
		case "gte", "gt":
			errs = append(errs, exception.OutOfRange{PropertyName: field, A: srcErr.Param(), B: "n", Actual: srcErr.Value()})
		case "lt", "lte":
			errs = append(errs, exception.OutOfRange{PropertyName: field, A: 0, B: srcErr.Param(), Actual: srcErr.Value()})
		case "oneof":
			errs = append(errs, exception.OneOf{
				PropertyName: field,
				Expect:       strings.Split(srcErr.Param(), " "),
				Value:        srcErr.Value(),
			})
		default:
			errs = append(errs, exception.InvalidFormat{
				PropertyName: field,
				Expect:       srcErr.Tag(),
				Value:        srcErr.Value(),
			})
		}
	}

	return errors.Join(errs...)
}
