package executor

import (
	"context"

	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/function"
)

// SyncExecutor is the synchronous implementation of Executor.
// It executes tasks in sequence and no extra goroutines are used.
type SyncExecutor[T any] struct {
}

var _ Executor[string] = SyncExecutor[string]{}

func (s SyncExecutor[T]) Execute(ctx context.Context, args T, delegate function.DelegateSafeFuncWithContext[T]) error {
	return delegate(ctx, args)
}

func (s SyncExecutor[T]) ExecuteAll(ctx context.Context, args T, delegates collection.Collection[function.DelegateSafeFuncWithContext[T]]) (err error) {
	delegates.ForEach(func(del function.DelegateSafeFuncWithContext[T]) bool {
		if errProc := del(ctx, args); errProc != nil {
			err = errProc
			return true
		}
		return false
	})
	return
}
