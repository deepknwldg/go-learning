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
	waitGroupExampleCopy()
	mutexExample()

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

func waitGroupExampleCopy() {
	var wg sync.WaitGroup

	wg.Add(1)
	go processWithWg(&wg) // нужно передавать через указатель

	wg.Wait()
}

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.value += n
}

func mutexExample() {
	var wg sync.WaitGroup

	c := Counter{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go c.Update(10, &wg)
	}

	wg.Wait()
	fmt.Printf("Result is %v\n", c.value)
}

type RWCounter struct {
	m     sync.RWMutex
	value int
}

func (c *RWCounter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.Lock()
	c.value += n
	c.m.Unlock()
}

func (c *RWCounter) GetValue() int {
	c.m.RLock()
	v := c.value
	defer c.m.RUnlock()

	return v
}
