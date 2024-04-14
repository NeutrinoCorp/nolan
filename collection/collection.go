package collection

type Collection[T any] interface {
	Add(v T)
	Len() int
}
