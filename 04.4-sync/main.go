package main

import (
	"fmt"
	"sync"
	"time"
)

func process() {
	time.Sleep(time.Millisecond)
	fmt.Println("process")
}

func processWithWg(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("process with wg")
}

func main() {
	waitGroupExample()

}

func waitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		process()
	}()

	wg.Wait()
}
