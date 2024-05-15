package executor_test

import (
	"context"
	"testing"
	"time"

	"github.com/neutrinocorp/nolan/collection/list"
	"github.com/neutrinocorp/nolan/concurrency/executor"
	"github.com/neutrinocorp/nolan/function"
)

func TestConcurrentExecutor_ExecuteAll(t *testing.T) {
	execPlan := list.NewSliceList[function.DelegateSafeFuncWithContext[string]](nil)
	execPlan.Add(func(a context.Context, b string) error {
		t.Logf("at delegate 0 with args: %s", b)
		return nil
	})
	execPlan.Add(func(a context.Context, b string) error {
		time.Sleep(time.Second * 2)
		t.Logf("at delegate 1 with args: %s", b)
		return nil
	})
	execPlan.Add(func(a context.Context, b string) error {
		t.Logf("at delegate 2 with args: %s", b)
		return nil
	})
	var exec executor.Executor[string] = executor.ConcurrentExecutor[string]{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	err := exec.ExecuteAll(ctx, "foo", execPlan)
	t.Log(err)
}
