package main

import (
	"fmt"
	"math/rand"
)

func main() {
	queue := generate()
	receiveFromQueue(queue)
}

func generate() <-chan int {
	queue := make(chan int)
	go func() {
		defer close(queue)
		for i := 0; i < 100; i++ {
			queue <- int(rand.Int31n(100))
		}
	}()
	return queue
}

func receiveFromQueue(queue <-chan int) {
	for v := range queue {
		fmt.Println(v)
	}
}
