package main

import (
	"fmt"
)

func main() {
	ca := make(chan int)
	go func() {
		ca <- 42
	}()
	fmt.Println(<-ca)

	ch := make(chan int, 4)
	ch <- 43
	fmt.Println(<-ch)
}
