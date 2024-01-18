package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWhell bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	myTruck := truck{
		vehicle: vehicle{
			doors: "2door",
			color: "white",
		},
		fourWhell: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: "4door",
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(myTruck)
	fmt.Println(mySedan)
}
