package main

import (
	"fmt"
)

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

func main() {
	x := person{
		firstName:               "kim",
		lastName:                "taeji",
		favoriteIcecreamFlavors: []string{"chocolate", "strawberry"},
	}

	fmt.Println(x)
	fmt.Printf("firstName = %v \nlastName = %v \nfavoriteIcecreamFlavors = %v \n", x.firstName, x.lastName, x.favoriteIcecreamFlavors[1])
}
