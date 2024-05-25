package values_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/values"
)

func TestNewPtr(t *testing.T) {
	someEmptyStr := ""
	someStr := "foo"
	tests := []struct {
		name string
		in   string
		exp  *string
	}{
		{
			name: "empty",
			in:   someEmptyStr,
			exp:  &someEmptyStr,
		},
		{
			name: "populated",
			in:   someStr,
			exp:  &someStr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := values.NewPtr(tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestNewNillablePtr(t *testing.T) {
	someStr := "foo"
	tests := []struct {
		name string
		in   string
		exp  *string
	}{
		{
			name: "empty",
			in:   "",
			exp:  nil,
		},
		{
			name: "populated",
			in:   someStr,
			exp:  &someStr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := values.NewNillablePtr(tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}
