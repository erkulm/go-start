package main

import "fmt"

func returnAFunc() func() string {
	return func() string {
		return "returned function"
	}
}

func main() {
	f := returnAFunc()
	fmt.Println(f())
}
