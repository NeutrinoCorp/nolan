package function

type BiFunction[A any, B any] func(a A, b B)
type BiPredicate[A, B, R any] func(a A, b B) R
