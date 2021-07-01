package main

import "fmt"

func main() {

	x := []int{10, 20, 25, 30, 35, 40, 45, 50, 55, 60}
	fmt.Println(x[:])
	fmt.Println(x[:])
	fmt.Println(x[:])
	fmt.Println(x[:])
	fmt.Println(x)
}
// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
//