package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, task <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range task {
		fmt.Printf("Task %d is picked up by the worker %d\n", task, id)
		time.Sleep(1*time.Second) // Simulate work
		result <- task * task // Simulate processing result
	}
}

func main() {

		const workers=3
		const jobs=6
		tasks:=make(chan int ,jobs)
		result:=make(chan int ,jobs)

		var wg sync.WaitGroup
		for i := 1; i < workers; i++ {
			wg.Add(1)
			go worker(i,tasks, result, &wg)
		}

		for i := 1; i <= jobs; i++ {
			tasks <- i
		}
		close(tasks)

		wg.Wait()
		close(result)
		for res := range result {
			fmt.Printf("Result: %d\n", res)
		}
		
	}