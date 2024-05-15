package concurrency

import (
	"context"
	"sync"
)

type Lock interface {
	Lock(ctx context.Context) error
	TryLock(ctx context.Context) (bool, error)
	Unlock(ctx context.Context) error
}

type MutexLock struct {
	Mutex sync.Mutex
}

var _ Lock = (*MutexLock)(nil)

func (d *MutexLock) Lock(_ context.Context) error {
	d.Mutex.Lock()
	return nil
}

func (d *MutexLock) TryLock(_ context.Context) (bool, error) {
	isLocked := d.Mutex.TryLock()
	return isLocked, nil
}

func (d *MutexLock) Unlock(_ context.Context) error {
	d.Mutex.Unlock()
	d.Mutex.TryLock()
	return nil
}
