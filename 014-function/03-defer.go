package main

import (
	"fmt"
)

func main() {
	defer foo()
	defer fmt.Println("what the fuck?")
	defer bar()
	bar()
	fmt.Println("kkamji")
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
