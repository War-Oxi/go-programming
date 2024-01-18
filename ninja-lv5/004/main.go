package main

import (
	"fmt"
)

func main() {
	x := struct {
		name      string
		age       int
		friends   map[string]int
		favDrinks []string
	}{
		name: "kimtaeji",
		age:  27,
		friends: map[string]int{
			"park-tae-hyong": 27,
			"kim-jin-su":     28,
		},
		favDrinks: []string{
			"soju",
			"beer",
			"water",
		},
	}
	fmt.Printf("name => %v\nage => %v\nfriends => %v\nfavDrinks => %v\n", x.name, x.age, x.friends, x.favDrinks)

	for key, value := range x.friends {
		fmt.Printf("key => %v\tvalue => %v\n", key, value)
	}
}
