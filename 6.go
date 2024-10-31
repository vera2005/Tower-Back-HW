package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// met1 использует канал для остановки горутины
func met1(ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Gorutine stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

// met2 использует контекст для остановки горутины
func met2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gorutine stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

// met3 использует WaitGroup для отслеживания завершения горутины
func met3(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

// met4 использует сигнал для остановки горутины
func met4(ch <-chan os.Signal) {
	for {
		select {
		case <-ch:
			fmt.Println("Gorutine stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var method int
	fmt.Println("Enter a number to select a method for stopping the goroutine:")
	fmt.Println("1 - Channel usage")
	fmt.Println("2 - Use of context")
	fmt.Println("3 - Use of WaitGroup")
	fmt.Println("4 - Use of signals")
	fmt.Scan(&method)

	switch method {
	case 1:
		signalChan := make(chan struct{})
		go met1(signalChan)
		time.Sleep(3 * time.Second)
		close(signalChan)
		time.Sleep(1 * time.Second)

	case 2:
		ctx, cancel := context.WithCancel(context.Background())
		go met2(ctx)
		time.Sleep(5 * time.Second)
		cancel()
		time.Sleep(1 * time.Second)

	case 3:
		waiter := new(sync.WaitGroup)
		waiter.Add(1)
		go met3(waiter)
		waiter.Wait()
		fmt.Println("Gorutine stopped")

	case 4:
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Добавлено SIGTERM для большей гибкости
		go met4(sigs)

		// Блокируем главный поток до получения сигнала
		select {}
	default:
		fmt.Println("Unexpected method")
	}
}
