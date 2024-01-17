package main

import (
	"fmt"
)
func main(){
	bd := 1998
	cur := 2024

	for {
		fmt.Println(bd)
		if(bd == cur){
			break
		}
		bd++
	}
}