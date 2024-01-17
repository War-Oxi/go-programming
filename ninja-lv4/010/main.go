package main

import(
	"fmt"
)

func main() {
	myMap := map[string][]string {
		"bond_james": []string{"dog", "cat", "elephant"},
		"moneypenny_miss": []string{"James Bond", "Literature"},
		"kkamji": []string{"1", "2", "3", "4"},
	}

	myMap["kim_tae_ji"] = []string{"babo", "babo1", "babo2", "babo3"}
	myMap["bond_james"] = append(myMap["bond_james"], "add_bond_james")

	for key, value := range myMap {
		fmt.Printf("key => %v\n\t", key)
		for _, rowValue := range value {
			fmt.Printf("%v ", rowValue)
		}
		fmt.Println()
	}
	fmt.Println("==========================================================")
	delete(myMap, "kkamji")
	
	for key, value := range myMap {
		fmt.Printf("key => %v\n\t", key)
		for _, rowValue := range value {
			fmt.Printf("%v ", rowValue)
		}
		fmt.Println()
	}
}