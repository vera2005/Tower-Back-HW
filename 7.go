package main

import (
	"fmt"
	"strconv"
	"sync"
)

func writer(m map[int]string, i int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	m[i] = "Message " + strconv.Itoa(i)
	mu.Unlock()
}

func main() {
	waiter := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	numberOfData := 10
	answer := make(map[int]string)
	for i := 1; i <= numberOfData; i++ {
		waiter.Add(1)
		go writer(answer, i, waiter, mutex)
	}
	waiter.Wait()
	for k, v := range answer {
		fmt.Printf("%d: %s\n", k, v)
	}
	fmt.Println("Complited")

}
