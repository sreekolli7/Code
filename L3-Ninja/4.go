package main

// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.

import (
	"fmt"
)

func main() {
	x := 2004
	for {
		if x > 2021 {
			break
		} //creating  a for loop
		if x%2 == 1 {
			x++
			continue
		}
		fmt.Println(x)
		x++

	}
}

// Same like previous but use for and move condition to if. declare and increment year in for
