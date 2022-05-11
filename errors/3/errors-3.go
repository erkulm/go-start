package main

import "fmt"

type customErr struct {
	error string
}

func (ce customErr) Error() string {
	return ce.error
}

func foo(err error) {
	fmt.Println(fmt.Errorf("error: %v", err))
}

func main() {
	ce := customErr{
		error: "custom error occured",
	}
	fmt.Printf("Type of custom error:%T\n", ce)
	foo(ce)
}
