package main

import (
	"fmt"
	"strconv"
)

func main() {
	Hello()
	fmt.Println("Sum of 3 and 4 is:", add(3, 4))
	fmt.Println("Subtraction of 3, 4, and 5 is:", subtract(3, 4, 5))


	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Alice")
	greet("Bob")

	counter := Counter(3)
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

	fmt.Println(Swap(5, 10)) // Swapping 5 and 10

	fmt.Println(Swap1(15, 10)) 
	i,e := Convert("123")	
	if e != nil {
		fmt.Println("Error converting string to int:", e)
		return
	} else {
		fmt.Println("Converted string to int:", i)
	}
}

func Convert(str string) (int,error){
	return strconv.Atoi(str)
}

func Hello() {
	fmt.Println("Hello, World!")
}
func add(a, b int) int {
	return a + b
}

func subtract(num ...int) int {
    if len(num) == 0 {
        return 0
    }
    
    total := num[0] // start from first number
    for i := 1; i < len(num); i++ {
        total -= num[i]
    }
    return total
}

// Counter returns a function that increments and returns a counter starting from the given integer.
// Each time the returned function is called, it increments the counter by 1 and returns the
func Counter(a int) func() int {
	count := a
	return func() int {
		count++
		return count
	}
}

func Swap(a, b int) ( int, int) {
	return b, a
}
func Swap1(a, b int) (x int, y int) {
	x=b
	y=a
	return
}
