package main

import (
	"fmt"
)

func main() {
	xi1 := 2
	xi2 := 3
	defer sum(xi1, xi2)

	xi1 += 3
	xi2 += 2
}

func sum(xi1 int, xi2 int) int {
	fmt.Printf("xi1 => %v\nxi2 => %v\n", xi1, xi2)
	return xi1 + xi2
}
