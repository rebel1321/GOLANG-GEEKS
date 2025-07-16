package main

import "fmt"

func operation(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}
func multiply(x, y int) int {
	return x * y
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
func main() {
	// 1.Function as input
	fmt.Println("Addition:", operation(5, 3, add))
	fmt.Println("Multiplication:", operation(5, 3, multiply))

	// /2.Function as output

	mult := multiplier(3)
	fmt.Println("Multiplication by 3:", mult(5))
}