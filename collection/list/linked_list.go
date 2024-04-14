package list

import "github.com/neutrinocorp/nolan/collection"

type doublyLinkedListNode[T any] struct {
	previous *doublyLinkedListNode[T]
	next     *doublyLinkedListNode[T]
	key      T
}

type DoublyLinkedList[T any] struct {
	head *doublyLinkedListNode[T]
	tail *doublyLinkedListNode[T]
	len  int
}

var _ collection.Collection[string] = &DoublyLinkedList[string]{}

func (l *DoublyLinkedList[T]) getNodeAt(i int) *doublyLinkedListNode[T] {
	currentNode := l.head
	count := 0
	for currentNode != nil {
		if count == i {
			return currentNode
		}
		count++
		currentNode = currentNode.next
	}
	return nil
}

func (l *DoublyLinkedList[T]) addToNode(prevNode *doublyLinkedListNode[T], v T) {
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
}

func (l *DoublyLinkedList[T]) Add(v T) {
	if l.head == nil {
		l.head = &doublyLinkedListNode[T]{
			key: v,
		}
		l.len++
		return
	}
	l.addToNode(l.getNodeAt(l.len-1), v)
}

func (l *DoublyLinkedList[T]) GetAt(pos int) T {
	node := l.getNodeAt(pos)
	if node == nil {
		var zeroVal T
		return zeroVal
	}
	return node.key
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}
