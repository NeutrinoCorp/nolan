package list_test

import (
	"testing"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestLinkedList_Add(t *testing.T) {
	var ls list.List[int] = &list.DoublyLinkedList[int]{}
	ls.Add(-1)
	ls.Add(1)
	ls.Add(2)
	ls.Add(3)
	ls.AddAt(3, 7)
	ls.Add(8)
	ls.Add(9)

	iter := ls.NewIterator()
	for iter.HasNext() {
		t.Logf("at fwd iter: %v", iter.Next())
	}

	for iter.HasPrevious() {
		t.Logf("at bwd iter: %v", iter.Previous())
	}

	t.Log(ls.GetAt(ls.Len() - 1))
	t.Log(ls.RemoveAt(ls.Len() - 1))
	t.Log(ls.GetAt(ls.Len() - 1))

	anotherLs := &list.DoublyLinkedList[int]{}
	anotherLs.Add(4)
	anotherLs.Add(5)
	anotherLs.Add(6)
	ls.AddAllAt(3, anotherLs)

	iter = ls.NewIterator()
	for iter.HasNext() {
		t.Logf("add buffer: %v", iter.Next())
	}

	sublist := ls.ToSubList(0, 3)
	iter = sublist.NewIterator()
	for iter.HasNext() {
		t.Logf("sublist buffer: %v", iter.Next())
	}
}

func TestSliceList_Add(t *testing.T) {
	var ls list.List[int] = &list.SliceList[int]{
		Source: make([]int, 0, 10),
	}
	ls.Add(-1)
	ls.Add(1)
	ls.Add(2)
	ls.Add(3)
	ls.AddAt(3, 7)
	ls.Add(8)
	ls.Add(9)

	iter := ls.NewIterator()
	for iter.HasNext() {
		t.Logf("at fwd iter: %v", iter.Next())
	}

	for iter.HasPrevious() {
		t.Logf("at bwd iter: %v", iter.Previous())
	}

	t.Log(ls.GetAt(ls.Len() - 1))
	t.Log(ls.RemoveAt(ls.Len() - 1))
	t.Log(ls.GetAt(ls.Len() - 1))

	anotherLs := &list.DoublyLinkedList[int]{}
	anotherLs.Add(4)
	anotherLs.Add(5)
	anotherLs.Add(6)
	ls.AddAllAt(3, anotherLs)

	iter = ls.NewIterator()
	for iter.HasNext() {
		t.Logf("add buffer: %v", iter.Next())
	}
	t.Log(ls.Len())

	sublist := ls.ToSubList(0, 3)
	iter = sublist.NewIterator()
	for iter.HasNext() {
		t.Logf("sublist buffer: %v", iter.Next())
	}
}
