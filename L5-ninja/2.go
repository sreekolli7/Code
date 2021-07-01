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
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"caramel",
			"vanilla ",
			"strawberry",
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

	m := map[string]adult{
		adult1.last: adult1,
		adult2.last: adult2,
	}
	/*
		"Bond": {
			first: "James",
			last:  "Bond",
			favFlavors: []string{
				"chocolate",
				"caramel",
				"vanilla ",
			},
		},
		"Kolli" : {
		first: "Sree",
		last:  "Kolli",
		favFlavors: []string{
			"cherry",
			"vanilla",
			"cookies and cream",
		},

	*/

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		// fmt.Println(i)
		for _, w := range v.favFlavors {
			fmt.Println(w)

		}
	}

}
