package main

import (
	"fmt"
	"sync"
)

func main() {
	arrOfNumbers := []int{2, 4, 6, 8, 10}
	ans := 0
	waiter := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	for _, elem := range arrOfNumbers {
		waiter.Add(1)
		go func(elem int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			ans += elem * elem
			mu.Unlock()
		}(elem, waiter, mutex)
	}
	waiter.Wait()
	fmt.Printf("Total sum is %d\n", ans)
	fmt.Println("Completed")
}
