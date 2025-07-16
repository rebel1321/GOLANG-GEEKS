package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	fmt.Println("Sending on the channel: ")
	time.Sleep(2 * time.Second) // Simulate some work
	ch <- 42
}

func receive(ch <-chan int) {
	fmt.Println("Waiting for the message...")
	msg := <-ch
	fmt.Println("Received:", msg) 
}
func main() {
	ch := make(chan int)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- 20
	}()
	receive(ch)
}