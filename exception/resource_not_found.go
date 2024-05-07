package exception

import (
	"errors"
	"fmt"
)

var ErrResourceNotFound = errors.New("domain: resource not found")

// ResourceNotFound is a sentinel structure used allocate an ErrResourceNotFound in a more detail manner.
type ResourceNotFound struct {
	Resource string
}

var _ Exception = ResourceNotFound{}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("resource '%s' was not found", e.Resource)
}

func (e ResourceNotFound) Unwrap() error {
	return ErrResourceNotFound
}
