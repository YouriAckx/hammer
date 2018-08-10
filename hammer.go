package main

import (
	"runtime"
	"time"
)

func main() {
	done := make(chan int)
	for i := 0; i < runtime.NumCPU()+1; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					// loop
				}
			}
		}()
	}
	time.Sleep(time.Hour * 1)
	close(done)
}

