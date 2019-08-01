package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sig() <-chan os.Signal {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return sigCh
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sig()
		cancel()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	reqCh := process1(ctx)
	process2(ctx, reqCh)
}

func process1(ctx context.Context) <-chan int {
	ticker := time.NewTicker(10 * time.Millisecond)
	req := make(chan int)
	var i int

	go func() {
		for {
			select {
			case <-ticker.C:
				i++
				req <- i
			case <-ctx.Done():
				ticker.Stop()
				close(req)
				return
			}
		}
	}()
	return req
}

func process2(ctx context.Context, req <-chan int) {
	var (
		concurrency        = make(chan struct{}, 100)
		processEndNotifyCh = make(chan struct{})
		waitCloseCh        = make(chan struct{})
	)
	defer close(concurrency)
	defer close(processEndNotifyCh)

	go func() {
		defer close(waitCloseCh)
		for {
			select {
			case <-ctx.Done():
				goto CLOSING_PHASE
			case <-processEndNotifyCh:
				<-concurrency
			}
		}

	CLOSING_PHASE:
		if len(concurrency) == 0 {
			return
		}
		for range processEndNotifyCh {
			<-concurrency
			if len(concurrency) == 0 {
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			goto WAIT_CLOSE
		case v := <-req:
			select {
			case <-ctx.Done():
			case concurrency <- struct{}{}:
				go func() {
					mainJob(v)
					processEndNotifyCh <- struct{}{}
				}()
			}
		}
	}

WAIT_CLOSE:
	<-waitCloseCh
}

func mainJob(req int) {
	time.Sleep(1 * time.Second)
	fmt.Println(req)
}
