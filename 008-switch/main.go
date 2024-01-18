package main

import (
	"fmt"
)

func main() {
	// switch "Bond" {
	switch "asd" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	default:
		fmt.Println("this is default")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("this is q")
	}
}
