package main

import (
	"fmt"
)

func main() {
	xi1 := foo()
	fmt.Println("xi1 =>", xi1)
	fmt.Println()

	xi2, xs2 := bar()
	fmt.Printf("xi2 type => %T\t\txi2 value => %v \nxs2 type => %T\txs2 value => %v", xi2, xi2, xs2, xs2)
}

func foo() int {
	var x int = 10
	return x
}

func bar() (int, string) {
	xi, xs := 20, "My name is kkamji"
	// fmt.Println(xi, xs)
	return xi, xs
}
