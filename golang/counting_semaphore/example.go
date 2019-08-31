package main

import "time"

func main() {
	cSem := NewCountingSemaphore(3)

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
