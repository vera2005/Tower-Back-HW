package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func reader(id int, ch <-chan string, w *sync.WaitGroup) {
	defer w.Done()
	for elem := range ch {
		fmt.Printf("Worker %d get: %s", id, elem)
	}
}

func main() {
	waiter := new(sync.WaitGroup)

	var numberOfWorkers int
	fmt.Println("Enter the number of workeres: ")
	fmt.Scan(&numberOfWorkers)
	mainChan := make(chan string)

	// Обработка сигналов для завершения программы
	sigs := make(chan os.Signal, 1)     // канал для получения системных сигналов
	signal.Notify(sigs, syscall.SIGINT) //signal.Notify - связка канала с сигналами для обработки
	// syscall.SIGIN - прерыавние (Ctrl+C)

	// запись даннных в главыный канал
	go func() {
		i := 1
		for {
			mainChan <- fmt.Sprintf("Message %d\n", i)
			i += 1
			time.Sleep(400 * time.Millisecond)
		}
	}()

	//запуск воркеров
	for i := 0; i < numberOfWorkers; i++ {
		waiter.Add(1)
		go reader(i+1, mainChan, waiter)
	}

	// ожидание сигнала для закрытия канала
	go func() {
		<-sigs
		close(mainChan)
	}()

	waiter.Wait()
	fmt.Println("Stopped..")

}
