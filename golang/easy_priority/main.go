package main

import (
	"context"
	"strconv"
	"time"
)

func runWorker(
	ctx context.Context,
	highPriorityReqCh <-chan string,
	lowPriorityReqCh <-chan string,
) <-chan struct{} {
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		for {
			select {
			case s := <-highPriorityReqCh:
				mainJob(s)
				continue
			case <-ctx.Done():
				return
			default:
			}

			select {
			case s := <-highPriorityReqCh:
				mainJob(s)
				continue
			case s := <-lowPriorityReqCh:
				mainJob(s)
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	return doneCh
}

func mainJob(s string) {
	println(s)
	time.Sleep(100 * time.Millisecond)
}

func retrieveQueue(
	highPriorityReqCh chan<- string,
	lowPriorityReqCh chan<- string,
) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			lowPriorityReqCh <- "LowPriorityRequest: " + strconv.Itoa(i)
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
	for i := 0; i < 3; i++ {
		go func(i int) {
			highPriorityReqCh <- "HighPriorityRequest: " + strconv.Itoa(i)
		}(i)
	}
}

func main() {
	highPriorityReqCh, lowPriorityReqCh := make(chan string), make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	workerDoneCh := runWorker(ctx, highPriorityReqCh, lowPriorityReqCh)

	go retrieveQueue(highPriorityReqCh, lowPriorityReqCh)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-workerDoneCh
	println("End!")
}
