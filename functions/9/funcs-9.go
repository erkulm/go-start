package main

import "fmt"

func main() {
	s := functionAsAnArgument("test", operateOnString)
	fmt.Println(s)
}

func functionAsAnArgument(s string, f func(s string) string) string {
	return f(s)
}

func operateOnString(s string) string {
	return s + " ***"
}
