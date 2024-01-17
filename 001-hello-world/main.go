package main

import (
	"fmt"
)

func main() {
	foo("str", "str2", "str2", "str2")
}

func foo(str string, str2 ...string){
	fmt.Println(str)
	fmt.Println(str2)
}