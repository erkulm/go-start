package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func main() {
	p := person{age: 34,
		name:    "james",
		surname: "bond",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).age++
}
