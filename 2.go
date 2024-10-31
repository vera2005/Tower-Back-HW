package main

import (
	"fmt"
	"sync"
)

func main() {
	arrOfNumbers := []int{2, 4, 6, 8, 10}
	waiter := new(sync.WaitGroup)
	for _, elem := range arrOfNumbers {
		waiter.Add(1)
		go func(elem int, wg *sync.WaitGroup) {
			fmt.Printf("%d^2 = %d\n", elem, elem*elem)
			wg.Done()
		}(elem, waiter)
	}
	waiter.Wait()
	fmt.Println("Completed")
}
