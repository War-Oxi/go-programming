package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(10))
	fmt.Println(loopfactorial(10))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfactorial(num int) int {
	var sum int = 1
	for ; num > 0; num-- {
		sum *= num
	}

	return sum
}
