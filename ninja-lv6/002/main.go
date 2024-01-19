package main

import (
	"fmt"
)

func main() {
	xi1 := foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(xi1)

	xi2 := bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(xi2)

}

func foo(xIn ...int) int {
	fmt.Printf("%T\t%v\n", xIn, xIn)
	sum := 0
	for _, value := range xIn {
		sum += value
	}
	return sum
}

func bar(sliceIn []int) int {
	sum := 0
	for _, value := range sliceIn {
		sum += value
	}
	return sum
}
