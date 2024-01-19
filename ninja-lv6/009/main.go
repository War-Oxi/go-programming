package main

import (
	"fmt"
)

func main() {

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
	n := f()
	n++
	return n
}
