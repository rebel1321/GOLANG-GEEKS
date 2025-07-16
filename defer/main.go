package main

import (
	"fmt"
	"os"
)

func demo(i int) {
	fmt.Println("Hello GFG",i)
}
func earlyExit() int {
	defer fmt.Println("Exiting early")
	return 1
}
func main() {
	// defer demo(1)
	// demo(2)
	// defer demo(3)
	// fmt.Println("Welcome to GFG")
	// fmt.Println("This is a defer example")
	// fmt.Println("Have a nice day!")
	// fmt.Println("Enjoy learning Go!")
	// fmt.Println("Goodbye!")



	// fmt.Println("Early exit value:", earlyExit())

	file,err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

}