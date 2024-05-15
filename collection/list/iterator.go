package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

// Iterator is the implementation of collection.Iterator using an underlying List.
type Iterator[T any] struct {
	source               List[T]
	currentForwardIndex  int
	currentBackwardIndex int
}

var _ collection.Iterator[string] = &Iterator[string]{}

// NewIterator allocates a new Iterator instance.
func NewIterator[T any](src List[T]) *Iterator[T] {
	return &Iterator[T]{
		source:               src,
		currentForwardIndex:  0,
		currentBackwardIndex: src.Len() - 1,
	}
}

// HasNext indicates if the iterator has another item to retrieve.
func (i *Iterator[T]) HasNext() bool {
	return i.currentForwardIndex <= i.source.Len()-1
}

// Next retrieves the next item.
func (i *Iterator[T]) Next() T {
	key := i.source.GetAt(i.currentForwardIndex)
	i.currentForwardIndex++
	return key
}

// HasPrevious indicates if the iterator has another item to retrieve.
func (i *Iterator[T]) HasPrevious() bool {
	return i.currentBackwardIndex >= 0
}

// Previous retrieves the previous item.
func (i *Iterator[T]) Previous() T {
	key := i.source.GetAt(i.currentBackwardIndex)
	i.currentBackwardIndex--
	return key
}

// Reset restarts the state of the Iterator to default values.
func (i *Iterator[T]) Reset() {
	i.currentForwardIndex = 0
	i.currentBackwardIndex = i.source.Len() - 1
}
