package proc_test

import (
	"context"
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/neutrinocorp/nolan/function"
	"github.com/neutrinocorp/nolan/proc"
)

func TestWorkStealingProcess_Start(t *testing.T) {
	var delFunc function.DelegateSafeFuncWithContext[string] = func(ctx context.Context, args string) error {
		time.Sleep(time.Millisecond * 500)
		log.Printf("at delegate with args %v", args)
		return nil
	}
	initRoutines := runtime.NumGoroutine()
	wsProc := proc.NewTaskScheduler(delFunc)
	go func() {
		_ = wsProc.Start()
	}()
	assert.Greater(t, runtime.NumGoroutine(), initRoutines)
	err := wsProc.SubmitWork("lemur")
	assert.GreaterOrEqual(t, runtime.NumGoroutine(), initRoutines+1)
	assert.NoError(t, err)
	err = wsProc.SubmitWork("lemur")
	assert.GreaterOrEqual(t, runtime.NumGoroutine(), initRoutines+2)
	assert.NoError(t, err)
	scopedCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	err = wsProc.Stop(scopedCtx)
	time.Sleep(time.Millisecond * 100)
	assert.LessOrEqual(t, runtime.NumGoroutine(), initRoutines)
	assert.NoError(t, err)
	err = wsProc.SubmitWork("lemur")
	assert.NotNil(t, err)
}
