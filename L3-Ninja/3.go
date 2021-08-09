package main

// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.

import (
	"fmt"
)

func main() {
	x := 2004 //assigning x to a varaible which is 2004 when I was born

	for x <= 2021 { // making a for statment and making it greater or equal to this year                         //used for incrementing
		fmt.Println(x)
		x++
	}
}

// declare variable and assign your year of birth to it and then use for less than present as condition for for to print years and increment it

//Code cleanup is needed
