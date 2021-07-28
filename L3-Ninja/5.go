package main

//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.


import "fmt"

func main() {
	for x:= 10; x <= 100; x++ {      // I am confused on how to do the for loop for this but I searched up 
		                            // for the print statemnt I used println and I used percent which is
									//used to divide the numbers by 4 
		fmt.Println(x,x%4)
	}
}
	// loop from 10 to 100 and print %4 of every iteration

}