package main

import "fmt"

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	truck := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	sedan := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(truck)

	fmt.Println(sedan)

	fmt.Println("Truck's number of doors: ", truck.doors)
	fmt.Println("Sedan's number of doors: ", sedan.doors)
}
