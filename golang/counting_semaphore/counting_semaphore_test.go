package counting_semaphore

import (
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

	cs := NewCountingSemaphore(concurrency)
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
