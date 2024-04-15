package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

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

var _ List[string] = &DoublyLinkedList[string]{}

func NewDoublyLinkedListFromCollection[T any](src collection.Collection[T]) *DoublyLinkedList[T] {
	ls := &DoublyLinkedList[T]{}
	iter := src.NewIterator()
	for iter.HasNext() {
		ls.Add(iter.Next())
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

func (l *DoublyLinkedList[T]) addToNode(prevNode, nextNode *doublyLinkedListNode[T]) {
	if prevNode == nil {
		return
	}

	l.len++
	nextNode.next = prevNode.next
	prevNode.next = nextNode
	nextNode.previous = prevNode
	if nextNode.next != nil {
		nextNode.next.previous = nextNode
	}
	l.tail = nextNode
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

func (l *DoublyLinkedList[T]) SetAt(index int, v T) T {
	node := l.getNodeAt(index)
	tmpKey := node.key
	node.key = v
	return tmpKey
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) T {
	node := l.getNodeAt(index)
	if node == nil {
		var zeroVal T
		return zeroVal
	}

	key := node.key
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

func (l *DoublyLinkedList[T]) GetAt(index int) T {
	node := l.getNodeAt(index)
	if node == nil {
		var zeroVal T
		return zeroVal
	}
	return node.key
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}

func (l *DoublyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

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

func (l *DoublyLinkedList[T]) AddAt(index int, v T) {
	node := l.getNodeAt(index)
	if node == nil {
		return
	}

	l.addToNewNode(node, v)
}

func (l *DoublyLinkedList[T]) NewIterator() collection.Iterator[T] {
	return NewIterator[T](l)
}

func (l *DoublyLinkedList[T]) AddAll(src collection.Collection[T]) bool {
	iter := src.NewIterator()
	wasMod := iter.HasNext()
	for iter.HasNext() {
		l.Add(iter.Next())
	}
	return wasMod
}

func (l *DoublyLinkedList[T]) AddAllAt(index int, src collection.Collection[T]) bool {
	if index > l.len-1 || src.Len() == 0 || index < 0 {
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

func (l *DoublyLinkedList[T]) ToSubList(fromIndex, toIndex int) List[T] {
	// TODO: IMPLEMENT ME!
	return nil
	// newList := &DoublyLinkedList[T]{}
	// currentNode := l.getNodeAt(fromIndex)
	// count := 0
	// for currentNode != nil {
	// 	if fromIndex == toIndex {
	// 		break
	// 	}
	//
	// 	newList.addToNode(count, currentNode)
	// 	fromIndex++
	// 	count++
	// 	currentNode = currentNode.next
	// }
	// return newList
}
