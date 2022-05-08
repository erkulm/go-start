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
	var mu sync.Mutex

	var counter int = 0

	for i := 0; i < delta; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			temp := counter
			temp++
			counter = temp
			fmt.Printf("Number of Goroutines\t%d\n", runtime.NumGoroutine())
			fmt.Printf("Current value of counter\t%d\n", counter)
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
