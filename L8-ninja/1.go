package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Sree",
		Last:  "Kolli",
		Age:   17,
	}

	p2 := person{
		First: "Kolli",
		Last:  "Sree",
		Age:   18,
	}

	person := []person{p1, p2}

	fmt.Println(person)

}
