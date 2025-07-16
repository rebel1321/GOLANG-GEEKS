package main

import "fmt"

func incr(a *[3]int) {
	(*a)[0]=100
}
func incrs(a []int) {
	a[0]=100
}
func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3 

	fmt.Println("Array a:", a)
	b := [3]int{4, 5, 6}
	fmt.Println("Array b:", b)


	var s []int
	s = append(s, 7, 8, 9)
	fmt.Println("Slice s:", s)

	s1 := make([]int, 5, 10)
	s1=append(s1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)
	fmt.Println("Slice s1:", s1)

	incr(&a)
	fmt.Println("Array a after incr:", a)
	incrs(s)
	fmt.Println("Slice s after incrs:", s)

	fmt.Println("Length of a:", len(a))
	fmt.Println("Length of b:", len(b))
	fmt.Println("Length of s:", len(s), "Capacity of s:", cap(s))
	fmt.Println("Length of s1:", len(s1), "Capacity of s1:", cap(s1))
}