package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo(2,3,4,5,7,8,9))
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x{
		sum += v
	}
	// fmt.Println(sum)
	return sum
}