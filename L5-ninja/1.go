package main

import (
	"fmt"
)

type adult struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	adult1 := adult{
		first: "Bill",
		last:  "Gates",
		favFlavors: []string{
			"chocolate",
			"caramel",
			"vanilla ",
		},
	}

	adult2 := adult{
		first: "Sree",
		last:  "Kolli",
		favFlavors: []string{
			"cherry",
			"vanilla",
			"cookies and cream",
		},
	}
	for _, v := range adult1.favFlavors {
		fmt.Println(v)
	}
	fmt.Println(adult1.first)
	fmt.Println(adult1.last)

	for _, i := range adult2.favFlavors {
		fmt.Println(i)
	}
	fmt.Println(adult2.first)
	fmt.Println(adult2.last)

}
