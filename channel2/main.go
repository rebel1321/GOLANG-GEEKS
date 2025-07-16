package main

import (
	"fmt"
	"time"
)

func worker(i int, ch chan int) {
	fmt.Println("Sending message on channel....")
	ch <- i

}

func main() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go worker(i, ch)
	}
	time.Sleep(2 * time.Second) // Sleep to ensure all goroutines have sent messages
	for i := 0; i < 3; i++ {
		val := <-ch
		println("Received message:", val)
	}
}
