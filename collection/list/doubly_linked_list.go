package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

// doublyLinkedListNode is a structure used to hold a key and both, previous and next, pointer reference to neighbor
// nodes.
type doublyLinkedListNode[T any] struct {
	previous *doublyLinkedListNode[T]
	next     *doublyLinkedListNode[T]
	key      T
}

// DoublyLinkedList is the doubly linked list implementation of List. The user of this interface has precise control
// over where in the list each element is inserted. The user can access elements by their integer
// index (position in the list), and search for elements in the list.
type DoublyLinkedList[T any] struct {
	head *doublyLinkedListNode[T]
	tail *doublyLinkedListNode[T]
	len  int
}

var _ List[string] = &DoublyLinkedList[string]{}

// NewDoublyLinkedList allocates a new DoublyLinkedList instance.
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// NewDoublyLinkedListFromCollection allocates a new DoublyLinkedList instance copying 'src' items into it.
func NewDoublyLinkedListFromCollection[T any](src collection.Collection[T]) *DoublyLinkedList[T] {
	ls := &DoublyLinkedList[T]{}
	iter := src.NewIterator()
	for iter.HasNext() {
		ls.Add(iter.Next())
	}
	return ls
}

// NewDoublyLinkedListFromSlice allocates a new DoublyLinkedList instance copying 'src' items into it.
func NewDoublyLinkedListFromSlice[T any](src []T) *DoublyLinkedList[T] {
	ls := &DoublyLinkedList[T]{}
	for _, v := range src {
		ls.Add(v)
	}
	return ls
}

func (l *DoublyLinkedList[T]) getNodeAt(index int) *doublyLinkedListNode[T] {
	if l.len > 1 && index == l.len-1 {
		return l.tail
	}

	currentNode := l.head
	count := 0
	for currentNode != nil {
		if count == index {
			return currentNode
		}
		count++
		currentNode = currentNode.next
	}
	return nil
}

func (l *DoublyLinkedList[T]) addToNewNode(prevNode *doublyLinkedListNode[T], v T) {
	if prevNode == nil {
		return
	}

	l.len++
	newNode := &doublyLinkedListNode[T]{
		key: v,
	}
	newNode.next = prevNode.next
	prevNode.next = newNode
	newNode.previous = prevNode
	if newNode.next != nil {
		newNode.next.previous = newNode
	}
	l.tail = newNode
}

// NewIterator returns an iterator over elements of type T.
func (l *DoublyLinkedList[T]) NewIterator() collection.Iterator[T] {
	return NewIterator[T](l)
}

// Add adds an element into this collection.
func (l *DoublyLinkedList[T]) Add(v T) bool {
	if l.head == nil {
		l.head = &doublyLinkedListNode[T]{
			key: v,
		}
		l.tail = l.head
		l.len++
		return true
	}
	l.addToNewNode(l.getNodeAt(l.len-1), v)
	return true
}

// AddAll adds all the elements into this collection.
func (l *DoublyLinkedList[T]) AddAll(src collection.Collection[T]) bool {
	wasMod := false
	src.ForEach(func(a T) bool {
		wasAdded := l.Add(a)
		if wasAdded && !wasMod {
			wasMod = true
		}
		return false
	})
	return wasMod
}

// AddSlice adds all the elements in the specified slice (variadic) to this collection.
func (l *DoublyLinkedList[T]) AddSlice(items ...T) bool {
	if len(items) == 0 {
		return false
	}

	newList := NewDoublyLinkedListFromSlice(items)
	node := l.getNodeAt(l.Len() - 1)
	newList.getNodeAt(0).previous = node
	if l.Len() == 0 {
		l.head = newList.getNodeAt(0)
	} else {
		node.next = newList.getNodeAt(0)
	}
	l.tail = newList.getNodeAt(newList.len - 1)
	l.len += newList.len
	return true
}

// Clear removes all the elements from this collection.
func (l *DoublyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

// Len returns the number of elements in this collection.
func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}

// IsEmpty returns true if this collection contains no elements.
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

// ToSlice returns all the elements from this collection as a slice of T.
func (l *DoublyLinkedList[T]) ToSlice() []T {
	if l.head == nil {
		return nil
	}

	currentNode := l.head
	buf := make([]T, 0, l.len)
	for currentNode != nil {
		buf = append(buf, currentNode.key)
		currentNode = currentNode.next
	}
	return buf
}

// ForEach traverses through all the elements from this collection.
// Use predicate's return value to indicate a break of the iteration, TRUE meaning a break.
func (l *DoublyLinkedList[T]) ForEach(predicateFunc collection.IterablePredicateFunc[T]) {
	currentNode := l.head
	for currentNode != nil {
		breakIter := predicateFunc(currentNode.key)
		if breakIter {
			break
		}
		currentNode = currentNode.next
	}
}

// AddAt inserts the specified element at the specified position in this list.
func (l *DoublyLinkedList[T]) AddAt(index int, v T) {
	if !isValidIndex(index, l.len) {
		return
	}
	node := l.getNodeAt(index)
	if node == nil {
		return
	}

	l.addToNewNode(node, v)
}

// AddAllAt inserts all the elements in the specified collection into this list at the specified position.
func (l *DoublyLinkedList[T]) AddAllAt(index int, src collection.Collection[T]) bool {
	if !isValidIndex(index, l.len) {
		return false
	}

	newList := NewDoublyLinkedListFromCollection(src)
	node := l.getNodeAt(index)
	newList.getNodeAt(0).previous = node
	newList.getNodeAt(newList.len - 1).next = node.next
	node.next = newList.getNodeAt(0)
	if index == l.len-1 {
		l.tail = newList.getNodeAt(newList.len - 1)
	}
	l.len += newList.len
	return true
}

// SetAt replaces the element at the specified position in this list with the specified element.
func (l *DoublyLinkedList[T]) SetAt(index int, v T) T {
	if !isValidIndex(index, l.len) {
		var zeroVal T
		return zeroVal
	}
	node := l.getNodeAt(index)
	tmpKey := node.key
	node.key = v
	return tmpKey
}

// GetAt returns the element at the specified position in this list.
func (l *DoublyLinkedList[T]) GetAt(index int) T {
	if !isValidIndex(index, l.len) {
		var zeroVal T
		return zeroVal
	}
	node := l.getNodeAt(index)
	if node == nil {
		var zeroVal T
		return zeroVal
	}
	return node.key
}

// RemoveAt removes the element at the specified position in this list.
func (l *DoublyLinkedList[T]) RemoveAt(index int) T {
	if !isValidIndex(index, l.len) {
		var zeroVal T
		return zeroVal
	}

	node := l.getNodeAt(index)
	if node == nil {
		var zeroVal T
		return zeroVal
	}

	key := node.key
	if index == 0 {
		l.head = node.next
	}
	if index == l.len-1 {
		l.tail = l.tail.previous
	}
	if node.previous != nil {
		node.previous.next = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	}
	l.len--
	return key
}

// ForEachWithIndex traverses through all the elements from this collection.
// Use predicate's return value to indicate a break of the iteration.
// 'A' is the index while 'B' is the item.
func (l *DoublyLinkedList[T]) ForEachWithIndex(predicateFunc collection.IterablePredicateBiFunc[int, T]) {
	currentNode := l.head
	count := 0
	for currentNode != nil {
		breakIter := predicateFunc(count, currentNode.key)
		if breakIter {
			break
		}
		currentNode = currentNode.next
		count++
	}
}

// ToSubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex,
// exclusive.
func (l *DoublyLinkedList[T]) ToSubList(fromIndex, toIndex int) List[T] {
	if fromIndex < 0 || toIndex >= l.len || toIndex < fromIndex {
		return nil
	}

	ls := &DoublyLinkedList[T]{}
	currentNode := l.head
	for currentNode != nil && fromIndex <= toIndex {
		ls.Add(currentNode.key)
		currentNode = currentNode.next
		fromIndex++
	}
	return ls
}
