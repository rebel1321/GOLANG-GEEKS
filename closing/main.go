package main

import "fmt"

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go sender(ch)
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
	fmt.Println("Channel is closed")
}