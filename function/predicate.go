package function

// PredicateFunc is a functional interface that accepts a sole argument of type A and returns a value of type bool.
type PredicateFunc[A any] DelegateFunc[A, bool]

// PredicateFuncSafe is a functional interface that accepts a sole argument of type A and returns a value of type bool.
// This routine differs from the PredicateFunc interface as this also returns an error if something fails.
type PredicateFuncSafe[A any] DelegateSafeFunc[A, bool]

// PredicateBiFunc is a functional interface that accepts two arguments of type A and B, returning a value of type bool.
type PredicateBiFunc[A, B any] DelegateBiFunc[A, B, bool]

// PredicateBiFuncSafe is a functional interface that accepts two arguments of type A and B, returning
// a value of type bool.
// This routine differs from the PredicateFunc interface as this also returns an error if something fails.
type PredicateBiFuncSafe[A, B any] DelegateBiFuncSafe[A, B, bool]
