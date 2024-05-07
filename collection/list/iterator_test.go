package list_test

import (
	"testing"

	"github.com/neutrinocorp/nolan/collection/list"
)

func TestNewIterator(t *testing.T) {
	ls := list.DoublyLinkedList[int]{}
	ls.AddSlice(1, 2, 3, 4, 5)
	t.Log(ls.ToSlice())
}
