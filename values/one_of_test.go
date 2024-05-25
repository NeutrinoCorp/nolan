package values_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/values"
)

func TestIsOneOf(t *testing.T) {
	tests := []struct {
		name       string
		inBaseCase string
		inOneOf    []string
		exp        bool
	}{
		{
			name:       "both empty",
			inBaseCase: "",
			inOneOf:    nil,
			exp:        false,
		},
		{
			name:       "one of nil",
			inBaseCase: "foo",
			inOneOf:    nil,
			exp:        false,
		},
		{
			name:       "base empty",
			inBaseCase: "",
			inOneOf:    []string{"foo"},
			exp:        false,
		},
		{
			name:       "single not equal",
			inBaseCase: "foo",
			inOneOf:    []string{"bar"},
			exp:        false,
		},
		{
			name:       "multi not equal",
			inBaseCase: "foo",
			inOneOf:    []string{"bar", "baz", "foobar"},
			exp:        false,
		},
		{
			name:       "single equal",
			inBaseCase: "foo",
			inOneOf:    []string{"foo"},
			exp:        true,
		},
		{
			name:       "multi one equal",
			inBaseCase: "foo",
			inOneOf:    []string{"bar", "baz", "foo"},
			exp:        true,
		},
		{
			name:       "multi all equal",
			inBaseCase: "foo",
			inOneOf:    []string{"foo", "foo", "foo"},
			exp:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := values.IsOneOf(tt.inBaseCase, tt.inOneOf...)
			assert.Equal(t, tt.exp, out)
		})
	}
}
