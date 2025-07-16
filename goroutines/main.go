package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Simulate some work
	}
	fmt.Println("I'm the another goroutine...")
}
func main() {
	runtime.GOMAXPROCS(4)
	go printNumbers(6)
	fmt.Println("Started printing numbers")
	time.Sleep(12 * time.Second) // Wait for goroutine to finish
	fmt.Println("I am the main goroutine...")
}