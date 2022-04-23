package main

import "fmt"

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) calculateArea() float64 {
	return s.length * s.length
}

func (c circle) calculateArea() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	calculateArea() float64
}

func info(s shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	sq := square{
		length: 5,
	}

	ci := circle{
		radius: 5,
	}

	info(sq)
	info(ci)
	fmt.Println("Area of the square", sq.calculateArea(), "Area of the circle", ci.calculateArea())
}
