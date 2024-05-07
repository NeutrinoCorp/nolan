package exception_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/exception"
)

func TestUnwrapErrors(t *testing.T) {
	var ErrSomeFoo = errors.New("some foo")

	tests := []struct {
		name string
		in   error
		exp  []error
	}{
		{
			name: "nil",
			in:   nil,
			exp:  nil,
		},
		{
			name: "arbitrary",
			in:   ErrSomeFoo,
			exp:  []error{ErrSomeFoo},
		},
		{
			name: "single",
			in:   errors.Join(ErrSomeFoo),
			exp:  []error{ErrSomeFoo},
		},
		{
			name: "multi",
			in:   errors.Join(ErrSomeFoo, errors.New("some error")),
			exp:  []error{ErrSomeFoo, errors.New("some error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := exception.UnwrapErrors(tt.in)
			assert.EqualValues(t, tt.exp, out)
		})
	}
}
