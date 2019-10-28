package counting_semaphore

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestCountingSemaphore(t *testing.T) {
	var (
		i                 int
		concurrency       = 3
		interimResultWant = 3
		want              = 5
		mu                sync.Mutex
		sleepAndIncrement = func() {
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			i++
			mu.Unlock()
		}
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cs := NewCountingSemaphore(ctx, concurrency)
	cs.Run(func() { sleepAndIncrement() })
	cs.Run(func() { sleepAndIncrement() })
	cs.Run(func() { sleepAndIncrement() })
	time.Sleep(150 * time.Millisecond)

	mu.Lock()
	interimResultGot := i
	mu.Unlock()
	if interimResultGot != interimResultWant {
		t.Errorf("Error. Interim result error. got = %d want = %d", interimResultGot, interimResultWant)
		return
	}

	cs.Run(func() { sleepAndIncrement() })
	cs.Run(func() { sleepAndIncrement() })
	cs.Wait()

	if i != want {
		t.Errorf("Error. Execution error. got = %d want = %d", i, want)
	}
}

func TestCountingSemaphore_InterruptCancel(t *testing.T) {
	var (
		i                 int
		concurrency       = 2
		want              = 2
		mu                sync.Mutex
		sleepAndIncrement = func() {
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			i++
			mu.Unlock()
		}
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cs := NewCountingSemaphore(ctx, concurrency)
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	cs.Run(func() { sleepAndIncrement() }) // execute
	cs.Run(func() { sleepAndIncrement() }) // execute
	cs.Run(func() { sleepAndIncrement() }) // cancel while waiting
	cs.Wait()
	cs.Run(func() { sleepAndIncrement() }) // cancel before waiting

	if i != want {
		t.Errorf("Error. Execution error. got = %d want = %d", i, want)
	}
}
