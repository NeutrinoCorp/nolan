package convert_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/convert"
)

func TestConvertMany(t *testing.T) {
	successFunc := func(a int) string {
		return fmt.Sprintf("%d", a)
	}

	tests := []struct {
		name          string
		in            []int
		inConvertFunc convert.ConverterFunc[int, string]
		exp           []string
		expPanic      bool
	}{
		{
			name:          "nil",
			in:            []int(nil),
			inConvertFunc: successFunc,
			exp:           []string(nil),
		},
		{
			name:          "empty",
			in:            []int{},
			inConvertFunc: successFunc,
			exp:           []string(nil),
		},
		{
			name:          "single",
			in:            []int{1},
			inConvertFunc: successFunc,
			exp:           []string{"1"},
		},
		{
			name:          "many",
			in:            []int{1, 2, 3, 4},
			inConvertFunc: successFunc,
			exp:           []string{"1", "2", "3", "4"},
		},
		{
			name:          "nil convert func",
			in:            []int{1},
			inConvertFunc: nil,
			exp:           []string(nil),
			expPanic:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				assert.Equal(t, tt.expPanic, r != nil)
			}()
			out := convert.ConvertMany[int, string](tt.in, tt.inConvertFunc)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestConvertManySafe(t *testing.T) {
	successFunc := func(a int) (string, error) {
		return fmt.Sprintf("%d", a), nil
	}
	failureFunc := func(a int) (string, error) {
		return "", fmt.Errorf("failure")
	}

	tests := []struct {
		name          string
		in            []int
		inConvertFunc convert.ConverterSafeFunc[int, string]
		exp           []string
		expErr        bool
		expPanic      bool
	}{
		{
			name:          "nil",
			in:            []int(nil),
			inConvertFunc: successFunc,
			exp:           []string(nil),
			expErr:        false,
		},
		{
			name:          "empty",
			in:            []int{},
			inConvertFunc: successFunc,
			exp:           []string(nil),
			expErr:        false,
		},
		{
			name:          "single",
			in:            []int{1},
			inConvertFunc: successFunc,
			exp:           []string{"1"},
			expErr:        false,
		},
		{
			name:          "many",
			in:            []int{1, 2, 3, 4},
			inConvertFunc: successFunc,
			exp:           []string{"1", "2", "3", "4"},
			expErr:        false,
		},
		{
			name:          "single fail",
			in:            []int{1},
			inConvertFunc: failureFunc,
			exp:           []string(nil),
			expErr:        true,
		},
		{
			name:          "many fail",
			in:            []int{1, 2, 3, 4},
			inConvertFunc: failureFunc,
			exp:           []string(nil),
			expErr:        true,
		},
		{
			name:          "nil convert func",
			in:            []int{1},
			inConvertFunc: nil,
			exp:           []string(nil),
			expErr:        false,
			expPanic:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				assert.Equal(t, tt.expPanic, r != nil)
			}()
			out, err := convert.ConvertManySafe[int, string](tt.in, tt.inConvertFunc)
			assert.Equal(t, tt.expErr, err != nil)
			assert.Equal(t, tt.exp, out)
		})
	}
}
