package await_termination

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestAwaitTermination(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		i                 int
		want              = 4
		mu                sync.Mutex
		sleepAndIncrement = func() {
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			i++
			mu.Unlock()
		}
	)

	at := NewAwaitTermination(ctx, 3)
	go func() {
		time.Sleep(time.Millisecond * 35)
		cancel()
	}()
	at.Run(func() { sleepAndIncrement() })
	at.Run(func() { sleepAndIncrement() })
	at.Run(func() { sleepAndIncrement() })
	at.Run(func() { sleepAndIncrement() })
	time.Sleep(time.Millisecond * 100)
	at.Run(func() { sleepAndIncrement() })
	at.Run(func() { sleepAndIncrement() })
	at.Wait()

	if i != want {
		t.Errorf("Error. Execution error. got = %d want = %d", i, want)
	}
}
