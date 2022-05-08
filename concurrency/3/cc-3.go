package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	const delta = 100
	wg.Add(delta)
	runtime.GOMAXPROCS(8)

	var counter int32

	for i := 0; i < delta; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
			fmt.Printf("Number of Goroutines\t%d\n", runtime.NumGoroutine())
			fmt.Printf("Current value of counter\t%d\n", atomic.LoadInt32(&counter))
		}()
	}
	wg.Wait()
	fmt.Println("end value", counter)

}
