package main

import "fmt"

func addingNumbers(a, b, c int) int {
	return a + b + c
}

func main() {
	varFunc := addingNumbers

	fmt.Println(varFunc(3, 4, 5))
}
