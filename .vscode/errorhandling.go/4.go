package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if x < 0 {
		x := fmt.Errorf("Coffee Needed", x)
		return 0, sqrtError{"50.2289 N", "99.4656 W", x}
		// write your error code here
	}

	return 42, nils
}
