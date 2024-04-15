package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

type Iterator[T any] struct {
	source               List[T]
	currentForwardIndex  int
	currentBackwardIndex int
}

var _ collection.Iterator[string] = &Iterator[string]{}

func NewIterator[T any](src List[T]) *Iterator[T] {
	return &Iterator[T]{
		source:               src,
		currentForwardIndex:  0,
		currentBackwardIndex: src.Len() - 1,
	}
}

func (i *Iterator[T]) HasNext() bool {
	return i.currentForwardIndex <= i.source.Len()-1
}

func (i *Iterator[T]) Next() T {
	key := i.source.GetAt(i.currentForwardIndex)
	i.currentForwardIndex++
	return key
}

func (i *Iterator[T]) HasPrevious() bool {
	return i.currentBackwardIndex >= 0
}

func (i *Iterator[T]) Previous() T {
	key := i.source.GetAt(i.currentBackwardIndex)
	i.currentBackwardIndex--
	return key
}
