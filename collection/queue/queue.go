package queue

import "github.com/neutrinocorp/nolan/collection"

// Queue A collection designed for holding elements prior to processing.
// Besides basic collection.Collection operations, queues provide additional insertion, extraction, and
// inspection operations.
type Queue[T any] interface {
	collection.Collection[T]
	// Push Inserts the specified element into this queue if it is possible to do so.
	Push(v T) error
	// Remove Retrieves and removes the head of this queue.
	Remove() (T, error)
	// Element Retrieves, but does not remove, the head of this queue.
	Element() (T, error)
	// Poll Retrieves and removes the head of this queue, or returns null (or zero-value) if this queue is empty.
	Poll() T
	// Peek Retrieves, but does not remove, the head of this queue, or returns null (or zero-value) if
	// this queue is empty.
	Peek() T
}

// BatchQueue is a special kind of Queue that exposes batch operations.
type BatchQueue[T any] interface {
	Queue[T]
	// RemoveMany Retrieves and removes, the head of this queue til N is reached.
	RemoveMany(n int) ([]T, error)
	// GetMany Retrieves, but does not remove, the head of this queue til N is reached.
	GetMany(n int) ([]T, error)
	// PollMany Retrieves and removes the head of this queue till N is reached, or returns null if this
	// queue is empty.
	PollMany(n int) []T
	// PeekMany Retrieves, but does not remove, the head of this queue, or returns null (or zero-value) if
	// this queue is empty.
	PeekMany(n int) []T
}
