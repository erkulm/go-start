package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen()

	receive(c, q)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		case recc, ok := <-c:
			if ok == false {
				return
			}
			fmt.Println("Message received from C:", recc)
		}
	}
}
