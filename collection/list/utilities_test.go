package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestIndexOf(t *testing.T) {
	var ls list.List[int] = &list.SliceList[int]{}
	ls.AddSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	tests := []struct {
		name string
		val  int
		exp  int
	}{
		{
			name: "not found",
			val:  0,
			exp:  -1,
		},
		{
			name: "mid",
			val:  5,
			exp:  4,
		},
		{
			name: "arbitrary",
			val:  7,
			exp:  6,
		},
		{
			name: "arbitrary",
			val:  2,
			exp:  1,
		},
		{
			name: "high",
			val:  10,
			exp:  9,
		},
		{
			name: "low",
			val:  1,
			exp:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := list.IndexOf(ls, tt.val)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestIndexOfOrdered(t *testing.T) {
	var ls list.List[int] = &list.SliceList[int]{}
	ls.AddSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	tests := []struct {
		name string
		val  int
		exp  int
	}{
		{
			name: "not found",
			val:  0,
			exp:  -1,
		},
		{
			name: "mid",
			val:  5,
			exp:  4,
		},
		{
			name: "arbitrary",
			val:  7,
			exp:  6,
		},
		{
			name: "arbitrary",
			val:  2,
			exp:  1,
		},
		{
			name: "high",
			val:  10,
			exp:  9,
		},
		{
			name: "low",
			val:  1,
			exp:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := list.IndexOfOrdered(ls, tt.val)
			assert.Equal(t, tt.exp, out)
		})
	}
}
