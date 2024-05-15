package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/neutrinocorp/nolan/collection/list"
)

// This file contains all tests for basic list.List implementations. We do this merge
// to keep nolan's List API homogeneous and thus, consistent, deterministic and predictable.

func TestList_Add(t *testing.T) {
	linkedLsCollFunc := func(src []int) list.List[int] {
		tmp := list.NewSliceList[int](src)
		return list.NewDoublyLinkedListFromCollection[int](tmp)
	}
	linkedLsSlcFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	linkedLsFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}

	sliceCollFunc := func(src []int) list.List[int] {
		tmp := list.NewDoublyLinkedListFromSlice[int](src)
		return list.NewSliceListFromCollection[int](tmp)
	}
	sliceFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		in          []int
		src         []int
		factoryFunc func(src []int) list.List[int]
		exp         []int
	}{
		// linked lists
		{
			name:        "linked_list_coll nil",
			in:          nil,
			src:         nil,
			factoryFunc: linkedLsCollFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list_coll empty",
			in:          []int{},
			src:         []int{},
			factoryFunc: linkedLsCollFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list_coll single",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: linkedLsCollFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "linked_list_coll multi",
			in:          []int{4, 5},
			src:         []int{1, 2, 3},
			factoryFunc: linkedLsCollFunc,
			exp:         []int{1, 2, 3, 4, 5},
		},
		{
			name:        "linked_list_slc nil",
			in:          nil,
			src:         nil,
			factoryFunc: linkedLsSlcFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list_slc empty",
			in:          []int{},
			factoryFunc: linkedLsSlcFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list_slc single",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: linkedLsSlcFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "linked_list_slc multi",
			in:          []int{4, 5},
			src:         []int{1, 2, 3},
			factoryFunc: linkedLsSlcFunc,
			exp:         []int{1, 2, 3, 4, 5},
		},
		{
			name:        "linked_list nil",
			in:          nil,
			src:         nil,
			factoryFunc: linkedLsFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list empty",
			in:          []int{},
			factoryFunc: linkedLsFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_list single",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: linkedLsFunc,
			exp:         []int{2},
		},
		{
			name:        "linked_list multi",
			in:          []int{4, 5},
			src:         []int{1, 2, 3},
			factoryFunc: linkedLsFunc,
			exp:         []int{4, 5},
		},
		// slice-backed
		{
			name:        "slice_coll nil",
			in:          nil,
			src:         nil,
			factoryFunc: sliceCollFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice_coll empty",
			in:          []int{},
			src:         []int{},
			factoryFunc: sliceCollFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice_coll single",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: sliceCollFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "slice_coll multi",
			in:          []int{4, 5},
			src:         []int{1, 2, 3},
			factoryFunc: sliceCollFunc,
			exp:         []int{1, 2, 3, 4, 5},
		},
		{
			name:        "slice nil",
			in:          nil,
			src:         nil,
			factoryFunc: sliceFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: sliceFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "slice multi",
			in:          []int{4, 5},
			src:         []int{1, 2, 3},
			factoryFunc: sliceFunc,
			exp:         []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ls list.List[int]
			if tt.factoryFunc != nil {
				ls = tt.factoryFunc(tt.src)
			}
			for _, item := range tt.in {
				ls.Add(item)
			}
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_AddAll(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name        string
		in          []int
		factoryFunc func() list.List[int]
		exp         []int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			in:          nil,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls empty",
			in:          []int{},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single",
			in:          []int{1},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1},
		},
		{
			name:        "linked_ls multi",
			in:          []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2, 3},
		},
		// slice
		{
			name:        "slice nil",
			in:          nil,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single",
			in:          []int{1},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1},
		},
		{
			name:        "slice multi",
			in:          []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			tmp := list.NewSliceList(tt.in)
			ls.AddAll(tmp)
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_AddSlice(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name        string
		in          []int
		factoryFunc func() list.List[int]
		exp         []int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			in:          nil,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls empty",
			in:          []int{},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single",
			in:          []int{1},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1},
		},
		{
			name:        "linked_ls multi",
			in:          []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2, 3},
		},
		// slice
		{
			name:        "slice nil",
			in:          nil,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single",
			in:          []int{1},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1},
		},
		{
			name:        "slice multi",
			in:          []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_Clear(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name        string
		in          []int
		factoryFunc func() list.List[int]
	}{
		// linked list
		{
			name:        "linked_ls nil",
			in:          nil,
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls empty",
			in:          []int{},
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls single",
			in:          []int{1},
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls multi",
			in:          []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
		},
		// slice
		{
			name:        "slice nil",
			in:          nil,
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice single",
			in:          []int{1},
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice multi",
			in:          []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			require.Equal(t, len(tt.in), ls.Len())
			ls.Clear()
			assert.Equal(t, 0, ls.Len())
			assert.True(t, ls.IsEmpty())
			assert.Len(t, ls.ToSlice(), 0)
		})
	}
}

func TestList_IsEmpty(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name        string
		in          []int
		factoryFunc func() list.List[int]
	}{
		// linked list
		{
			name:        "linked_ls nil",
			in:          nil,
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls empty",
			in:          []int{},
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls single",
			in:          []int{1},
			factoryFunc: linkedLsFactoryFunc,
		},
		{
			name:        "linked_ls multi",
			in:          []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
		},
		// slice
		{
			name:        "slice nil",
			in:          nil,
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice single",
			in:          []int{1},
			factoryFunc: sliceLsFactoryFunc,
		},
		{
			name:        "slice multi",
			in:          []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			require.Equal(t, len(tt.in), ls.Len())
			assert.Equal(t, len(tt.in) == 0, ls.IsEmpty())
		})
	}
}

func TestList_ForEach(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name         string
		in           []int
		factoryFunc  func() list.List[int]
		breakAtIndex int
		exp          []int
	}{
		// linked list
		{
			name:         "linked_ls nil",
			in:           nil,
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "linked_ls empty",
			in:           []int{},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "linked_ls single",
			in:           []int{1},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1},
		},
		{
			name:         "linked_ls multi",
			in:           []int{1, 2, 3},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1, 2, 3},
		},
		{
			name:         "linked_ls multi break",
			in:           []int{1, 2, 3},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: 1,
			exp:          []int{1, 2},
		},
		// slice
		{
			name:         "slice nil",
			in:           nil,
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "slice empty",
			in:           []int{},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "slice single",
			in:           []int{1},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1},
		},
		{
			name:         "slice multi",
			in:           []int{1, 2, 3},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1, 2, 3},
		},
		{
			name:         "slice multi break",
			in:           []int{1, 2, 3},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: 1,
			exp:          []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			if tt.breakAtIndex < 0 {
				require.Equal(t, len(tt.exp), ls.Len())
			}

			i := 0
			buf := make([]int, 0)
			out := ls.ToSlice()
			ls.ForEach(func(a int) bool {
				buf = append(buf, out[i])
				if tt.breakAtIndex == i {
					return true
				}
				i++
				return false
			})

			if out != nil {
				assert.Equal(t, tt.exp, buf)
			}
		})
	}
}

func TestList_AddAt(t *testing.T) {
	linkedLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	sliceLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		src         []int
		addAtIndex  int
		in          int
		factoryFunc func([]int) list.List[int]
		exp         []int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			src:         nil,
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls empty",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single merge",
			src:         []int{1},
			in:          2,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "linked_ls multi",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls multi merge head",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 4, 2, 3},
		},
		{
			name:        "linked_ls multi merge non-zero idx",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  1,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2, 4, 3},
		},
		{
			name:        "linked_ls multi merge append",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  2,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2, 3, 4},
		},
		// slice
		{
			name:        "slice nil",
			src:         nil,
			in:          0,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single merge",
			src:         []int{1},
			in:          2,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "slice multi",
			src:         make([]int, 0),
			in:          1,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice multi merge head",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 4, 2, 3},
		},
		{
			name:        "slice multi merge non-zero idx",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  1,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2, 4, 3},
		},
		{
			name:        "slice multi merge append",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  2,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc(tt.src)
			ls.AddAt(tt.addAtIndex, tt.in)
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_AddAllAt(t *testing.T) {
	linkedLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	sliceLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		in          []int
		src         []int
		addAtIndex  int
		factoryFunc func([]int) list.List[int]
		exp         []int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			in:          nil,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls empty",
			in:          []int{},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single",
			in:          []int{1},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single merge",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "linked_ls multi",
			in:          []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls multi merge head",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 4, 5, 6, 2, 3},
		},
		{
			name:        "linked_ls multi merge non-zero idx",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
			addAtIndex:  1,
			exp:         []int{1, 2, 4, 5, 6, 3},
		},
		{
			name:        "linked_ls multi merge append",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			factoryFunc: linkedLsFactoryFunc,
			addAtIndex:  2,
			exp:         []int{1, 2, 3, 4, 5, 6},
		},
		// slice
		{
			name:        "slice nil",
			in:          nil,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			in:          []int{},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single",
			in:          []int{1},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single merge",
			in:          []int{2},
			src:         []int{1},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2},
		},
		{
			name:        "slice multi",
			in:          []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice multi merge head",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 4, 5, 6, 2, 3},
		},
		{
			name:        "slice multi merge non-zero idx",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
			addAtIndex:  1,
			exp:         []int{1, 2, 4, 5, 6, 3},
		},
		{
			name:        "slice multi merge append",
			in:          []int{4, 5, 6},
			src:         []int{1, 2, 3},
			factoryFunc: sliceLsFactoryFunc,
			addAtIndex:  2,
			exp:         []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc(tt.src)
			tmp := list.NewSliceList(tt.in)
			ls.AddAllAt(tt.addAtIndex, tmp)
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_SetAt(t *testing.T) {
	linkedLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	sliceLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		src         []int
		addAtIndex  int
		in          int
		factoryFunc func([]int) list.List[int]
		exp         []int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			src:         nil,
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls empty",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls single merge",
			src:         []int{1},
			in:          2,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{2},
		},
		{
			name:        "linked_ls multi",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "linked_ls multi merge head",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{4, 2, 3},
		},
		{
			name:        "linked_ls multi merge non-zero idx",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  1,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 4, 3},
		},
		{
			name:        "linked_ls multi merge append",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  2,
			factoryFunc: linkedLsFactoryFunc,
			exp:         []int{1, 2, 4},
		},
		// slice
		{
			name:        "slice nil",
			src:         nil,
			in:          0,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice empty",
			src:         make([]int, 0),
			in:          0,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice single merge",
			src:         []int{1},
			in:          2,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{2},
		},
		{
			name:        "slice multi",
			src:         make([]int, 0),
			in:          1,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int(nil),
		},
		{
			name:        "slice multi merge head",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{4, 2, 3},
		},
		{
			name:        "slice multi merge non-zero idx",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  1,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 4, 3},
		},
		{
			name:        "slice multi merge append",
			src:         []int{1, 2, 3},
			in:          4,
			addAtIndex:  2,
			factoryFunc: sliceLsFactoryFunc,
			exp:         []int{1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc(tt.src)
			ls.SetAt(tt.addAtIndex, tt.in)
			require.Equal(t, len(tt.exp), ls.Len())
			assert.Equal(t, tt.exp, ls.ToSlice())
		})
	}
}

func TestList_GetAt(t *testing.T) {
	linkedLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	sliceLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		src         []int
		addAtIndex  int
		factoryFunc func([]int) list.List[int]
		exp         int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			src:         nil,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls empty",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls single merge",
			src:         []int{1},
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "linked_ls multi",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls multi merge head",
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "linked_ls multi merge non-zero idx",
			src:         []int{1, 2, 3},
			addAtIndex:  1,
			factoryFunc: linkedLsFactoryFunc,
			exp:         2,
		},
		{
			name:        "linked_ls multi merge append",
			src:         []int{1, 2, 3},
			addAtIndex:  2,
			factoryFunc: linkedLsFactoryFunc,
			exp:         3,
		},
		// slice
		{
			name:        "slice nil",
			src:         nil,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice empty",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice single merge",
			src:         []int{1},
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "slice multi",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice multi merge head",
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "slice multi merge non-zero idx",
			src:         []int{1, 2, 3},
			addAtIndex:  1,
			factoryFunc: sliceLsFactoryFunc,
			exp:         2,
		},
		{
			name:        "slice multi merge append",
			src:         []int{1, 2, 3},
			addAtIndex:  2,
			factoryFunc: sliceLsFactoryFunc,
			exp:         3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc(tt.src)
			out := ls.GetAt(tt.addAtIndex)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestList_RemoveAt(t *testing.T) {
	linkedLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewDoublyLinkedListFromSlice[int](src)
	}
	sliceLsFactoryFunc := func(src []int) list.List[int] {
		return list.NewSliceList[int](src)
	}

	tests := []struct {
		name        string
		src         []int
		addAtIndex  int
		factoryFunc func([]int) list.List[int]
		exp         int
	}{
		// linked list
		{
			name:        "linked_ls nil",
			src:         nil,
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls empty",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls single merge",
			src:         []int{1},
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "linked_ls multi",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "linked_ls multi merge head",
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: linkedLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "linked_ls multi merge non-zero idx",
			src:         []int{1, 2, 3},
			addAtIndex:  1,
			factoryFunc: linkedLsFactoryFunc,
			exp:         2,
		},
		{
			name:        "linked_ls multi merge append",
			src:         []int{1, 2, 3},
			addAtIndex:  2,
			factoryFunc: linkedLsFactoryFunc,
			exp:         3,
		},
		// slice
		{
			name:        "slice nil",
			src:         nil,
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice empty",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice single merge",
			src:         []int{1},
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "slice multi",
			src:         make([]int, 0),
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         0,
		},
		{
			name:        "slice multi merge head",
			src:         []int{1, 2, 3},
			addAtIndex:  0,
			factoryFunc: sliceLsFactoryFunc,
			exp:         1,
		},
		{
			name:        "slice multi merge non-zero idx",
			src:         []int{1, 2, 3},
			addAtIndex:  1,
			factoryFunc: sliceLsFactoryFunc,
			exp:         2,
		},
		{
			name:        "slice multi merge append",
			src:         []int{1, 2, 3},
			addAtIndex:  2,
			factoryFunc: sliceLsFactoryFunc,
			exp:         3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc(tt.src)
			snapshotLen := ls.Len()
			out := ls.RemoveAt(tt.addAtIndex)
			assert.Equal(t, tt.exp, out)
			if len(tt.src) == 0 {
				return
			}
			assert.Equal(t, snapshotLen-1, ls.Len())
		})
	}
}

func TestList_ToSubList(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	slcLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name           string
		in             []int
		subListIndexes [2]int
		factoryFunc    func() list.List[int]
		exp            []int
	}{
		// linked list
		{
			name:           "linked_ls nil",
			in:             nil,
			subListIndexes: [2]int{},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "linked_ls empty",
			in:             make([]int, 0),
			subListIndexes: [2]int{},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		}, {
			name:           "linked_ls single",
			in:             []int{1},
			subListIndexes: [2]int{0},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            []int{1},
		},
		{
			name:           "linked_ls single arbitrary",
			in:             []int{1},
			subListIndexes: [2]int{-1, 0},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "linked_ls single out of bounds",
			in:             []int{1},
			subListIndexes: [2]int{0, 1},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "linked_ls multi head",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            []int{1},
		},
		{
			name:           "linked_ls multi mid",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 1},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            []int{1, 2},
		},
		{
			name:           "linked_ls multi tail",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 2},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            []int{1, 2, 3},
		},
		{
			name:           "linked_ls multi arbitrary",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{-1, 2},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "linked_ls multi out of bounds",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 5},
			factoryFunc:    linkedLsFactoryFunc,
			exp:            nil,
		},
		// slice
		{
			name:           "slice_ls nil",
			in:             nil,
			subListIndexes: [2]int{},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "slice_ls empty",
			in:             make([]int, 0),
			subListIndexes: [2]int{},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "slice_ls single",
			in:             []int{1},
			subListIndexes: [2]int{0},
			factoryFunc:    slcLsFactoryFunc,
			exp:            []int{1},
		},
		{
			name:           "slice_ls single arbitrary",
			in:             []int{1},
			subListIndexes: [2]int{-1, 0},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "slice_ls single out of bounds",
			in:             []int{1},
			subListIndexes: [2]int{0, 1},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "slice_ls multi head",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0},
			factoryFunc:    slcLsFactoryFunc,
			exp:            []int{1},
		},
		{
			name:           "slice_ls multi mid",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 1},
			factoryFunc:    slcLsFactoryFunc,
			exp:            []int{1, 2},
		},
		{
			name:           "slice_ls multi tail",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 2},
			factoryFunc:    slcLsFactoryFunc,
			exp:            []int{1, 2, 3},
		},
		{
			name:           "slice_ls multi arbitrary",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{-1, 2},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
		{
			name:           "slice_ls multi out of bounds",
			in:             []int{1, 2, 3},
			subListIndexes: [2]int{0, 5},
			factoryFunc:    slcLsFactoryFunc,
			exp:            nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			subLs := ls.ToSubList(tt.subListIndexes[0], tt.subListIndexes[1])
			if len(tt.exp) == 0 {
				assert.Nil(t, subLs)
				return
			}
			assert.Equal(t, tt.exp, subLs.ToSlice())
		})
	}
}

func TestList_ForEachWithIndex(t *testing.T) {
	linkedLsFactoryFunc := func() list.List[int] {
		return list.NewDoublyLinkedList[int]()
	}
	sliceLsFactoryFunc := func() list.List[int] {
		return list.NewSliceList[int](nil)
	}

	tests := []struct {
		name         string
		in           []int
		factoryFunc  func() list.List[int]
		breakAtIndex int
		exp          []int
	}{
		// linked list
		{
			name:         "linked_ls nil",
			in:           nil,
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "linked_ls empty",
			in:           []int{},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "linked_ls single",
			in:           []int{1},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1},
		},
		{
			name:         "linked_ls multi",
			in:           []int{1, 2, 3},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1, 2, 3},
		},
		{
			name:         "linked_ls multi break",
			in:           []int{1, 2, 3},
			factoryFunc:  linkedLsFactoryFunc,
			breakAtIndex: 1,
			exp:          []int{1, 2},
		},
		// slice
		{
			name:         "slice nil",
			in:           nil,
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "slice empty",
			in:           []int{},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int(nil),
		},
		{
			name:         "slice single",
			in:           []int{1},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1},
		},
		{
			name:         "slice multi",
			in:           []int{1, 2, 3},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: -1,
			exp:          []int{1, 2, 3},
		},
		{
			name:         "slice multi break",
			in:           []int{1, 2, 3},
			factoryFunc:  sliceLsFactoryFunc,
			breakAtIndex: 1,
			exp:          []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := tt.factoryFunc()
			ls.AddSlice(tt.in...)
			if tt.breakAtIndex < 0 {
				require.Equal(t, len(tt.exp), ls.Len())
			}

			buf := make([]int, 0)
			out := ls.ToSlice()
			ls.ForEachWithIndex(func(i, a int) bool {
				buf = append(buf, out[i])
				if tt.breakAtIndex == i {
					return true
				}
				i++
				return false
			})

			if out != nil {
				assert.Equal(t, tt.exp, buf)
			}
		})
	}
}
