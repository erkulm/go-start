package main

import "fmt"

type Person struct {
	First           string   `json:"first"`
	Last            string   `json:"last"`
	FavoriteFlavors []string `json:"favorite_flavors"`
}

func main() {

	person1 := Person{
		First:           "James",
		Last:            "Bond",
		FavoriteFlavors: []string{"vanilla", "chocolate"},
	}

	fmt.Println(person1)

	person2 := Person{
		First:           "Miss",
		Last:            "MoneyPenny",
		FavoriteFlavors: []string{"vanilla", "lemon", "pistachio", "caramel"},
	}
	fmt.Println(person2)

	fmt.Printf("name: %s\t surname: %s\n", person1.First, person1.Last)

	fmt.Println("Favorite Flavors of Person 1:")
	for i, v := range person1.FavoriteFlavors {
		fmt.Printf("%d: %s\t", i+1, v)
	}

	fmt.Printf("\nname: %s\t surname: %s\n", person2.First, person2.Last)

	fmt.Println("Favorite Flavors of Person 2:")
	for i, v := range person2.FavoriteFlavors {
		fmt.Printf("%d: %s\t", i+1, v)
	}
}
