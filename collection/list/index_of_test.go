package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestIndexOf(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name        string
		src         []int
		factoryFunc func() list.List[int]
		val         int
		exp         int
	}{
		{
			name:        "linked_ls not found",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         0,
			exp:         -1,
		},
		{
			name:        "linked_ls mid",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         5,
			exp:         4,
		},
		{
			name:        "linked_ls arbitrary",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         7,
			exp:         6,
		},
		{
			name:        "linked_ls arbitrary",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         2,
			exp:         1,
		},
		{
			name:        "linked_ls high",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         10,
			exp:         9,
		},
		{
			name:        "linked_ls low",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: linkedLsFactoryFunc,
			val:         1,
			exp:         0,
		},
		// slice
		{
			name:        "slice not found",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         0,
			exp:         -1,
		},
		{
			name:        "slice mid",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         5,
			exp:         4,
		},
		{
			name:        "slice arbitrary",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         7,
			exp:         6,
		},
		{
			name:        "slice arbitrary",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         2,
			exp:         1,
		},
		{
			name:        "slice high",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         10,
			exp:         9,
		},
		{
			name:        "slice low",
			src:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			factoryFunc: sliceLsFactoryFunc,
			val:         1,
			exp:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.src...)
			out := list.IndexOf(ls, tt.val)
			assert.Equal(t, tt.exp, out)

			out = list.IndexOfOrdered(ls, tt.val)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func BenchmarkIndexOf_Slice(b *testing.B) {
	ls := list.NewSliceList[int](nil)
	lim := 999999
	for i := 0; i < lim; i++ {
		ls.Add(i)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = list.IndexOf[int](ls, lim/4)
	}
}

func BenchmarkIndexOf_LinkedList(b *testing.B) {
	ls := list.NewDoublyLinkedList[int]()
	lim := 999999
	for i := 0; i < lim; i++ {
		ls.Add(i)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = list.IndexOf[int](ls, lim/4)
	}
}

func BenchmarkIndexOfOrdered_Slice(b *testing.B) {
	ls := list.NewSliceList[int](nil)
	lim := 999999
	for i := 0; i < lim; i++ {
		ls.Add(i)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = list.IndexOfOrdered[int](ls, lim/4)
	}
}

func BenchmarkIndexOfOrdered_LinkedList(b *testing.B) {
	ls := list.NewDoublyLinkedList[int]()
	lim := 999999
	for i := 0; i < lim; i++ {
		ls.Add(i)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = list.IndexOfOrdered[int](ls, lim/4)
	}
}
