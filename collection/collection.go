package collection

// Collection a collection represents a group of objects, known as its elements.
// Some collections allow duplicate elements and others do not.
// Some are ordered and others unordered.
type Collection[T any] interface {
	Iterable[T]
	// Add adds an element into this collection.
	Add(v T) bool
	// AddAll adds all the elements into this collection.
	AddAll(src Collection[T]) bool
	// AddSlice adds all the elements in the specified slice (variadic) to this collection.
	AddSlice(items ...T) bool
	// Clear removes all the elements from this collection.
	Clear()
	// Len returns the number of elements in this collection.
	Len() int
	// IsEmpty returns true if this collection contains no elements.
	IsEmpty() bool
	// ToSlice returns all the elements from this collection as a slice of T.
	ToSlice() []T
	// ForEach traverses through all the elements from this collection.
	// Use predicate's return value to indicate a break of the iteration, TRUE meaning a break.
	ForEach(predicateFunc IterablePredicateFunc[T])
}
