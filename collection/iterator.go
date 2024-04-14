package collection

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
}

type Iterable[T any] interface {
	NewIterator() Iterator[T]
}
