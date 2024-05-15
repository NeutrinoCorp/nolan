package executor

import (
	"context"
	"errors"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"

	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/function"
)

var (
	DefaultMaxGoroutines = int64(runtime.NumCPU() * 2)
)

// ConcurrentExecutor is the concurrency-backed implementation of Executor.
// This implementation executes tasks in different goroutines and reduces host's resource exhaustion by
// using an internal semaphore.Weighted.
//
// Zero-value is concurrent-safe and ready to use.
type ConcurrentExecutor[T any] struct {
	MaxGoroutines int64

	singletonLock *sync.Mutex
	procSemaphore *semaphore.Weighted
}

// NewConcurrentExecutor allocates a new ConcurrentExecutor instance.
// The 'maxGoroutines' value is by default DefaultMaxGoroutines.
func NewConcurrentExecutor[T any](maxGoroutines int64) ConcurrentExecutor[T] {
	return ConcurrentExecutor[T]{
		MaxGoroutines: maxGoroutines,
		singletonLock: &sync.Mutex{},
		procSemaphore: semaphore.NewWeighted(maxGoroutines),
	}
}

var _ Executor[string] = (*ConcurrentExecutor[string])(nil)

func (s ConcurrentExecutor[T]) Execute(ctx context.Context, args T, delegate function.DelegateSafeFuncWithContext[T]) error {
	return delegate(ctx, args) // this is just to comply with the Executor interface
}

func (s ConcurrentExecutor[T]) ExecuteAll(ctx context.Context, args T, delegates collection.Collection[function.DelegateSafeFuncWithContext[T]]) error {
	if s.singletonLock == nil {
		s.singletonLock = &sync.Mutex{}
	}
	s.singletonLock.Lock()
	if s.MaxGoroutines <= 0 {
		s.MaxGoroutines = DefaultMaxGoroutines
	}
	if s.procSemaphore == nil {
		s.procSemaphore = semaphore.NewWeighted(s.MaxGoroutines)
	}
	s.singletonLock.Unlock()

	errsMu := sync.Mutex{}
	errs := make([]error, 0, delegates.Len())
	wg := sync.WaitGroup{}
	wg.Add(delegates.Len())
	delegates.ForEach(func(del function.DelegateSafeFuncWithContext[T]) bool {
		if err := s.procSemaphore.Acquire(ctx, 1); err != nil {
			wg.Done()
			errs = append(errs, err)
			return true
		}
		go func(delCopy function.DelegateSafeFuncWithContext[T]) {
			defer s.procSemaphore.Release(1)
			defer wg.Done()
			if errExec := delCopy(ctx, args); errExec != nil {
				errsMu.Lock()
				errs = append(errs, errExec)
				errsMu.Unlock()
			}
		}(del)
		return false
	})

	doneChan := make(chan struct{}, 1)
	defer close(doneChan)
	go func() {
		wg.Wait()
		doneChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		errs = append(errs, ctx.Err())
		return errors.Join(errs...)
	case <-doneChan:
	}

	return errors.Join(errs...)
}
