package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Sree Kolli",
	}
	fmt.Println(p1)

	fmt.Println(p1)
}
