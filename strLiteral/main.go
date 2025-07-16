package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr, 6,8,9) // Appending an element to the slice
	fmt.Println(len(arr)) // Length of the slice
	fmt.Println(cap(arr)) // Capacity of the slice

var arr1 []int
for i := 0; i < 10; i++ {
	arr1 = append(arr1, i)
	fmt.Printf("Length: %v, Capacity: %v, Value: %v\n", len(arr1), cap(arr1), arr1)
}
}