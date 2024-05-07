package exception

import (
	"errors"
	"fmt"
)

var ErrInvalidFormat = errors.New("domain: property has an invalid format")

// InvalidFormat is a sentinel structure used allocate an ErrInvalidFormat in a more detail manner.
type InvalidFormat struct {
	PropertyName string
	Expect       any
	Value        any
}

var _ Exception = InvalidFormat{}

func (e InvalidFormat) Error() string {
	return fmt.Sprintf("property '%s' has an invalid format, expected '%v', got <<%v>>", e.PropertyName, e.Expect, e.Value)
}

func (e InvalidFormat) Unwrap() error {
	return ErrInvalidFormat
}
