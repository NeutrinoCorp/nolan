package exception

import (
	"errors"
	"fmt"
	"strings"
)

var ErrOneOf = errors.New("domain: property must have one of the expected values")

// OneOf is a sentinel structure used allocate an ErrOneOf in a more detail manner.
type OneOf struct {
	PropertyName string
	Expect       []string
	Value        any
}

var _ Exception = OneOf{}

func (o OneOf) Error() string {
	return fmt.Sprintf("property '%s' is <<%v>>, expected one of [%v]", o.PropertyName,
		o.Value, strings.Join(o.Expect, ","))
}

func (o OneOf) Unwrap() error {
	return ErrOneOf
}
