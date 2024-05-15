package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestNewIterator(t *testing.T) {
	tests := []struct {
		name string
		ls   list.List[int]
	}{
		{
			name: "slice list empty",
			ls:   list.NewSliceList[int](nil),
		},
		{
			name: "slice list single",
			ls:   list.NewSliceList[int]([]int{1}),
		},
		{
			name: "slice list multi",
			ls:   list.NewSliceList[int]([]int{1, 2, 3}),
		},
		{
			name: "linked list empty",
			ls:   list.NewDoublyLinkedList[int](),
		},
		{
			name: "linked list single",
			ls:   list.NewDoublyLinkedListFromSlice[int]([]int{1}),
		},
		{
			name: "linked list multi",
			ls:   list.NewDoublyLinkedListFromSlice[int]([]int{1, 2, 3}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := tt.ls.NewIterator()
			i := 0
			for iter.HasNext() {
				item := iter.Next()
				assert.Equal(t, tt.ls.GetAt(i), item)
				i++
			}
			assert.Equal(t, tt.ls.Len(), i)
			i = tt.ls.Len() - 1
			for iter.HasPrevious() {
				item := iter.Previous()
				assert.Equal(t, tt.ls.GetAt(i), item)
				i--
			}
			assert.Equal(t, -1, i)
		})
	}
}
