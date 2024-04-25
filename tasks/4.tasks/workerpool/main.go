package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	go func() {
		defer w.wg.Done()
		for task := range w.tasks {
			hash := md5.Sum([]byte(task))
			w.out <- hex.EncodeToString(hash[:])
		}
	}()
}

func main() {
	var tasks chan string = make(chan string, 1)
	var out chan string = make(chan string, 1)
	var wg sync.WaitGroup

	var w *Worker = NewWorker(tasks, &wg, out)
	wg.Add(1)
	w.Run()

	tasks <- "text"
	close(tasks)

	wg.Wait()

	fmt.Println(<-out)

	close(out)
}
