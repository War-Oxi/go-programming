package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.age = 27
}

func main() {
	p1 := person{
		first: "kim",
		last:  "tae",
		age:   20,
	}

	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}
