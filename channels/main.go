package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		
		ch <- 1
		fmt.Println("Sent message 1 on channel...")
		ch <- 2
		fmt.Println("Sent message 2 on channel...")
		ch <- 3
		fmt.Println("Sent message 3 on channel...")
	}()
	// Receive the message from the channel
	fmt.Println("Waiting for sender...")
	time.Sleep(5* time.Second) // Sleep to ensure the sender has sent messages
	i1 := <-ch
	fmt.Println("Received message:", i1)
	i2 := <-ch
	fmt.Println("Received message:", i2)
	i3 := <-ch
	fmt.Println("Received message:", i3)
}