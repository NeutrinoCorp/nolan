package queue

// Deque a linear collection that supports element insertion and removal at both ends.
// The name deque is short for "double ended queue" and is usually pronounced "deck".
type Deque[T any] interface {
	Queue[T]
	// PeekFirst retrieves, but does not remove, the first element of this deque.
	PeekFirst() T
	// PeekLast retrieves, but does not remove, the last element of this deque.
	PeekLast() T
	// PollFirst retrieves and removes the first element of this deque.
	PollFirst() T
	// PollLast retrieves and removes the last element of this deque.
	PollLast() T
}
