package main

import "fmt"

func incr(n *int) int{
	return *n+1
}
func main() {
	var p *int
	x := 42

	p = &x                         // p now points to x
	fmt.Println("Address of x:", p)          // Printing the address of x
	fmt.Println("Value of x:", *p) // Dereferencing p to get the value of x

	*p = incr(p)                  // Incrementing the value at the address p points to
	fmt.Println("New value of x:", *p) // Printing the new value of x after

	pp:=&p
	fmt.Println("Address of p:", pp) // Printing the address of p
	fmt.Println("Value of p:", *pp) // Dereferencing pp to get the value of
	fmt.Println("Value at address p points to:", **pp) // Dereferencing pp to get the value at the address p points to
}