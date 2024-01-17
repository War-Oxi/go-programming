package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d", 'a'+1)

	for i := 'A'; i <= 'Z'; i++ {
		count := 1
		fmt.Println(i)
		for {
			fmt.Printf("\t%#U\n", i)
			count++
			if count > 3 {
				break
			}
		}
	}
}
