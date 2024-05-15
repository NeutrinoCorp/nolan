package collection

import (
	"github.com/neutrinocorp/nolan/function"
)

// ComparatorFunc a functional interface used to indicate if the first argument is greater, less or equals than
// the second.
// If equals, this will return 0, if less -1 and if greater, then it returns 1.
type ComparatorFunc[T any] function.DelegateBiFunc[T, T, int]

// ComparableCollection a kind of Collection for comparable types.
// Note that nolan offers a functional way to operate with Collection and comparable types
// without the need of this interface (e.g., Contains, ContainsAll, ContainsSlice).
type ComparableCollection[T comparable] interface {
	Collection[T]
	// Contains returns true if this collection contains the specified element.
	Contains(v T) bool
	// ContainsAll returns true if this collection contains all the elements in the specified collection.
	ContainsAll(src Collection[T]) bool
	// ContainsSlice returns true if this collection contains all the elements in the specified slice.
	ContainsSlice(src ...T) bool
}

// NOTE: This is used internally to share iterator instances.
// This helps to reduce malloc as we can reuse iterator instances.
func containsWithIterator[T comparable](iter Iterator[T], v T) bool {
	for iter.HasNext() {
		if val := iter.Next(); val == v {
			return true
		}
	}
	return false
}

// Contains returns true if this collection contains the specified element.
func Contains[T comparable](src Collection[T], v T) bool {
	iter := src.NewIterator()
	return containsWithIterator[T](iter, v)
}

// ContainsAll returns true if this collection contains all the elements in the specified collection.
func ContainsAll[T comparable](src Collection[T], cmpColl Collection[T]) bool {
	// Time complexity of O(mn)
	srcIter := src.NewIterator()
	cmpCollIter := cmpColl.NewIterator()
	for cmpCollIter.HasNext() {
		item := cmpCollIter.Next()
		wasFound := containsWithIterator[T](srcIter, item)
		if !wasFound {
			return false
		}
		srcIter.Reset()
	}
	return true
}

// ContainsSlice returns true if this collection contains all the elements in the specified slice.
func ContainsSlice[T comparable](src Collection[T], cmpSlice []T) bool {
	iter := src.NewIterator()
	for _, item := range cmpSlice {
		wasFound := containsWithIterator[T](iter, item)
		if !wasFound {
			return false
		}
		iter.Reset()
	}
	return true
}
