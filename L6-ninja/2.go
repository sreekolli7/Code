package main

import (
	"fmt"
)

func main() {
	i := []int{10, 20, 25, 30, 35, 40, 45, 50}
	n := foo(i)
	fmt.Println(n)

	x := []int{10, 20, 25, 30, 35, 40, 45, 50}
	n2 := bar(x)
	fmt.Println(n2)
}

func foo(x int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
