package exception

import (
	"errors"
	"fmt"
)

var ErrArgumentRequired = errors.New("domain: argument missing")

// ArgumentRequired is a sentinel structure used allocate an ErrArgumentRequired in a more detail manner.
type ArgumentRequired struct {
	ArgName string
}

var _ Exception = ArgumentRequired{}

func (e ArgumentRequired) Error() string {
	return fmt.Sprintf("argument '%s' is required", e.ArgName)
}

func (e ArgumentRequired) Unwrap() error {
	return ErrArgumentRequired
}
