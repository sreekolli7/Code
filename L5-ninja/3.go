package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	// s := sedan{
	// 	vehicle: vehicle{
	// 		doors: 4,
	// 		color: "black",
	// 	},
	// 	luxury: false,
	// }

	m := map[string]truck{
		"abc": t,
	}

	for _, b := range m {
		fmt.Println(b.fourWheel)

	}

}
