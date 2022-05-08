package main

import (
	"fmt"
	"sync"
	"time"
)

func writeSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Test 1 started")
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Test 1 ended")
}
func writeSomethingElse(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Test 2 started")
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Test 2 ended")
}

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		writeSomething(&wg)
	}()

	go func() {
		writeSomethingElse(&wg)
	}()
	wg.Wait()
	fmt.Println("The end!")

}
