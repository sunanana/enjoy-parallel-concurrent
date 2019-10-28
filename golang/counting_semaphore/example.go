package counting_semaphore

import (
	"context"
	"time"
)

func example() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cSem := NewCountingSemaphore(ctx, 3)

	cSem.Run(func() { time.Sleep(time.Second * 1); println("1-1") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("1-2") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("1-3") })

	cSem.Run(func() { time.Sleep(time.Second * 1); println("2-1") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("2-2") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("2-3") })

	cSem.Run(func() { time.Sleep(time.Second * 1); println("3-1") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("3-2") })
	cSem.Run(func() { time.Sleep(time.Second * 1); println("3-3") })

	cSem.Wait()
}
