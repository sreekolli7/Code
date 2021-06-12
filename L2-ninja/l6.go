package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 2021
		b
		c
		d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

/* Use the itoa command and assign values like a 2021, b 2022,c 2023*/
