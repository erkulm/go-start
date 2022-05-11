package main

import (
	"fmt"
	"math/rand"
)

func main() {
	queue, done := generate()
	receiveFromQueue(queue, done)
}

func generate() (q, d <-chan int) {
	queue := make(chan int)
	done := make(chan int)

	for i := 0; i < 10; i++ {
		done <- 1
		go func() {
			fmt.Println("Goroutine launched! number:", i)
			for j := 0; j < 10; j++ {
				queue <- int(rand.Int31n(100))
			}
			current := <-done
			current++
			done <- current
		}()
	}
	return queue, done
}

func receiveFromQueue(queue, done <-chan int) {
	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case d := <-done:
			if d == 9 {
				return
			}
		}
	}
}
