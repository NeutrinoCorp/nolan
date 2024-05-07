package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/list"
)

func TestContains(t *testing.T) {
	ls := &list.SliceList[int]{
		Source: make([]int, 0),
	}
	ls.AddSlice(0, 1, 2, 3, 4, 5, 6, 7, 8)
	tests := []struct {
		name string
		in   int
		exp  bool
	}{
		{
			name: "zero-value",
			in:   0,
			exp:  true,
		},
		{
			name: "found",
			in:   0,
			exp:  true,
		},
		{
			name: "not found",
			in:   -1,
			exp:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := collection.Contains[int](ls, tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestContainsAll(t *testing.T) {
	ls := &list.SliceList[int]{
		Source: make([]int, 0),
	}
	ls.AddSlice(0, 1, 2, 3, 4, 5, 6, 7, 8)
	tests := []struct {
		name string
		in   collection.Collection[int]
		exp  bool
	}{
		{
			name: "empty",
			in: &list.SliceList[int]{
				Source: []int{},
			},
			exp: true,
		},
		{
			name: "found",
			in: &list.SliceList[int]{
				Source: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			exp: true,
		},
		{
			name: "not found",
			in: &list.SliceList[int]{
				Source: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, -1},
			},
			exp: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := collection.ContainsAll[int](ls, tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestContainsSlice(t *testing.T) {
	ls := &list.SliceList[int]{
		Source: make([]int, 0),
	}
	ls.AddSlice(0, 1, 2, 3, 4, 5, 6, 7, 8)
	tests := []struct {
		name string
		in   []int
		exp  bool
	}{
		{
			name: "empty",
			in:   []int{},
			exp:  true,
		},
		{
			name: "found",
			in:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			exp:  true,
		},
		{
			name: "not found",
			in:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, -1},
			exp:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := collection.ContainsSlice[int](ls, tt.in)
			assert.Equal(t, tt.exp, out)
		})
	}
}
