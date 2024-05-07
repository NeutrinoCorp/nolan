package collection

import "github.com/neutrinocorp/nolan/function"

type (
	// IterablePredicateFunc A functional interface used by function-oriented iterators to traverse data.
	//
	// T: Item's type.
	//
	// Return: A boolean. Tells the iterable concrete implementation routine to break the traversing process.
	IterablePredicateFunc[T any] function.Predicate[T, bool]
	// IterablePredicateBiFunc A functional interface used by function-oriented iterators to traverse data.
	// Used for collections with two items like Maps.
	//
	// A: First item's type.
	//
	// B: Second item's type.
	//
	// Return: A boolean. Tells the iterable concrete implementation routine to break the traversing process.
	IterablePredicateBiFunc[A, B any] function.BiPredicate[A, B, bool]
)

type ForwardIterator[T any] interface {
	HasNext() bool
	Next() T
}

type ReverseIterator[T any] interface {
	HasPrevious() bool
	Previous() T
}

type Iterator[T any] interface {
	ForwardIterator[T]
	ReverseIterator[T]
	Reset()
}

type Iterable[T any] interface {
	NewIterator() Iterator[T]
}
