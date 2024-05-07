package validate_test

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/exception"
	"github.com/neutrinocorp/nolan/validate"
)

func TestAdaptValidatorError(t *testing.T) {
	type someStruct struct {
		SomeString string  `validate:"required"`
		SomeFloat  float64 `validate:"gt=1"`
		SomeIP     string  `validate:"ip"`
		SomeAlpha  string  `validate:"alpha"`
		SomeOneOf  string  `validate:"oneof=REQUIRED PASSED UNKNOWN"`
	}

	val := someStruct{
		SomeString: "",
		SomeFloat:  0,
		SomeIP:     "someelver",
		SomeAlpha:  "somefoo99",
		SomeOneOf:  "somefoo99",
	}
	errValidate := validator.New().Struct(val)
	err := validate.AdaptValidatorError(errValidate)

	assert.True(t, errors.Is(err, exception.ErrOutOfRange))
	assert.True(t, errors.Is(err, exception.ErrArgumentRequired))
	assert.True(t, errors.Is(err, exception.ErrInvalidFormat))
	assert.True(t, errors.Is(err, exception.ErrOneOf))
}

func TestAdaptValidatorError_Cast(t *testing.T) {
	type someStruct struct {
		SomeString string  `validate:"required"`
		SomeFloat  float64 `validate:"gt=1"`
		SomeIP     string  `validate:"ip"`
		SomeAlpha  string  `validate:"alpha"`
		SomeOneOf  string  `validate:"oneof=REQUIRED PASSED UNKNOWN"`
	}
	val := someStruct{
		SomeString: "",
		SomeFloat:  0,
		SomeIP:     "someelver",
		SomeAlpha:  "somefoo99",
		SomeOneOf:  "somefoo99",
	}

	err := validate.GlobalValidator.Validate(val)
	var errFoo exception.InvalidFormat
	assert.True(t, errors.As(err, &errFoo))
}
