package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", rectangle.Area())
	rectangle.Scale(10)
	fmt.Println("Scaled rectangle area:", rectangle.Area())
}