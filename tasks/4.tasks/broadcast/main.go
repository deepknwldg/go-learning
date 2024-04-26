package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	workers []*Worker
	mu      sync.Mutex
	cond    *sync.Cond
}

type Worker struct {
	id int
}

func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		workers: make([]*Worker, size),
	}
	wp.cond = sync.NewCond(&wp.mu)
	for i := 0; i < size; i++ {
		wp.workers[i] = &Worker{id: i}
	}
	return wp
}

func (w *Worker) DoWork(wp *WorkerPool) {
	wp.cond.L.Lock()
	wp.cond.Wait()
	fmt.Printf("Worker %d received broadcast\n", w.id)
	wp.cond.L.Unlock()
}

func (wp *WorkerPool) Broadcast() {
	wp.cond.L.Lock()
	wp.cond.Broadcast()
	wp.cond.L.Unlock()
}

func main() {
	wp := NewWorkerPool(5)
	for _, w := range wp.workers {
		go w.DoWork(wp)
	}

	time.Sleep(time.Second)
	wp.Broadcast()

	time.Sleep(time.Second)
}
