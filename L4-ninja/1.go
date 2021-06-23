package main

import "fmt"

func main() {
	x := [5]int{10, 15, 20, 25, 30}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
