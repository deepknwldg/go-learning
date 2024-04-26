package main

import (
	"fmt"
	"sync"
)

func main() {
	raceCondition()
	// withoutRaceCondition()
}

// go run -race racedecetor.go
func raceCondition() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
func withoutRaceCondition() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
