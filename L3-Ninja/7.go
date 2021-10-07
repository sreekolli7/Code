package main

//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

import "fmt"

func main() {

	//x := []int{62, 72, 82, 92, 102}
	//for _, grades := range x {
	//	switch {
	//	case grades >= 90:
	//		fmt.Println("Student 1")
	//	case grades >= 80:
	//		fmt.Println("Student 2")
	//	case grades >= 70:
	//		fmt.Println("Student 3")
	//	case grades >= 60:
	//		fmt.Println("Student 4")
	//	case grades >= 50:
	//		fmt.Println("Student 5")
	//	}
	//}
	count := 0
	for i := 3; i <= 1000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				count++
			}

		}
		if count == 0 {
			fmt.Println(i)
		} else {
			count = 0
			smallnumber := i
			for _, smallnumber := range i {
				if smallnumber > i {

				}
				fmt.Println(smallnumber)
			}

			// continue before program with else if and else
		}
	}
}
