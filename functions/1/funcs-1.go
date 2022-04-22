package main

import "fmt"

func main() {

	fooReturned := foo()
	fmt.Println(fooReturned)

	barInt, barStr := bar()
	fmt.Println(barInt, barStr)
}

func foo() int {
	return 0
}

func bar() (int, string) {
	return 1, "bar"
}
