package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter input:")
	var a int
	fmt.Scanln(&a)

	fmt.Printf("%d\t%x\t%b\n", a, a, a)
}
