package set

import "github.com/neutrinocorp/nolan/collection"

type Set[K comparable] interface {
	collection.ComparableCollection[K]
}
