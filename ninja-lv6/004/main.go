package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "I am", p.age, "Years old.")
}

func main() {
	p1 := person{
		first: "kim",
		last:  "taeji",
		age:   27,
	}
	p1.speak()
}
