package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	cond          *sync.Cond
	workers       int
	isSystemAlive bool
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		cond:          sync.NewCond(&sync.Mutex{}),
		workers:       workers,
		isSystemAlive: true,
	}
}

func (wp *WorkerPool) Worker(id int) {
	for {
		wp.cond.L.Lock()
		for !wp.isSystemAlive {
			fmt.Printf("Worker %d is sleeping\n", id)
			wp.cond.Wait()
		}
		fmt.Printf("Worker %d is working\n", id)
		wp.cond.L.Unlock()
		// Здесь должна быть реализация получения и выполнения задачи
		time.Sleep(1 * time.Second) // Имитация работы
	}
}

func (wp *WorkerPool) CheckSystemLiveness() {
	for {
		time.Sleep(3 * time.Second) // Имитация проверки
		alive := checkSystemLiveness()
		wp.cond.L.Lock()
		wp.isSystemAlive = alive
		if alive {
			fmt.Println("System is alive, notifying all workers")
			wp.cond.Broadcast() // Уведомление всех горутин
		} else {
			fmt.Println("System is down")
		}
		wp.cond.L.Unlock()
	}
}

func checkSystemLiveness() bool {
	// Здесь должна быть реализация проверки liveness внешней системы
	return true // Имитация работы
}

func main() {
	wp := NewWorkerPool(5)

	for i := 0; i < wp.workers; i++ {
		go wp.Worker(i)
	}

	go wp.CheckSystemLiveness()

	select {} // Бесконечный цикл, чтобы горутины не завершились
}
