package main

import (
	"fmt"
)

func main(){
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooo, James."}
	xxs := [][]string{xs1, xs2}
	fmt.Println("xxs =>", xxs)
	fmt.Println("xs1 =>", xs1)
	fmt.Println("xs2 =>", xs2)

	for i, xs :=range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

	xs1 = append(xs1, "kkamji", "kkamtaeji")
	xxs = append(xxs, xs1)
	fmt.Println(xxs)
	// xxs = append(xxs[0], "kkamji")
	// fmt.Println(xxs[0])
}