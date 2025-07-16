package main

import "fmt"

type Car interface {
	Drive() string
	StartEngine() string
}
type SUV struct {
	Brand  string
	Height int
}

type Sedan struct {
	Brand  string
	Length int
}

func (s SUV) Drive() string {
	return fmt.Sprintf("%s Brand SUV is driving %vmm height of car", s.Brand, s.Height)
}

func (s Sedan) Drive() string {
	return fmt.Sprintf("%s Sedan is driving %vfeet length of car", s.Brand, s.Length)
}
func (s SUV) StartEngine() string {
	return fmt.Sprintf("%s SUV engine started", s.Brand)
}
func (s Sedan) StartEngine() string {
	return fmt.Sprintf("%s Sedan engine started", s.Brand)
}
// Polymorphism
func TestDrive(c Car){
	fmt.Println(c.Drive())
	fmt.Println(c.StartEngine())
}
func main() {
	var myCar Car
	myCar = SUV{Brand: "Toyota", Height: 170}
	TestDrive(myCar)

	myCar = Sedan{Brand: "Honda", Length: 45}
	TestDrive(myCar)
}