package main

import (
	"fmt"
	"time"
)

func writer(ch chan<- interface{}) {
	for i := 1; ; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		time.Sleep(600 * time.Millisecond)
	}
}

func reader(ch <-chan interface{}) {
	for msg := range ch {
		fmt.Println(msg)
	}

}

func main() {
	var duration int
	fmt.Println("Enter the duration in seconds: ")
	_, err := fmt.Scan(&duration)
	if err != nil || duration <= 0 {
		fmt.Println("Something wrong... need int > 0")
		return
	}
	msgChan := make(chan interface{})
	go writer(msgChan)
	go reader(msgChan)
	time.Sleep(time.Duration(duration) * time.Second)
	close(msgChan)
	fmt.Println("Completed")
}
