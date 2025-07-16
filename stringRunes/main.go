package main

import "fmt"

func main() {
	s := "Hello,GFG!🤗"
	s="hello🤣"

	r := []rune(s)

	fmt.Println("String:", s)
	fmt.Println("Runes:", r)

	for i, c := range s {
		fmt.Printf("Character: %c, and index: %d with ASCII value: %v\n", c, i, c)
	}


	for i, c := range r {
		fmt.Printf("Rune Character: %c, and index: %d with ASCII/Unicode value: %U\n", c, i, c)
	}

	s = s + " World! 🌍"
	fmt.Println("Updated String:", s)
}