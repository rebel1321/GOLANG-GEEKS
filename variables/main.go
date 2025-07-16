package main

import (
	"fmt"
)

func main() {
	// var a int

	// a = 3

// 	var a=5
// 	fmt.Println(a)

// 	var i int //0
// 	var fl float64 //0.0
//  var b bool //false
//  var str string //""
//  var arr [3]int // [0, 0, 0]	
//  var slice []int // []
//  var m map[string]int // map[]
//  var ptr *int //nil

// fmt.Printf("Int: %d\n",i)
// fmt.Printf("Float: %f\n",fl)
// fmt.Printf("Bool: %t\n",b)
// fmt.Printf("String: %s\n",str)
// fmt.Printf("Array: %v\n",arr)
// fmt.Printf("Slice: %v\n",slice)
// fmt.Printf("Map: %v\n",m)
// fmt.Printf("Pointer: %v\n",ptr)


//Shorthand declaration
//  a := 5
//  fmt.Println(a)

//  var b=7
//  fmt.Println(b)

//Type conversion
//  var i int = 42
//  f:= float64(i) // Convert int to float64	
//  fmt.Println("Integer:", i)	
//  fmt.Printf("Float:%f\n", f)

//  var fl float64 = 3.14
// i2 := int(fl) // Convert float64 to int
//  fmt.Println("Float:", fl)
//  fmt.Println("Integer:", i2)

//  var s string = "123"
//  b:=[]byte(s) // Convert string to byte slice
//  fmt.Println("String:", s)
//  fmt.Println("Byte slice:", b)


//  var bt []byte = []byte{72, 101, 108, 108, 111} // Byte slice
//  str := string(bt) // Convert byte slice to string
//  fmt.Println("Byte slice:", bt)
//  fmt.Println("String:", str)


//  s = strconv.Itoa(i)
//  fmt.Println("Converted Integer to String:", s)

//  i, _ = strconv.Atoi(s) // Convert string back to integer
//  fmt.Println("Converted String back to Integer:", i)


const a = 10
const Pi float64 = 3.14
fmt.Printf("a = %v , Type a = %T\n Pi = %v,Type Pi=%T\n",a,a,Pi,Pi)
const (
	L=5
	W=10
	A=L*W
)
fmt.Printf("L = %v, W = %v, Area = %v\n", L, W, A)
const(
	One = iota
	Two
	Three
)
fmt.Printf("One = %v, Two = %v, Three = %v\n", One, Two, Three)

z:=float64(a)
fmt.Printf("z = %.3f, Type z = %T\n", z, z)
}