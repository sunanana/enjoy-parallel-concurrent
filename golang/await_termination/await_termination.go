package await_termination

import (
	"context"
)

type AwaitTermination struct {
	ctx             context.Context
	concurrency     chan struct{}
	funcEndNotifyCh chan struct{}
	terminatedCh    chan struct{}
}

func NewAwaitTermination(ctx context.Context, concurrency int) *AwaitTermination {
	at := &AwaitTermination{
		ctx:             ctx,
		concurrency:     make(chan struct{}, concurrency),
		funcEndNotifyCh: make(chan struct{}),
		terminatedCh:    make(chan struct{}),
	}

	go func() {
		defer close(at.terminatedCh)
		for {
			select {
			case <-at.ctx.Done():
				if len(at.concurrency) == 0 {
					return
				}
				for range at.funcEndNotifyCh {
					<-at.concurrency
					if len(at.concurrency) == 0 {
						return
					}
				}
			case <-at.funcEndNotifyCh:
				<-at.concurrency
			}
		}
	}()

	return at
}

func (at *AwaitTermination) Run(f func()) {
	select {
	case <-at.ctx.Done():
	case at.concurrency <- struct{}{}:
		select {
		case <-at.ctx.Done():
		default:
			go func() {
				f()
				at.funcEndNotifyCh <- struct{}{}
			}()
		}
	}
}

func (at *AwaitTermination) Wait() {
	<-at.terminatedCh
	at.concurrency = nil
	close(at.funcEndNotifyCh)
}
