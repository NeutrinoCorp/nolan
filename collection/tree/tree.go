package tree

import (
	"github.com/neutrinocorp/nolan/collection"
)

type Tree[T collection.Ordered] interface {
}

type elver struct {
	ComparatorFunc collection.ComparatorFunc[string]
}

func (e elver) name() {
	cmpResult := e.ComparatorFunc("a", "b")
	if cmpResult == 0 {

	}
}
