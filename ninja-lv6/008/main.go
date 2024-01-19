package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Printf("%T\t %v\n", foo(), foo())
	fmt.Printf("%T\t %v\n", f(), f())
}

func foo() func() int {
	// fmt.Println("kkamji")
	return func() int {
		return 42
	}
}
