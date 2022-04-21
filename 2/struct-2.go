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

	person2 := Person{
		First:           "Miss",
		Last:            "MoneyPenny",
		FavoriteFlavors: []string{"vanilla", "lemon", "pistachio", "caramel"},
	}

	mapOfPersons := make(map[string]Person, 2)

	mapOfPersons[person1.Last] = person1
	mapOfPersons[person2.Last] = person2

	for i, v := range mapOfPersons {
		tempPerson := v
		fmt.Printf("name: %s\t surname: %s\n", tempPerson.First, tempPerson.Last)

		fmt.Println("Favorite Flavors of Person :", i)
		for i, v := range tempPerson.FavoriteFlavors {
			fmt.Printf("%d: %s\t", i+1, v)
		}
	}
}
