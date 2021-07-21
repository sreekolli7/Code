package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	fmt.Println(string(bs))
	if err = range of x{
		fmt.println(err)
	}
//would using range of x work in this situation? also the guy used nil where I put range of x why is that? //
}
