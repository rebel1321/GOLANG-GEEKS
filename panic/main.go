package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
func main() {
	
	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0)) //Panicked here
	fmt.Println(divide(4, 1))
}