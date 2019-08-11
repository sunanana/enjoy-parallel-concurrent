package main

import "sync"

type Semaphore struct {
	concurrency chan struct{}
	wg          sync.WaitGroup
}

func NewSemaphore(concurrency int) *Semaphore {
	return &Semaphore{
		concurrency: make(chan struct{}, concurrency),
		wg:          sync.WaitGroup{},
	}
}

func (s *Semaphore) Run(f func()) {
	s.wg.Add(1)
	s.concurrency <- struct{}{}
	go func() {
		f()
		s.wg.Done()
		<-s.concurrency
	}()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
