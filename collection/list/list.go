package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

// List an ordered collection.Collection (also known as a sequence).
// The user of this interface has precise control over where in the list each element is inserted.
// The user can access elements by their integer index (position in the list), and search for elements in the list.
//
// Unlike sets, lists typically allow duplicate elements. More formally, lists typically allow pairs of elements e1
// and e2 such that e1 == e2, and they typically allow multiple nil elements if they allow null elements at all.
type List[T any] interface {
	collection.Collection[T]
	// AddAt inserts the specified element at the specified position in this list.
	AddAt(index int, v T)
	// AddAllAt inserts all the elements in the specified collection into this list at the specified position.
	AddAllAt(index int, src collection.Collection[T]) bool
	// SetAt replaces the element at the specified position in this list with the specified element.
	SetAt(index int, v T) T
	// GetAt returns the element at the specified position in this list.
	GetAt(index int) T
	// RemoveAt removes the element at the specified position in this list.
	RemoveAt(index int) T
	// ForEachWithIndex traverses through all the elements from this collection.
	// Use predicate's return value to indicate a break of the iteration.
	// 'A' is the index while 'B' is the item.
	ForEachWithIndex(predicateFunc collection.IterablePredicateBiFunc[int, T])
	// ToSubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex,
	// exclusive.
	ToSubList(fromIndex, toIndex int) List[T]
}
