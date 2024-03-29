package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   32,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	fmt.Println(string(bs))
	fmt.Println(err)

}
