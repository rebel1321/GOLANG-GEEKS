package main

import "fmt"
func update(m map[string]int) {
	m["cherry"] = 45
}
func main() {
	m := map[string]int{
		"apple":  200,
		"banana": 100,
		"cherry": 150,
	}
	fmt.Println("Original map:", m)
	fmt.Println("Value for 'banana':", m["banana"])

	m["banana"] = 120 // Mutating the value for "banana"
	fmt.Println("Updated map:", m)

	if val,ok := m["banana"]; ok {
		fmt.Println("Value for 'banana' after mutation:", val)
		m["banana"] = 130 // Further mutating the value for "banana"
		fmt.Println("Final value for 'banana':", m["banana"])
	} else {
		fmt.Println("'banana' not found in the map")
	}

	delete(m, "cherry") // Deleting the key "cherry"
	fmt.Println("Map after deleting 'cherry':", m)
	fmt.Println(m["cherry"]) // Accessing a deleted key returns zero value
	fmt.Println(m)
	update(m)
	fmt.Println("Map after update function call:", m)
}