package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/neutrinocorp/nolan/collection/list"
	"github.com/neutrinocorp/nolan/collection/queue"
)

func TestDeque_AddFirst(t *testing.T) {
	tests := []struct {
		name    string
		in      list.List[int]
		exp     []int
		isStack bool
	}{
		{
			name:    "nil",
			in:      nil,
			exp:     nil,
			isStack: false,
		},
		{
			name:    "linked list empty",
			in:      list.NewDoublyLinkedList[int](),
			exp:     []int{},
			isStack: false,
		},
		{
			name:    "linked list single queue",
			in:      list.NewDoublyLinkedListFromSlice[int]([]int{1}),
			exp:     []int{1},
			isStack: false,
		},
		{
			name:    "linked list single stack",
			in:      list.NewDoublyLinkedListFromSlice[int]([]int{1}),
			exp:     []int{1},
			isStack: true,
		},
		{
			name:    "linked list queue",
			in:      list.NewDoublyLinkedListFromSlice[int]([]int{1, 2, 3, 4, 5}),
			exp:     []int{1, 2, 3, 4, 5},
			isStack: false,
		},
		{
			name:    "linked list stack",
			in:      list.NewDoublyLinkedListFromSlice[int]([]int{1, 2, 3, 4, 5}),
			exp:     []int{5, 4, 3, 2, 1},
			isStack: true,
		},
		{
			name:    "slice empty",
			in:      list.NewSliceList[int](nil),
			exp:     []int{},
			isStack: false,
		},
		{
			name:    "slice single queue",
			in:      list.NewSliceList[int]([]int{1}),
			exp:     []int{1},
			isStack: false,
		},
		{
			name:    "slice single stack",
			in:      list.NewSliceList[int]([]int{1}),
			exp:     []int{1},
			isStack: true,
		},
		{
			name:    "slice queue",
			in:      list.NewSliceList[int]([]int{1, 2, 3, 4, 5}),
			exp:     []int{1, 2, 3, 4, 5},
			isStack: false,
		},
		{
			name:    "slice stack",
			in:      list.NewSliceList[int]([]int{1, 2, 3, 4, 5}),
			exp:     []int{5, 4, 3, 2, 1},
			isStack: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := queue.NewDequeList[int](tt.in)
			require.Equal(t, len(tt.exp), deque.Len())
			for _, item := range tt.exp {
				var polledItem int
				if tt.isStack {
					polledItem = deque.PollLast()
				} else {
					polledItem = deque.PollFirst()
				}
				assert.Equal(t, item, polledItem)
			}

			// validate empty scenarios
			assert.Zero(t, deque.PollFirst())
			assert.Zero(t, deque.PollLast())
		})
	}
}
