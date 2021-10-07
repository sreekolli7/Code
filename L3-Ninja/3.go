package main

import "fmt"

func main() {

	numbers := []int{11, 9, 1, 45, 411}
	lowestnumber := numbers[0]
	for _, i := range numbers {
		if i < lowestnumber {
			lowestnumber = i

		}
	}
	fmt.Println(lowestnumber)

}
