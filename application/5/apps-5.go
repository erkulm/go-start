package main

import (
	"fmt"
	"sort"
)

type users []user

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u users) Len() int {
	return len(u)
}
func (u users) Less(i, j int) bool {
	if u[i].Age < u[j].Age {
		return true
	} else if u[i].Age == u[j].Age {
		return u[i].Last < u[j].Last
	}
	return false
}
func (u users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u users) Print() {
	for i, v := range u {
		sort.Strings(v.Sayings)
		fmt.Printf("User #%d\n", i+1)
		fmt.Println("First", v.First)
		fmt.Println("Last", v.Last)
		fmt.Println("Age", v.Age)
		fmt.Println("Sayings")
		for j, saying := range v.Sayings {
			fmt.Printf("\t%d - %s\n", j+1, saying)
		}
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := users{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(users)

	fmt.Println(users)

	users.Print()
	// your code goes here

}
