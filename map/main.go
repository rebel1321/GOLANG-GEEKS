package main

import "fmt"

func change(m map[int]string) {
	m[1] = "two"
}

func main() {
	var m map[int]string
	if m == nil {
		fmt.Println("map is nil")
	}

	m1 := make(map[int]string)
	m1[1] = "one"

	fmt.Println("m1[1]:", m1[1])

	val, exists := m1[2]
	if exists {
		fmt.Println("m1[2]:", val)
	} else {
		fmt.Println("m1[2] does not exist")
	}
	change(m)
	fmt.Println("After change, m1[1]:", m1[1])
}