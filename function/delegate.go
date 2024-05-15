package function

import "context"

// DelegateFunc is a type of function that passes responsibility of an action.
type DelegateFunc[A, R any] func(a A) R

// DelegateSafeFunc is a type of function that passes responsibility of an action, returning either a response
// or an error.
type DelegateSafeFunc[A, R any] func(a A) (R, error)

// DelegateBiFunc is a type of function that passes responsibility of an action using two arguments.
type DelegateBiFunc[A, B, R any] func(a A, b B) R

// DelegateBiFuncSafe is a type of function that passes responsibility of an action using two arguments,
// returning either a response or an error.
type DelegateBiFuncSafe[A, B, R any] func(a A, b B) (R, error)

// DelegateFuncWithContext is a type of DelegateFunc that passes responsibility of an action.
type DelegateFuncWithContext[A, R any] func(ctx context.Context, a A) R

// DelegateSafeFuncWithContext is a type of function that passes responsibility of an action, returning either a
// response or an error.
type DelegateSafeFuncWithContext[A any] DelegateBiFunc[context.Context, A, error]
