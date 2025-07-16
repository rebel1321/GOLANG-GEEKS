package main

import "fmt"

func adder() func(int) int {
	x := 5
	return func(i int) int {
		return x + i
	}
}
func fun() (func(),func()){
	x:=0
	inc:= func() {
		x++
		fmt.Println("Incremented value:", x)
	}
	dec := func() {	
		x--
		fmt.Println("Decremented value:", x)
	}
	return inc, dec
}
func main() {
	// Closure example
	a := adder()
	fmt.Println(a(10)) // Output: 15
	fmt.Println(a(20)) // Output: 25
	

	in ,de :=fun()
	in() // Incremented value: 1
	de() // Decremented value: 0
	in() // Incremented value: 1
	in()
	in()
	in()
	in()
	in()
	in()
	de()
	de()
	de()
	de()
}