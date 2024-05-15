package values_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/values"
)

func TestCoalesce_Nillable(t *testing.T) {
	type someStr struct{}
	var someNilStr *someStr
	somePopulatedStr := &someStr{}

	tests := []struct {
		name string
		in   []*someStr
		exp  *someStr
	}{
		{
			name: "nil",
			in:   []*someStr(nil),
			exp:  nil,
		},
		{
			name: "zero val",
			in:   []*someStr{someNilStr, someNilStr, someNilStr, someNilStr},
			exp:  nil,
		},
		{
			name: "last val",
			in:   []*someStr{someNilStr, someNilStr, someNilStr, somePopulatedStr},
			exp:  somePopulatedStr,
		},
		{
			name: "all val",
			in:   []*someStr{somePopulatedStr, somePopulatedStr, somePopulatedStr, somePopulatedStr},
			exp:  somePopulatedStr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := values.Coalesce[*someStr](tt.in...)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestCoalesce_ZeroValue(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  int
	}{
		{
			name: "nil",
			in:   []int(nil),
			exp:  0,
		},
		{
			name: "zero val",
			in:   []int{0, 0, 0, 0},
			exp:  0,
		},
		{
			name: "last val",
			in:   []int{0, 0, 0, 1},
			exp:  1,
		},
		{
			name: "last val",
			in:   []int{0, 0, 0, -1},
			exp:  -1,
		},
		{
			name: "all val",
			in:   []int{4, 3, 2, 1},
			exp:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := values.Coalesce[int](tt.in...)
			assert.Equal(t, tt.exp, out)
		})
	}
}
