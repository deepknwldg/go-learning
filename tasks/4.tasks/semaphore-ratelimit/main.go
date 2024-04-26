package main

import (
	"fmt"
	"time"
)

type semaphore chan struct{}

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	sem := NewSemaphore(10) // ограничиваем количество одновременных запросов до 10

	for i := 0; i < 1000; i++ {
		sem.Acquire(1) // захватываем семафор

		go func(n int) {
			defer sem.Release(1) // освобождаем семафор после выполнения запроса
			sendRequest(n)
		}(i)
	}

	// Ждем, пока все горутины завершат свою работу
	sem.Acquire(10)
}

func sendRequest(n int) {
	start := time.Now()
	time.Sleep(100 * time.Millisecond) // имитация отправки запроса
	elapsed := time.Since(start)
	fmt.Printf("Request %d took %s\n", n, elapsed)
}
