package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "and I am", p.age, "years old.")
}

func main() {

	james := person{
		firstName: "James",
		lastName:  "Bond",
		age:       34,
	}

	james.speak()

}
