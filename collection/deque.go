package collection

type Deque[T any] interface {
	AddFirst(v T)
	AddLast(v T)
	OfferFirst(v T)
	OfferLast(v T)
	RemoveFirst() bool
	RemoveLast() bool
	PollFirst() T
	PollLast() T
	GetFirst() T
	GetLast() T
}
