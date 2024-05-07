package bag

import (
	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/set"
)

// Bag Defines a collection that counts the number of times an object appears in
// the collection.
type Bag[T any] interface {
	collection.Collection[T]
	AddCopies(v T, numberCopies int) bool
	GetCount(v T) int
}

// NewUniqueSet Returns a set.Set of unique elements in the Bag.
func NewUniqueSet[T comparable](src Bag[T]) set.Set[T] {
	buf := set.HashSet[T]{}
	iter := src.NewIterator()
	for iter.HasNext() {
		item := iter.Next()
		buf.Add(item)
	}
	return buf
}
