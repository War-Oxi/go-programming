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
	p1 := person{
		firstName:               "kim",
		lastName:                "taeji",
		favoriteIcecreamFlavors: []string{"chocolate", "strawberry"},
	}

	p2 := person{
		firstName:               "ryu",
		lastName:                "suhyun",
		favoriteIcecreamFlavors: []string{"soda", "mint"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m["suhyun"])

	for k, v := range m {
		fmt.Println("key =>", k)
		fmt.Println("\t", v.firstName)
		fmt.Println("\t", v.lastName)
		for i, val := range v.favoriteIcecreamFlavors {
			fmt.Println("\t\t", i, val)
		}
	}
}
