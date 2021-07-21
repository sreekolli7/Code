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

	bs, err := toJSON(p1)

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
}
// my own cutom create//
if err = range for x {
	return []byte, fmt.Errorf("Error in the code", err)
}
fmt.Println(x)
}