package queue

import "github.com/neutrinocorp/nolan/collection/list"

// DequeList is the list.List backed implementation of Deque.
type DequeList[T any] struct {
	list.List[T]
}

var _ Deque[string] = (*DequeList[string])(nil)

// NewDequeList allocates a new DequeList instance.
// If no 'src' list is passed, a list.SliceList is allocated by default.
func NewDequeList[T any](src list.List[T]) DequeList[T] {
	if src == nil {
		src = list.NewSliceList[T](nil)
	}
	return DequeList[T]{
		List: src,
	}
}

func (d DequeList[T]) Push(v T) error {
	d.Add(v)
	return nil
}

func (d DequeList[T]) Remove() (T, error) {
	v := d.Poll()
	return v, nil
}

func (d DequeList[T]) Element() (T, error) {
	v := d.Peek()
	return v, nil
}

func (d DequeList[T]) Poll() T {
	return d.PollFirst()
}

func (d DequeList[T]) Peek() T {
	return d.PeekFirst()
}

func (d DequeList[T]) PeekFirst() T {
	return d.GetAt(0)
}

func (d DequeList[T]) PeekLast() T {
	return d.GetAt(d.Len() - 1)
}

func (d DequeList[T]) PollFirst() T {
	val := d.PeekFirst()
	d.RemoveAt(0)
	return val
}

func (d DequeList[T]) PollLast() T {
	val := d.PeekLast()
	d.RemoveAt(d.Len() - 1)
	return val
}
