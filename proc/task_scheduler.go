package proc

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/neutrinocorp/nolan/function"
)

var (
	DefaultWorkerExecTimeout    = time.Second * 15
	DefaultMaxConcurrentWorkers = int64(runtime.NumCPU() * 2)
)

type TaskScheduler[T any] struct {
	DelegateFunc         function.DelegateSafeFuncWithContext[T]
	WorkerExecTimeout    time.Duration
	MaxConcurrentWorkers int64

	baseCtx           context.Context
	jobQueue          chan T
	inFlightWaitGroup sync.WaitGroup
	procSemaphore     *semaphore.Weighted
	isShuttingDown    atomic.Bool
}

var _ BootableProcess = (*TaskScheduler[string])(nil)

func NewTaskScheduler[T any](del function.DelegateSafeFuncWithContext[T]) *TaskScheduler[T] {
	return &TaskScheduler[T]{
		DelegateFunc:         del,
		WorkerExecTimeout:    DefaultWorkerExecTimeout,
		MaxConcurrentWorkers: DefaultMaxConcurrentWorkers,
		jobQueue:             make(chan T),
		inFlightWaitGroup:    sync.WaitGroup{},
		procSemaphore:        semaphore.NewWeighted(DefaultMaxConcurrentWorkers),
		isShuttingDown:       atomic.Bool{},
	}
}

func (f *TaskScheduler[T]) Start() error {
	if f.jobQueue == nil {
		f.jobQueue = make(chan T)
	}
	if f.baseCtx == nil {
		f.baseCtx = context.Background()
	}

	for job := range f.jobQueue {
		if err := f.procSemaphore.Acquire(f.baseCtx, 1); err != nil {
			return err
		}
		go f.execute(job)
	}
	return nil
}

func (f *TaskScheduler[T]) Stop(ctx context.Context) error {
	if f.isShuttingDown.Load() {
		return ErrAlreadyTerminated
	}

	f.isShuttingDown.Store(true)
	shutdownChan := make(chan struct{}, 1)
	defer close(shutdownChan)
	go func() {
		// terminating in parallel (inside a goroutine) to avoid deadlocks
		f.inFlightWaitGroup.Wait()
		close(f.jobQueue)
		shutdownChan <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-shutdownChan:
	}
	return nil
}

func (f *TaskScheduler[T]) SubmitWork(args T) error {
	if f.isShuttingDown.Load() {
		return ErrAlreadyTerminated
	}
	f.inFlightWaitGroup.Add(1)
	f.jobQueue <- args
	return nil
}

func (f *TaskScheduler[T]) execute(args T) {
	defer f.inFlightWaitGroup.Done()
	defer f.procSemaphore.Release(1)
	if f.WorkerExecTimeout <= time.Duration(0) {
		_ = f.DelegateFunc(f.baseCtx, args)
		return
	}

	scopedCtx, cancelFunc := context.WithTimeout(f.baseCtx, f.WorkerExecTimeout)
	defer cancelFunc()
	_ = f.DelegateFunc(scopedCtx, args)
}
