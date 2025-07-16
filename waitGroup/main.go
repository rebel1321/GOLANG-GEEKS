package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i, delay int, wg *sync.WaitGroup){
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker No: %v Processing started\n", i)
	time.Sleep(time.Duration(delay) * time.Second) // Simulate some work
	fmt.Printf("Worker No: %v Processing finished\n", i)
}
func main() {
	var wg sync.WaitGroup
	// for i:=0; i < 5; i++ {
	// 	wg.Add(1) // Increment the WaitGroup counter
	// 	go worker(i+1, &wg)
	// }
	wg.Add(5) // Set the number of goroutines to wait for

	go worker(1, 2, &wg)
	go worker(2, 4, &wg)
	go worker(3, 1, &wg)
	go worker(4, 3, &wg)
	go worker(5, 5, &wg)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed their tasks")
}