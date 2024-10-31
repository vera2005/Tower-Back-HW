package main

import (
	"fmt"
	"sync"
)

func writer1(ch chan<- int, data []int, w *sync.WaitGroup) {
	defer w.Done()
	for _, elem := range data {
		ch <- elem
	}
	close(ch)
}

func writer2(ch1 <-chan int, ch2 chan<- int, w *sync.WaitGroup) {
	defer w.Done()
	for elem := range ch1 {
		ch2 <- elem * 2
	}
	close(ch2)
}

func reader(ch <-chan int, w *sync.WaitGroup) {
	defer w.Done()
	for elem := range ch {
		fmt.Println(elem)
	}
}
func main() {
	waiter := new(sync.WaitGroup)
	waiter.Add(3)
	dataArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chan1 := make(chan int)
	chan2 := make(chan int)
	go writer1(chan1, dataArray, waiter)
	go writer2(chan1, chan2, waiter)
	go reader(chan2, waiter)
	waiter.Wait()
	fmt.Println("Completed")

}
