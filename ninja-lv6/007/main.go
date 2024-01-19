package main

import (
	"fmt"
)

func main() {
	x := sum
	fmt.Println(x(1, 2))
	fmt.Printf("x type => %T", x)
}

func sum(xi1 int, xi2 int) int {
	return xi1 + xi2
}
