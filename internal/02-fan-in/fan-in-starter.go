package main

import (
	"fmt"
	"sync"
)

func main() {
	var producerWg sync.WaitGroup
	fanInChan := make(chan int)

	numWorkers := 10
	numMessages := 10
	for i := 0; i < numWorkers; i++ {
		producerWg.Add(1)
		go func(workerID int) { // worker goroutines
			defer func() {
				producerWg.Done()
				if workerID == 0 {
					close(fanInChan)
				}
			}()

			for i := 0; i < numMessages; i++ {
				fanInChan <- i
			}
		}(i)
	}

	var consumerWg sync.WaitGroup

	consumerWg.Add(1)
	go func() {
		defer consumerWg.Done()
		for msg := range fanInChan {
			fmt.Printf("received %d\n", msg)
		}
	}()

	producerWg.Wait()
	consumerWg.Wait()
}
