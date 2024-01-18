package main

import (
	"fmt"
)

func main() {
	bd := 1998
	cur := 2024
	for bd <= cur {
		fmt.Println(bd)
		bd++
	}
}
