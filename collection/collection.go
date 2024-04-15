package collection

// Collection A collection represents a group of objects, known as its elements. Some collections allow duplicate
// elements and others do not. Some are ordered and others unordered.
type Collection[T any] interface {
	Iterable[T]
	// Add Ensures that this collection contains the specified element.
	Add(v T) bool
	// AddAll Adds all the elements in the specified collection to this collection.
	AddAll(src Collection[T]) bool
	// Clear Removes all the elements from this collection.
	Clear()
	// Len Returns the number of elements in this collection.
	Len() int
	// IsEmpty Returns true if this collection contains no elements.
	IsEmpty() bool
	// ToSlice Returns all the elements from this collection as a slice of T.
	ToSlice() []T
}

type ComparableCollection[T comparable] interface {
	// Contains Returns true if this collection contains the specified element.
	Contains(v T) bool
	// ContainsAll Returns true if this collection contains all the elements in the specified collection.
	ContainsAll(src Collection[T]) bool
}
