package main

import "fmt"

func main() {

	anonymous := struct {
		name      string
		topSecret bool
	}{
		name:      "testName",
		topSecret: true,
	}

	fmt.Println(anonymous)

	m := make(map[string]struct {
		name      string
		topSecret bool
	}, 1)

	m[anonymous.name] = anonymous

	for _, v := range m {
		fmt.Println(v)
	}

}
