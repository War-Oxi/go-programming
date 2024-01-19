package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	wg.Wait()

	func() {
		for i := 0; i < 200; i++ {
			fmt.Println(i)
		}
	}()
}
