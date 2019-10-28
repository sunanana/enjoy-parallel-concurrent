package counting_semaphore

import (
	"context"
	"sync"
)

type CountingSemaphore struct {
	ctx         context.Context
	concurrency chan struct{}
	wg          sync.WaitGroup
}

func NewCountingSemaphore(ctx context.Context, concurrency int) *CountingSemaphore {
	return &CountingSemaphore{
		ctx:         ctx,
		concurrency: make(chan struct{}, concurrency),
		wg:          sync.WaitGroup{},
	}
}

func (s *CountingSemaphore) Run(f func()) {
	select {
	case <-s.ctx.Done():
	case s.concurrency <- struct{}{}:
		select {
		case <-s.ctx.Done():
		default:
			s.wg.Add(1)
			go func() {
				f()
				s.wg.Done()
				<-s.concurrency
			}()
		}
	}
}

func (s *CountingSemaphore) Wait() {
	s.wg.Wait()
}
