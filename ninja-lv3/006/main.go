package main

import (
	"fmt"
)

func main() {
	var str string = "kkamji"

	if str != "kkamji" {
		fmt.Println("str is not kkamji", str)
	} else if str == "kkamji" {
		fmt.Println("str is kkamji")
	} else {
		fmt.Println("str is not matching")
	}
}
