package list_test

import (
	"testing"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestLinkedList_Add(t *testing.T) {
	ls := list.DoublyLinkedList[int]{}
	ls.Add(0)
	ls.Add(1)
	ls.Add(2)
	ls.Add(3)
	ls.Add(10)
	t.Log(ls.GetAt(ls.Len() - 1))
}
