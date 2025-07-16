package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		ch1 <- "Hello from channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from channel 2"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- "Hello from channel 3"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch4 <- "Hello from channel 4"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch5 <- "Hello from channel 5"
	}()
		for i := 0; i < 2; i++ {
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case msg3 := <-ch3:
		fmt.Println("Received:", msg3)
	case msg4 := <-ch4:
		fmt.Println("Received:", msg4)
	case msg5 := <-ch5:
		fmt.Println("Received:", msg5)

	case <-time.After(1 * time.Second):
		fmt.Println("No messages received within the timeout period")
		default:
		fmt.Println("No messages received, continuing to wait...")
	}

}
}