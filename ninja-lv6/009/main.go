package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 9
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	foo(g)
}

func foo(f func(xi []int) int) int {
	n := f([]int{1, 2, 3, 4, 5, 6,})
	n++
	return n
}
