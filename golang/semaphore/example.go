package main

import "time"

func main() {
	sem := NewSemaphore(3)

	sem.Run(func() { time.Sleep(time.Second * 1); println("1-1") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("1-2") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("1-3") })

	sem.Run(func() { time.Sleep(time.Second * 1); println("2-1") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("2-2") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("2-3") })

	sem.Run(func() { time.Sleep(time.Second * 1); println("3-1") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("3-2") })
	sem.Run(func() { time.Sleep(time.Second * 1); println("3-3") })

	sem.Wait()
}
