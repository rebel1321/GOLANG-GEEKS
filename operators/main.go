package main

import "fmt"

func main() {

	//Arthimetic operation
	a := 3
	b := 4
	fmt.Printf("Arthimetic operations on a=%d and b=%d:\n\n", a, b)
	var c int
	c = a + b
	fmt.Println("Addition: ", c)

	c = a - b
	fmt.Println("Subtraction: ", c)

	c = a * b	
	fmt.Println("Multiplication: ", c)
	d := float64(a) / float64(b)	
	fmt.Println("Division: ", d)
	c = a % b
	fmt.Println("Modulus: ", c)


	// Relational Operator
	fmt.Printf("\nRelational operations on a=%d and b=%d:\n\n", a, b)
	fmt.Println("Equal: ",a==b)
	fmt.Println("Not Equal: ", a != b)
	fmt.Println("Greater Than: ", a > b)
	fmt.Println("Less Than: ", a < b)
	fmt.Println("Greater Than or Equal: ", a >= b)
	fmt.Println("Less Than or Equal: ", a <= b)

	// Logical Operator on a and b (non-zero values are considered true)
	fmt.Printf("\nLogical operations on a=%d and b=%d (as booleans):\n\n", a, b)
	fmt.Println("AND (a != 0 && b != 0):", a != 0 && b != 0)
	fmt.Println("OR (a != 0 || b != 0):", a != 0 || b != 0)
	fmt.Println("NOT a (! (a != 0)):", !(a != 0))
	fmt.Println("NOT b (! (b != 0)):", !(b != 0))

	// Bitwise Operator
	fmt.Printf("\nBitwise operations on a=%d and b=%d:\n\n", a, b)
	fmt.Printf("Bitwise AND (a & b): %d\n", a & b)
	fmt.Printf("Bitwise OR (a | b): %d\n", a | b)
	fmt.Printf("Bitwise XOR (a ^ b): %d\n", a ^ b)
	fmt.Printf("Bitwise AND NOT (a &^ b): %d\n", a &^ b)
	fmt.Printf("Left Shift (a << 1): %d\n", a << 1)
	fmt.Printf("Right Shift (b >> 1): %d\n", b >> 1)

	// Assignment Operator
	fmt.Printf("\nAssignment operations on a=%d and b=%d:\n\n", a, b)
	a=7
	fmt.Println("a = ", a)
	a += b
	fmt.Println("a += b:", a)
	a -= b
	fmt.Println("a -= b:", a)
	a *= b
	fmt.Println("a *= b:", a)
	a /= b
	fmt.Println("a /= b:", a)
	a %= b
	fmt.Println("a %= b:", a)

	// Reset a for bitwise assignment operations
	a = 3

	a &= b
	fmt.Println("a &= b:", a)
	a |= b
	fmt.Println("a |= b:", a)
	a ^= b
	fmt.Println("a ^= b:", a)
	a &^= b
	fmt.Println("a &^= b:", a)
	a = 3 // reset a
	a <<= 1
	fmt.Println("a <<= 1:", a)
	a >>= 1
	fmt.Println("a >>= 1:", a)


	// Miscellaneous operations
	x:=&a
	fmt.Println("Address ",x)
	y:=*x
	fmt.Println("Value ",y)
}