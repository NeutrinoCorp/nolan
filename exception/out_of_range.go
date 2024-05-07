package exception

import (
	"errors"
	"fmt"
)

var ErrOutOfRange = errors.New("domain: property out of range")

// OutOfRange is a sentinel structure used allocate an ErrOutOfRange in a more detail manner.
type OutOfRange struct {
	PropertyName string
	A, B         any
	Actual       any
}

var _ Exception = OutOfRange{}

func (e OutOfRange) Error() string {
	return fmt.Sprintf("property '%s' has an invalid range [%v,%v], got <<%v>>", e.PropertyName, e.A, e.B, e.Actual)
}

func (e OutOfRange) Unwrap() error {
	return ErrOutOfRange
}
