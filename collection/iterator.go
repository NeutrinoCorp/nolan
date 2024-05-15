package collection

import "github.com/neutrinocorp/nolan/function"

type (
	// IterablePredicateFunc a functional interface used by function-oriented iterators to traverse data.
	//
	// T: Item's type.
	//
	// Return: A boolean. Tells the iterable concrete implementation routine to break the traversing process.
	IterablePredicateFunc[T any] function.PredicateFunc[T]
	// IterablePredicateBiFunc a functional interface used by function-oriented iterators to traverse data.
	// Used for collections with two items like Maps.
	//
	// A: First item's type.
	//
	// B: Second item's type.
	//
	// Return: A boolean. Tells the iterable concrete implementation routine to break the traversing process.
	IterablePredicateBiFunc[A, B any] function.PredicateBiFunc[A, B]
)

// ForwardIterator an iterator over a collection traversing in ascending order.
type ForwardIterator[T any] interface {
	// HasNext indicates if the iterator has another item to retrieve.
	HasNext() bool
	// Next retrieves the next item.
	Next() T
}

// ReverseIterator an iterator over a collection traversing in descending order.
type ReverseIterator[T any] interface {
	// HasPrevious indicates if the iterator has another item to retrieve.
	HasPrevious() bool
	// Previous retrieves the previous item.
	Previous() T
}

// Iterator an iterator over a collection.
// It is stateful, meaning the user of this interface needs to control iteration routine calls.
type Iterator[T any] interface {
	ForwardIterator[T]
	ReverseIterator[T]
	// Reset restarts the state of the Iterator to default values.
	Reset()
}

// Iterable implementing this interface allows an object to be the target of the "for-each loop"-like statement.
type Iterable[T any] interface {
	// NewIterator returns an iterator over elements of type T.
	NewIterator() Iterator[T]
}
