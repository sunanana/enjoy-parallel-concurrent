package counting_semaphore

import "sync"

type CountingSemaphore struct {
	concurrency chan struct{}
	wg          sync.WaitGroup
}

func NewCountingSemaphore(concurrency int) *CountingSemaphore {
	return &CountingSemaphore{
		concurrency: make(chan struct{}, concurrency),
		wg:          sync.WaitGroup{},
	}
}

func (s *CountingSemaphore) Run(f func()) {
	s.wg.Add(1)
	s.concurrency <- struct{}{}
	go func() {
		f()
		s.wg.Done()
		<-s.concurrency
	}()
}

func (s *CountingSemaphore) Wait() {
	s.wg.Wait()
}
