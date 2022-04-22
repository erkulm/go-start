package main

import "fmt"

func main() {
	sliceOfIntegers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(sliceOfIntegers...)
	fmt.Println(sum)

	sum2 := bar(sliceOfIntegers)
	fmt.Println(sum2)
}

func foo(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func bar(sliceOfIntegers []int) int {
	sum := 0
	for _, v := range sliceOfIntegers {
		sum += v
	}
	return sum
}
