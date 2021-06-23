package main

import "fmt"

func main() {

	x := []int{10, 15, 20, 25, 30, 35, 40, 45, 50, 55}
	x = append(x, 80)
	fmt.Println(x)
}
