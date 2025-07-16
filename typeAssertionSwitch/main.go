package main

import "fmt"

func main() {
	var output interface{}
	output = 45
	// i, ok := output.(int)
	// if !ok {
	// 	fmt.Println("Unable to convert output to int")
	// } else {
	// 	fmt.Println("Value:", i)
	// }
	// output = "Hello, World!"
	// s, ok := output.(int)
	// if !ok {
	// 	fmt.Println("Unable to convert output to int")
	// } else {
	// 	fmt.Println("Value:", s)
	// }
	// fmt.Println(s)

	switch v := output.(type) {
	case int:
		fmt.Println("Integer value:", v)
	case string:
		fmt.Println("String value:", v)
	default:
		fmt.Println("Unknown type")
	}


}