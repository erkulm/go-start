package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
