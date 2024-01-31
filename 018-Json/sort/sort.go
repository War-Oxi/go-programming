package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var a []int

	for i := 1; i <= 100; i++ {
		randNum := rand.Intn(100)
		a = append(a, randNum)
	}
	fmt.Println(a, "\n")
	sort.IntSlice(a).Sort()
	fmt.Println(a, "\n")

	var b []string = []string{
		"a-kkamji", "d-taeji", "c-ji", "z-jiral",
	}
	fmt.Println(b, "\n")
	sort.Strings(b)
	fmt.Println(b, "\n")
}
