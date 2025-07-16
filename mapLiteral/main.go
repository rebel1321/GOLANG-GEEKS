package main

func main() {
	m1:=map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	v:= m1["two"]
	println(v) // Output: 2
}