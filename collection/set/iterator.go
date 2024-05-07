package set

import (
	"github.com/neutrinocorp/nolan/collection"
)

type Iterator[K comparable] struct {
	keySet               []K
	currentForwardIndex  int
	currentBackwardIndex int
}

var _ collection.Iterator[string] = &Iterator[string]{}

func NewIterator[K comparable](src Set[K]) *Iterator[K] {
	return &Iterator[K]{
		keySet:               src.ToSlice(),
		currentForwardIndex:  0,
		currentBackwardIndex: src.Len() - 1,
	}
}

func (i *Iterator[K]) HasNext() bool {
	return i.currentForwardIndex <= len(i.keySet)-1
}

func (i *Iterator[K]) Next() K {
	key := i.keySet[i.currentForwardIndex]
	i.currentForwardIndex++
	return key
}

func (i *Iterator[K]) HasPrevious() bool {
	return i.currentBackwardIndex >= 0
}

func (i *Iterator[K]) Previous() K {
	key := i.keySet[i.currentBackwardIndex]
	i.currentBackwardIndex--
	return key
}

func (i *Iterator[K]) Reset() {
	i.currentForwardIndex = 0
	i.currentBackwardIndex = len(i.keySet) - 1
}
