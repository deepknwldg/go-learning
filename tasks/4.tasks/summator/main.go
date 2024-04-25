package main

import (
	"fmt"
	"sync"
)

func main() {
	channels := []chan int64{
		make(chan int64, 2),
		make(chan int64, 2),
	}

	channels[0] <- 7
	channels[0] <- 3
	channels[1] <- 6
	channels[1] <- 4

	close(channels[0])
	close(channels[1])

	result := sumChannels(channels)

	fmt.Println(result)
}

func sumChannels(inputs []chan int64) int64 {
	var wg sync.WaitGroup
	partialSums := make(chan int64, len(inputs))

	for _, inputChan := range inputs {
		wg.Add(1)
		go func(ch <-chan int64) {
			defer wg.Done()
			var tempSum int64

			for num := range ch {
				tempSum += num
			}

			partialSums <- tempSum
		}(inputChan)
	}

	wg.Wait()
	close(partialSums)

	var totalSum int64
	for partialSum := range partialSums {
		totalSum += partialSum
	}

	return totalSum
}
