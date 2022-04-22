package main

import "fmt"

func main() {
	defer deferredFunc("so it should be last")
	fmt.Println("This should be first")
}

func deferredFunc(key string) {
	fmt.Println("This function is deferred", key)
}
