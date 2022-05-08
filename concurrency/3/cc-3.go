package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const delta = 100
	wg.Add(delta)
	runtime.GOMAXPROCS(8)

	var counter int = 0

	for i := 0; i < delta; i++ {
		go func() {
			defer wg.Done()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			fmt.Println("Number of Goroutines", runtime.NumGoroutine())
			fmt.Println("Current value of counter", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
